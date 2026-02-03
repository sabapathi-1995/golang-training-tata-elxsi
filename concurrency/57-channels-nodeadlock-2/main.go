package main

import (
	"sync"
	"time"
)

func main() {

	ch1 := make(chan int) // This is to instantiate a channel. This is unbuffered channel
	wg := new(sync.WaitGroup)

	wg.Add(1)

	go func() {
		println("waiting to receive a value")
		v := <-ch1 // This is a receiver, the arrow mark away from the channel
		println(v)
		wg.Done()
	}()

	time.Sleep(time.Second * 3) // What would happen
	ch1 <- 100                  // This is a sender, the arrow mark towards the channel

	wg.Wait()
}

// chan is a keyword to create a channel
// channels are used for the sync+data-transfer
// there is a sender and there will be a receiver, sender and the receiver are generally goroutines

// 1. The sender is blocked until the receiver receives the data
// 2. The receiver is blocked until the sender sends the data
// 3. the block is very subjective, it is based on the size of the channel (buffered and unbuffered channels)
// 4. there is no order that the sender goroutine has to be the first or next

// a channel can be nil, until make is used
// unbuffered channel that means at any point, only one data value can be sent at a time.The next value can be send only if the previous valus has been received
// to send a value ch1 <- 100
// to received a value from the channel <-ch1
