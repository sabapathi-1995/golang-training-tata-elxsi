package main

func main() {

	ch := make(chan int, 10)

	//go func() {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
	//}()

	done := make(chan struct{})
	go func() {
		for c := range ch {
			print(c, " >> ")
		}

		// for {
		// 	time.Sleep(time.Millisecond * 100)
		// 	v, ok := <-ch
		// 	println(v, ok)
		// }

		done <- struct{}{}
	}()

	<-done

}
