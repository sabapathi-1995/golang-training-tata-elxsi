package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	ch := Generate("Gen-1 ", time.Millisecond*200, time.Second*1)
	outCh1 := make(chan string)
	outCh2 := make(chan string)
	outCh3 := make(chan string)

	FanOut(ch, outCh1, outCh2, outCh3)

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func(chs ...chan string) {
		for _, ch := range chs {
			wg.Add(1)
			go func() {
				for c := range ch {
					println(c)
				}
				wg.Done()
			}()

		}
		wg.Done()
	}(outCh1, outCh2, outCh3)
	wg.Wait()
}

func Generate(name string, s time.Duration, d time.Duration) <-chan string {
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

func FanOut(inCh <-chan string, outChs ...chan<- string) {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		for v := range inCh {
			wg.Add(1)
			go func(data string) {
				for i, ch := range outChs {
					ch <- fmt.Sprint("FanOut-", i, " ---->", data)
				}
				wg.Done()
			}(v)
		}
		wg.Done()
	}()
	go func() {
		wg.Wait()
		for _, ch := range outChs {
			close(ch)
		}
	}()
}
