package main

import (
	"time"
)

func main() {

	chFib := make(chan int)
	done := make(chan struct{})
	r := 20
	go fib(r, chFib)
	go receiveFib(chFib, done)

	<-done

}

func fib(r int, ch chan<- int) {
	a, b := 0, 1
	for i := 1; i <= r; i++ {
		time.Sleep(time.Millisecond * 200)
		ch <- a // sender
		a, b = b, a+b
	}
	close(ch) // close a channel so that it gives notification to the receiver that the channel has been closed and no longer sends the value
}

func receiveFib(ch <-chan int, done chan<- struct{}) {
	// for {
	// 	v, ok := <-ch
	// 	if !ok { // the channel is closed
	// 		done <- struct{}{}
	// 		//break
	// 		return
	// 		//runtime.Goexit()
	// 	} else {
	// 		println(v)
	// 	}
	// }

	for c := range ch {
		// the range loop iterates until the channel is closed
		// range loop on channels does not give two values only a value that is received from the channel
		println(c)
	}
	done <- struct{}{}
}

// closing a channel is not making the channel nil
// once closed cannot be opened
// once closed cannot send a value
// only the sender can close a channel
// a serder can close but can never check whether a channel is closed or not
