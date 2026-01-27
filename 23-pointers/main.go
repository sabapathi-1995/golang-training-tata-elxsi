package main

var G int

func main() {

	//var num1 int = 100
	//var ptr1 = &num1

	ptr2 := new(int) // what is the value that the address is referenced to

	var arr1 [100000]int // 8 lakh bytes , moved to heap
	// var arr2 [10]int
	println(&G)
	println(ptr2)
	println(&arr1[0])
	//println(&arr2[0])

	//ptr2 = nil

	println(*ptr2)
	// for i := range arr1 {
	// 	println(i)
	// }

	*ptr2 = 100

	println(*ptr2)

	sq1 := getSq1(100)
	println(sq1)

	sqptr1 := getSq2(100)
	println(*sqptr1)

	sqptr2 := getSq3(100)
	println(*sqptr2)

	ptr3 := new(int)
	getSq4(100, ptr3)
	println(*ptr3)

	slice1 := make([]int, 10)

	println(&slice1[0])

	map1 := make(map[string]string, 10)
	println(map1)

}

// 1. Null pointer dereference --> Yes this is problem, so everywhere you need to check whether nil or not
// 2. Double Free --> we dont free , so double free
// 3. Deallocate a null pointer --> we dont allocate or deallocate dirctly like c, so no issues, even use new, it would be automatically deallocated
// 4. Dangling pointer --> Handled by escape analysis in Go
// 5. Memory leak --> by GC in Go
// 6. Use after Null -> Yes need to be chcked like the first one

// Many of these issues are handled by Go Runtime+GC
// Nil or not to be checked by the developer, if it is a pointer or some data structures like slice, map, chan etc..

func getSq1(num int) int {
	sq := num * num // I created a new variable, what is the lifetime of the variable
	return sq
}

func getSq2(num int) *int {
	sq := new(int)
	*sq = num * num // I created a new variable, what is the lifetime of the variable
	return sq       //escaped to heap
}

func getSq3(num int) *int {
	sq := num * num // I created a new variable, what is the lifetime of the variable
	return &sq
}

func getSq4(num int, ptr *int) {
	*ptr = num * num // I created a new variable, what is the lifetime of the variable

}
