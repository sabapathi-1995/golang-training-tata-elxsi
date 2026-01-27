package main

import "fmt"

var Counter int = 100

func main() {
	fmt.Println(&Counter)

	var num1 int = 100
	var ptr1 *int = &num1 // what is the size of this

	var ok1 bool = true

	var ptr2 *bool = &ok1 // What is the size of this

	fmt.Println(ptr1, ptr2)
	fmt.Println(*ptr1, *ptr2)

	*ptr2 = false

	var ptr3 *string //

	var str1 = "Hello World"
	ptr3 = &str1 // What addresss of the string is there
	fmt.Println(ptr3)

	//var ptr4 *byte =

	bytes := ([]byte)(str1) // This is the ptr inside string header

	ptr4 := &bytes[0]
	fmt.Println(ptr4)
	*ptr4 = 66

	fmt.Println(string(bytes))

}

// Pointers hold addressers
// address --> memory location of a data
//
