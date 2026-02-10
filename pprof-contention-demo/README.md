# pprof contention demo (goroutines + channels)

This is a **small Go service** that intentionally creates:
- goroutines blocked on **channel send/recv**
- **mutex contention**
- artificial delays to make blocking visible

So you can learn to use:
- `pprof` **goroutine**, **block**, **mutex**, **cpu** profiles
- `go tool trace` (runtime execution trace)

## Prereqs
- Go 1.22+

## Run
```bash
go run .
```

Open:
- pprof index: http://localhost:8080/debug/pprof/
- stats:       http://localhost:8080/stats

## Create load (causes channel & mutex contention)
```bash
curl -XPOST http://localhost:8080/load \
  -H 'content-type: application/json' \
  -d '{"goroutines":400,"items":4000}'
```

Wait ~5-10 seconds and keep it running while you capture profiles.

---

## 1) Goroutine profile (find blocking stacks)
```bash
go tool pprof -http=:0 http://localhost:8080/debug/pprof/goroutine
```

What to look for:
- `chan send` / `chan receive` in stacks
- functions like `(*Server).handleLoad` goroutines stuck in `s.workCh <- v`
- workers blocked sending to `resultCh`

Tip: switch to **"top"**, **"peek"**, and **"graph"** views.

---

## 2) Block profile (BEST for channel delays / blocking time)
The block profile attributes **time spent blocked** (channel ops, mutex, etc).

```bash
go tool pprof -http=:0 http://localhost:8080/debug/pprof/block
```

What to look for:
- `runtime.chan*` and your `(*Server).worker` / `handleLoad`
- this highlights *where time is lost due to blocking*

---

## 3) Mutex profile (lock contention hotspots)
```bash
go tool pprof -http=:0 http://localhost:8080/debug/pprof/mutex
```

Look for:
- `(*Server).worker` holding `s.mu` (intentional contention)
- percent / samples show where lock wait time accumulates

---

## 4) CPU profile (optional)
```bash
go tool pprof -http=:0 'http://localhost:8080/debug/pprof/profile?seconds=10'
```

This will show CPU hotspots (e.g., `spin()`).

---

## 5) Execution trace (best end-to-end: delays, scheduling, net/http, goroutines)
### Option A: Use runtime trace endpoint in this demo (writes trace.out)
```bash
curl 'http://localhost:8080/trace/start?seconds=10'
# then open it:
go tool trace trace.out
```

### Option B: Use built-in pprof trace endpoint (downloads a trace)
```bash
curl -o trace.pprof 'http://localhost:8080/debug/pprof/trace?seconds=5'
go tool trace trace.pprof
```

In `go tool trace`, use:
- **Goroutine analysis**
- **User regions / tasks**
- **Network blocking profile**
- **Sync block profile**

---

## How this demo creates contention

- `workCh` has **buffer=2** and `/load` launches many goroutines that send work fast ⇒
  goroutines block on `s.workCh <- v` (channel send blocking).
- Workers send results on an **unbuffered** `resultCh`, while a slow consumer sleeps ⇒
  workers block on `s.resultCh <- v`.
- `s.mu` is held ~2ms per work item ⇒ lock contention shows in mutex profile.

---

## Cleanup
Stop the server with Ctrl+C.
