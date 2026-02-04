package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//wg := new(sync.WaitGroup)

	ch := Generate("Gen-1 ", time.Millisecond*200, time.Second*10)

	// wg.Add(1)
	// go Receive(wg, "rec-1", ch)

	// wg.Add(1)
	// go Receive(wg, "rec-2", ch)

	// wg.Add(1)
	// go Receive(wg, "rec-3", ch)

	// wg.Add(1)
	// go Receive(wg, "rec-4", ch)

	// wg.Wait()

	done := Workers(10, ch)
	<-done
	println("Done work")
}

func Generate(name string, s time.Duration, d time.Duration) chan string {
	ch := make(chan string)
	timeoutCh := time.After(d)
	i := 1

	go func() {
		for {
			select {
			case <-timeoutCh:
				close(ch)
				runtime.Goexit()
			default:
				time.Sleep(s)
				ch <- fmt.Sprint("Goroutine:", name, ">> Value:", i)
			}
			i++
		}
	}()
	return ch
}

func Receive(wg *sync.WaitGroup, name string, ch chan string) {
	if wg != nil {
		defer wg.Done()
	}
	for c := range ch {
		println(name, ">>>>>>  ", c)
	}
}

func Workers(workers uint, ch chan string) chan struct{} {
	wg := new(sync.WaitGroup)
	done := make(chan struct{})
	go func() {
		for i := range workers {
			wg.Add(1)
			go func() {
				for c := range ch {
					println(fmt.Sprint("Wroker->", i+1), ">>>>>>  ", c)
				}
				wg.Done()
			}()
		}
		wg.Wait()
		done <- struct{}{}
		close(done)
	}()
	return done
}
