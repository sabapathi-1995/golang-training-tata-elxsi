package main

import "time"

func main() {

	chFib := make(chan int)
	done := make(chan struct{})
	r := 20
	go fib(r, chFib)
	go receiveFib(r, chFib, done)

	<-done

}

func fib(r int, ch chan<- int) {
	a, b := 0, 1
	for i := 1; i <= r; i++ {
		time.Sleep(time.Millisecond * 200)
		ch <- a // sender
		a, b = b, a+b
	}
}

func receiveFib(r int, ch <-chan int, done chan<- struct{}) {
	for i := 1; i <= r; i++ {
		println(<-ch)
	}

	done <- struct{}{}
}
