package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"runtime/trace"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

// This demo intentionally creates:
//  1) Goroutines blocked on channels (send/recv)
//  2) Mutex contention
//  3) Artificial latency (sleep / work)
// So you can observe them in pprof "goroutine" profile and via "block" profile,
// and in execution trace (go tool trace).

type Server struct {
	mu        sync.Mutex
	counter   int64
	workCh    chan int
	resultCh  chan int
	workers   int
	inFlight  int64
	blockedS  int64
	blockedR  int64
	startTime time.Time
}

func NewServer() *Server {
	s := &Server{
		// Small buffer makes it easy to create sender blocking.
		workCh:    make(chan int, 2),
		// Unbuffered resultCh makes receivers/senders block more often.
		resultCh:  make(chan int),
		workers:   4,
		startTime: time.Now(),
	}
	return s
}

func (s *Server) Start(ctx context.Context) {
	// Start workers that receive from workCh and send to resultCh.
	for i := 0; i < s.workers; i++ {
		go s.worker(ctx, i)
	}

	// Result consumer: intentionally slow, so worker goroutines block on sending results.
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case v := <-s.resultCh:
				// Slow consumer -> causes workers to block on send to resultCh.
				time.Sleep(15 * time.Millisecond)
				_ = v
			}
		}
	}()
}

func (s *Server) worker(ctx context.Context, id int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(id)))
	for {
		select {
		case <-ctx.Done():
			return
		case v := <-s.workCh:
			atomic.AddInt64(&s.inFlight, 1)

			// Mutex contention hotspot: many workers trying to lock.
			s.mu.Lock()
			// Hold lock a bit to amplify contention.
			time.Sleep(2 * time.Millisecond)
			s.counter += int64(v)
			s.mu.Unlock()

			// Simulated CPU-ish work.
			spin(r.Intn(8000) + 2000)

			// Send result; with slow consumer + unbuffered channel this blocks often.
			select {
			case s.resultCh <- v:
			default:
				// If send would block, count it, then do a blocking send.
				atomic.AddInt64(&s.blockedS, 1)
				s.resultCh <- v
			}

			atomic.AddInt64(&s.inFlight, -1)
		}
	}
}

func spin(n int) {
	x := 0
	for i := 0; i < n; i++ {
		x = (x*33 + i) ^ (x >> 1)
	}
	_ = x
}

// /load will create many goroutines trying to enqueue work faster than workers can consume.
// That makes goroutines block on workCh sends.
func (s *Server) handleLoad(w http.ResponseWriter, r *http.Request) {
	type req struct {
		Goroutines int `json:"goroutines"`
		Items      int `json:"items"`
	}
	in := req{Goroutines: 200, Items: 2000}

	_ = json.NewDecoder(r.Body).Decode(&in)
	if in.Goroutines <= 0 {
		in.Goroutines = 200
	}
	if in.Items <= 0 {
		in.Items = 2000
	}

	for g := 0; g < in.Goroutines; g++ {
		go func(seed int) {
			rnd := rand.New(rand.NewSource(time.Now().UnixNano() + int64(seed)))
			for i := 0; i < in.Items/in.Goroutines; i++ {
				v := rnd.Intn(100)

				// Attempt a non-blocking send to detect and count contention.
				select {
				case s.workCh <- v:
				default:
					atomic.AddInt64(&s.blockedR, 1) // reuse counter for "work send blocked"
					// Blocking send (this is what shows as goroutine blocking on chan send).
					s.workCh <- v
				}

				// Tiny jitter to vary schedules.
				if i%7 == 0 {
					time.Sleep(time.Duration(rnd.Intn(3)) * time.Millisecond)
				}
			}
		}(g)
	}

	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{
		"started":    true,
		"goroutines": in.Goroutines,
		"items":      in.Items,
	})
}

// /stats shows internal counters.
func (s *Server) handleStats(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(s.startTime).String()
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{
		"uptime":                   uptime,
		"counter":                  atomic.LoadInt64(&s.counter),
		"in_flight_workers":        atomic.LoadInt64(&s.inFlight),
		"blocked_on_result_send":   atomic.LoadInt64(&s.blockedS),
		"blocked_on_work_send":     atomic.LoadInt64(&s.blockedR),
		"goroutines":               runtime.NumGoroutine(),
		"workCh_buffered_capacity": cap(s.workCh),
	})
}

// /trace/start starts runtime tracing to a file (trace.out) and stops after duration.
// Usage: curl 'http://localhost:8080/trace/start?seconds=10'
func handleTraceStart(w http.ResponseWriter, r *http.Request) {
	sec := 10
	if v := r.URL.Query().Get("seconds"); v != "" {
		fmt.Sscanf(v, "%d", &sec)
		if sec <= 0 {
			sec = 10
		}
		if sec > 60 {
			sec = 60
		}
	}
	f, err := os.Create("trace.out")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if err := trace.Start(f); err != nil {
		_ = f.Close()
		http.Error(w, err.Error(), 500)
		return
	}

	go func() {
		time.Sleep(time.Duration(sec) * time.Second)
		trace.Stop()
		_ = f.Close()
	}()

	w.WriteHeader(200)
	_, _ = w.Write([]byte("trace started; will stop automatically\n"))
}

func main() {
	// Enable block profiling (captures goroutine blocking events like channel ops, mutex).
	// Rate is in nanoseconds: 1 records every blocking event (high overhead).
	// 10_000 means sample roughly every 10µs of blocking time.
	runtime.SetBlockProfileRate(10_000)

	// Optional mutex profile sampling (captures lock contention).
	runtime.SetMutexProfileFraction(5)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := NewServer()
	s.Start(ctx)

	mux := http.NewServeMux()
	mux.HandleFunc("/load", s.handleLoad)
	mux.HandleFunc("/stats", s.handleStats)
	mux.HandleFunc("/trace/start", handleTraceStart)

	// pprof endpoints:
	//   /debug/pprof/
	//   /debug/pprof/goroutine
	//   /debug/pprof/block
	//   /debug/pprof/mutex
	//   /debug/pprof/profile?seconds=10
	//   /debug/pprof/trace?seconds=5
	// Register default pprof handlers on this mux as well
	mux.Handle("/debug/pprof/", http.DefaultServeMux)
	mux.Handle("/debug/pprof/cmdline", http.DefaultServeMux)
	mux.Handle("/debug/pprof/profile", http.DefaultServeMux)
	mux.Handle("/debug/pprof/symbol", http.DefaultServeMux)
	mux.Handle("/debug/pprof/trace", http.DefaultServeMux)

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		log.Printf("listening on http://localhost:8080")
		log.Printf("pprof at     http://localhost:8080/debug/pprof/")
		log.Printf("trigger load: curl -XPOST http://localhost:8080/load -d '{\"goroutines\":400,\"items\":4000}'")
		log.Printf("stats:        curl http://localhost:8080/stats")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	// Graceful shutdown
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	cancel()

	ctxTO, cancelTO := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelTO()
	_ = srv.Shutdown(ctxTO)
	log.Printf("bye")
}
