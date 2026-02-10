run:
	go run .

load:
	curl -XPOST http://localhost:8080/load -H 'content-type: application/json' -d '{"goroutines":400,"items":4000}'

stats:
	curl http://localhost:8080/stats

pprof-goroutine:
	go tool pprof -http=:0 http://localhost:8080/debug/pprof/goroutine

pprof-block:
	go tool pprof -http=:0 http://localhost:8080/debug/pprof/block

pprof-mutex:
	go tool pprof -http=:0 http://localhost:8080/debug/pprof/mutex

pprof-cpu:
	go tool pprof -http=:0 'http://localhost:8080/debug/pprof/profile?seconds=10'

trace:
	curl 'http://localhost:8080/trace/start?seconds=10'
	go tool trace trace.out
