package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	println("start of main")
	defer println("End of main")

	/*
		// uncomment to check divide by zero panic
			num := 100
			 for i := 10; i >= 0; i-- {
				println(num / i) // at some point of time , i become zero and it would be divid by aero error which is a runtime panic
			}*/
	// panic: runtime error: integer divide by zero

	/*
			//uncomment to check divide by zero panic
		arr := [5]int{10, 11, 12, 13, 14}

		for i := 0; i <= len(arr); i++ {
			println(arr[i])
		}
	*/

	// panic: runtime error: index out of range [5] with length 5

	/*
				//uncomment to check divide by zero panic
		var ptr *int

		*ptr = 100 // there is a panic bvz it is nil pointer dereference

		println(*ptr)

		// panic: runtime error: invalid memory address or nil pointer dereference
		// [signal SIGSEGV: segmentation violation code=0x2 addr=0x0 pc=0x100af1690]
	*/
	func() { // func1
		println(" This is before panic in func1")
		defer println("End of func1")
		func() { // func2
			defer println("End of func2")
			println(" This is before panic in func2")
			PrintEven() //
			println("This is after panic inside func2")
		}()
	}()
	defer println("End of after panic in main.func1.func2.PrintEven")
	println("This is after panic in main")

}

func PrintEven() {
	for i := 1; ; i++ {
		r := rand.IntN(1000)
		if r%2 == 0 {
			println(r)
			if r%9 == 0 {
				panic(fmt.Sprintf("The value %d is divisible by 9 so simply panicking it", r))
			}
		}

	}

}

// What is panic --> The call cannot be executed or continued to execute so it panics and crashes the application
// runtime panic and userdefined panics
// panics panics the whole call stack
// panic panics the whole call stack but looks for deffered functions beofre stacked in the entire call stack ..

// error --> panic --> fatal
