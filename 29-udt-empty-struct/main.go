package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {

	area1 := A1{}.Area(10.3, 14.5)
	area2 := A2{}.Area(10.3, 14.5, 15.34)
	fmt.Printf("a1:%.3f\n", area1)
	fmt.Println("a2:", math.Round(area2))

	a1 := A1{}
	a2 := A1{}
	a3 := A2{}
	area1 = a1.Area(10.3, 14.5)
	fmt.Printf("a1:%.3f\n", area1)
	fmt.Println("size of a1:", unsafe.Sizeof(a1))
	fmt.Printf("address of:%p\n", &a1) // 0x104a94980
	fmt.Printf("address of:%p\n", &a2) // 0x104a94980
	fmt.Printf("address of:%p\n", &a3) // 0x104a94980

	//	address of:0x102ffc980
	//  address of:0x102ffc980
	//  address of:0x102ffc980
}

type A1 struct{} // Size is zero
type A2 struct{}

func (a A1) Area(l, b float32) float64 {
	return float64(l * b)
}

func (a A2) Area(l, b, h float32) float64 {
	return float64(l * b * h)
}
