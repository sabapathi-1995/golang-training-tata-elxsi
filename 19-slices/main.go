package main

import "fmt"

func main() {

	slice1 := make([]int, 10, 15) // Ptr:0x1400123 Len:10, Cap:15

	for i := range len(slice1) {
		slice1[i] = i + 1
	}
	fmt.Println(slice1)

	slice2 := slice1
	slice3 := slice1[:]
	slice4 := slice1[:5]  // 1,2,3,4,5
	slice5 := slice1[3:8] //[4,5,6,7,8]
	slice6 := slice1[8:]  // [9,10]
	fmt.Println("slice2:", slice2)
	fmt.Println("slice3:", slice3)
	fmt.Println("slice4:", slice4)
	fmt.Println("slice5:", slice5)
	fmt.Println("slice6:", slice6)
	// [1 2 3 4 5 6 7 8 9 10 ? ? ? ? ?]
	// slice6: [9 10 ? ? ? ? ?]
	// slice6[0]= 9999 wha would happen to slice1 and slice2

	slice6[0] = 9999
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)
	fmt.Println("slice3:", slice3)
	slice6 = append(slice6, 12, 13) // The slice6 is nolonger has the same address of slice1
	slice6[1] = 8888
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)
	fmt.Println("slice3:", slice3)
	fmt.Println("slice6:", slice6)
	fmt.Println("Hello", "World", "How", "Are", "You", "Doing", 1, true, 1.1)

	arr1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	println("cap:", cap(arr1)) // there is no cap for array still cap is not there it is written in such a way that len is given
	println(SumOf())
	println(SumOf(1, 2))
	println(SumOf(1, 2, 3, 34, 34, 3, 54, 6, 6, 6, 232, 454, 4523, 232))
	println(SumOf(slice1...))
	println(SumOf(arr1[:]...))

	slice7 := append(slice1[0:5], slice1[6:]...)
	fmt.Println("Slice7:", slice7)

	slice8 := []any{1, 2, 3, "hello", true, nil}
	fmt.Println(slice8...)

	// slice9 := arr1[:] // Len:10 , Cap: len(arr1)

}

// Some functions can alloc to pass any number of arguments.. those are called as variadic arguments..
// Variadic parameter should be the last parameter of a function , cannot keep it as first or second, should be the last or only one
// Variadic parameter can only be used in functions and methods, cannot be used as a variable or a field in a struct
// Variadic paramter can allow to pass 0 or N number of arguments.
// A slice can be passed as a variadic argument using some special form
// can use range loop on variadic argument

func SumOf(nums ...int) int { // ...int variadic parameter
	println("len:", len(nums), "cap:", cap(nums))
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func MultiSumOf(num int, nums ...int) int { // MultiSumOf(nums ...int,num int) int  nums must be the last // ...int variadic parameter
	sum := 0
	for _, v := range nums {
		sum += v * num
	}
	return sum
}
