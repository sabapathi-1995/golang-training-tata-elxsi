package main

import (
	"fmt"
	"math/rand/v2"
	"reflect"
	"unsafe"
)

func main() {

	slice1 := make([]int, 5, 10) // or make([]int,5) automatically the capasity is 5
	// [0 0 0 0 0] There are zero values to the slice if not assigned any value
	// make is a built in func --> it allocates memory based on the arguments passed
	// make to be used in three instances slice, map, chan
	// make does it allocate on heap? --> It depends

	fmt.Println(slice1)
	for i := range slice1 {
		slice1[i] = rand.IntN(999)
	}
	fmt.Println(slice1)

	slice2 := slice1 // created a new slice from the existing slice

	PrintSliceHeader(slice1)
	PrintSliceHeader(slice2)

	slice2[0] = 99999
	fmt.Println(slice1)

	slice1 = append(slice1, 7777)
	fmt.Println(slice2)
	// slice1 : Ptr:0x1400001a140 Len:6 Cap:10
	// slice2:  Ptr:0x1400001a140 Len:5 Cap:10
	slice1 = append(slice1, 8888)
	slice1 = append(slice1, 9999, 1111, 333, 444)
	PrintSliceHeader(slice1) // 0x140000b6000 0x140000aa050
	slice2[1] = 434343
	fmt.Println(slice1)

}

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
