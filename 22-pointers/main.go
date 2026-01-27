package main

import (
	"fmt"
)

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
	if ptr3 == nil {
		println("nil pointer string")
	} else {

		fmt.Println(ptr3, *ptr3)
	}

	/*
		String Header
		------------
		Ptr
		Len

	*/

	//var ptr4 *byte =

	*ptr3 = "你好，世界, How are you doing! "
	// ，,

	bytes := ([]byte)(str1) // This is the ptr inside string header

	for i, v := range *ptr3 {
		println(i, string(v))
	}
	ptr4 := &bytes[0]
	fmt.Println(ptr4)
	*ptr4 = 66

	println("len of bytes:", len(bytes))
	for i := 0; i < len(bytes); i++ {
		print(string(bytes[i]))
	}

	println()
	//fmt.Println(string(bytes))

	runes := ([]rune)(str1) // ptr, len, cap

	println("len of runes:", len(runes))
	for i := 0; i < len(runes); i++ {
		print(string(runes[i]))
	}

	println()
}

// Pointers hold addressers
// address --> memory location of a data
// Raw pointers --> The addresses of data
