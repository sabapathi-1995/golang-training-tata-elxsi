package main

import (
	"fmt"
	"math/rand/v2"
	"reflect"
	"unsafe"
)

func main() {

	var slice1 []int // This is a slice // Ptr: nil, Len: 0, Cap: 0 --> 24

	fmt.Println("Type of the slice1:", reflect.TypeOf(slice1))
	fmt.Println("Size of the slice1:", unsafe.Sizeof(slice1))

	fmt.Printf("Address of the slice1:%p\n", &slice1)
	if slice1 != nil && len(slice1) > 0 {
		fmt.Printf("Ptr of the slice1:%p\n", &slice1[0])

	}
	fmt.Println("Len of slice1:", len(slice1))
	fmt.Println("Cap of slice1:", cap(slice1))

	PrintSliceHeader(slice1)

	slice2 := []int{10, 20, 30} // It  gets instantiated automatically --> {} Len:3 Cap:3

	fmt.Printf("Address of the slice2:%p\n", &slice2)
	fmt.Printf("Ptr of the slice2:%p\n", &slice2[0])

	PrintSliceHeader(slice2)

	// slice3 := []int{} // Is it a nil slice ?

	// str1 := ""
	// var str2 string

	// var slice4 []int  // nil
	slice5 := []int{}        // not nil , Ptr: dummy ptr Len:0 Cap:0
	slice6 := make([]int, 0) // not nil, Len:0 Cap:0

	if slice5 == nil {
		println("slice5 is nil")
	}

	if slice6 == nil {
		println("slice6 is nil")
	}
	// println(&slice5[0]) // This would lead to panic as there is no allocation yet ptr has some value so,it is not nil

	slice1 = make([]int, 5, 10) // or make([]int,5) automatically the capasity is 5
	// [0 0 0 0 0] There are zero values to the slice if not assigned any value
	// make is a built in func --> it allocates memory based on the arguments passed
	// make to be used in three instances slice, map, chan
	// make does it allocate on heap? --> It depends

	fmt.Println(slice1)
	for i := range slice1 {
		slice1[i] = rand.IntN(999)
	}
	fmt.Println(slice1)

}

// slice is collection elements of the same type like array but ..
// slice can be extended/grown or can be cutshot/shrunk --> The size of the slice is dynamic
// declare a slice vs instantiate a slice
// declartion does not allocate memory for the original data

func PrintSliceHeader(slice []int) {
	println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println(slice)
	fmt.Println("Type of the slice:", reflect.TypeOf(slice))
	fmt.Println("Size of the slice:", unsafe.Sizeof(slice))
	fmt.Println("Len of slice:", len(slice))
	fmt.Println("Cap of slice:", cap(slice))
	fmt.Printf("Address of the slice:%p\n", &slice)
	if slice != nil && len(slice) > 0 {
		fmt.Printf("Ptr of the slice:%p\n", &slice[0])
	}
	println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
}

// Slice header
// Ptr
// Len
// Cap

// 0x1400012e000
// Slice1{
// 	Ptr:nil
// 	Len:0
// 	Cap:0
// }

// slice2{
// 	Ptr:0x14000016120
// 	Len:3
// 	Cap:3
// }

// What can be nil ? --> slice, interface, map, chan, ptr, func
// Where ever there is a ptr directly or inside the header of the data structure (SliceHeader, Interface Heder etc), that can be nil
// Except string
