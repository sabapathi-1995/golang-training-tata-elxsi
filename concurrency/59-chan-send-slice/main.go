package main

import "time"

func main() {

	chFib := make(chan int)

	r := 20

	go fib(r, chFib)

	for i := 1; i <= r; i++ {
		println(<-chFib)
	}

}

func fib(r int, ch chan int) {
	a, b := 0, 1
	for i := 1; i <= r; i++ {
		time.Sleep(time.Millisecond * 200)
		ch <- a // sender
		a, b = b, a+b
	}
}
