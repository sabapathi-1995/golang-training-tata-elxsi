package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	ch := make(chan string)
	wg := new(sync.WaitGroup)

	// go PublishProblem("publisher-1", ch, time.Millisecond*200, time.Second*1)
	// go PublishProblem("publisher-2", ch, time.Millisecond*200, time.Second*2)
	wg.Add(2)
	go PublishSolution("publisher-1", wg, ch, time.Millisecond*200, time.Millisecond*500)
	go PublishSolution("publisher-2", wg, ch, time.Millisecond*200, time.Millisecond*600)

	// for c := range ch {
	// 	println(c)
	// }

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(ch)
		done <- struct{}{}
		//close(done)
	}()

	//	wg.Add(1) this is creates problems
	go func() {
		for {
			v, ok := <-ch
			if ok {
				println(v)
			} else {
				//wg.Done()
				break
			}
		}
	}()

	<-done
	//runtime.Goexit()
}

func PublishProblem(name string, ch chan string, s time.Duration, d time.Duration) {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println(r)
	// 	}
	// }()
	i := 1
	timeOut := time.After(d)
	for {
		select {
		case <-timeOut:
			close(ch)
			runtime.Goexit()
		//	break
		default:
			time.Sleep(s)
			ch <- fmt.Sprint("Publisher:", name, ">> Value:", i)
		}
		i++
	}
}

func PublishSolution(name string, wg *sync.WaitGroup, ch chan string, s time.Duration, d time.Duration) {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println(r)
	// 	}
	// }()
	i := 1
	timeOut := time.After(d)
	for {
		select {
		case <-timeOut:
			//close(ch)
			wg.Done()
			runtime.Goexit()
		//	break
		default:
			time.Sleep(s)
			ch <- fmt.Sprint("Publisher:", name, ">> Value:", i)
		}
		i++
	}
}

// How does a goroutne know that a channel is closed?
// One is on a range loop once it exit the loop ,can think that the channel is closed
// can use v,ok:=<-ch the ok tells that the channel is closed or not false/true

// sender can only close a channel , but sender never knows whether a channel is closed or not
