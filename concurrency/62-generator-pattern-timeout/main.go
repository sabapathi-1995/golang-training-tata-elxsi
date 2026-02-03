package main

import (
	"runtime"
	"time"
)

func main() {

	ch := GenerateFib(time.Microsecond * 10)
	<-receiveFib(ch)

}

// generator pattern is to genersate data and send thru the channel

func GenerateFib(d time.Duration) <-chan int {
	ch := make(chan int) // send channel
	go func() {
		a, b := 0, 1
		for {
			select {
			case <-TimeOut(d):
				close(ch)
				println("Done generating fib numbers")
				runtime.Goexit()
			case ch <- a:
				//time.Sleep(time.Millisecond * 200)
				// sender
				a, b = b, a+b
			}
		}
		// close a channel so that it gives notification to the receiver that the channel has been closed and no longer sends the value
	}()
	return ch
}

// This is our own time out
// There is something called TimeAfter --> please use that ..

func TimeOut(d time.Duration) chan struct{} {
	timeout := make(chan struct{})
	go func() {
		time.Sleep(d)
		timeout <- struct{}{}
		close(timeout)
	}()
	return timeout
}

func receiveFib(ch <-chan int) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		for c := range ch {
			// the range loop iterates until the channel is closed
			// range loop on channels does not give two values only a value that is received from the channel
			println(c)
		}
		done <- struct{}{}
		close(done)
	}()
	return done
}
