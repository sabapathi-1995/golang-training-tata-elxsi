package main

import "time"

func main() {

	ch := GenerateFib(10)
	<-receiveFib(ch)

}

// generator pattern is to genersate data and send thru the channel

func GenerateFib(r int) <-chan int {
	ch := make(chan int) // send channel
	go func() {
		a, b := 0, 1
		for i := 1; i <= r; i++ {
			time.Sleep(time.Millisecond * 200)
			ch <- a // sender
			a, b = b, a+b
		}
		close(ch) // close a channel so that it gives notification to the receiver that the channel has been closed and no longer sends the value
	}()
	return ch
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
	}()
	return done
}
