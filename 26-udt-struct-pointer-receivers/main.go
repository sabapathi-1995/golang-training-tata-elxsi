package main

import "fmt"

func main() {

	r1 := New(10.5, 12.4)

	a1 := r1.Area()
	p1 := r1.Perimeter()

	fmt.Println("Area of r1:", a1)
	fmt.Println("Perimeter of r1:", p1)

	fmt.Println("Area of r1:", r1.A)
	fmt.Println("Perimeter of r1:", r1.P)

	r2 := New(10.5, 12.4)

	a2 := (*r2).AreaP() // Go can understand that the receiver is a pointer
	p2 := r2.PerimeterP()

	fmt.Println("Area of r2:", a2)
	fmt.Println("Perimeter of r2:", p2)

	fmt.Println("Area of r2:", r2.A)
	fmt.Println("Perimeter of r2:", r2.P)

	r3 := Rect{L: 10.3, B: 14.5}

	a3 := (&r3).AreaP() // Go can understand that the receiver is a pointer, mo need to use as a ref and call the method since it is a pointer receiver
	p3 := r3.PerimeterP()

	fmt.Println("Area of r3:", a3)
	fmt.Println("Perimeter of r3:", p3)

	fmt.Println("Area of r3:", r3.A)
	fmt.Println("Perimeter of r3:", r3.P)

}

type Rect struct {
	L, B float32
	A, P float32
}

// There is no constructor in Go

// The idomatic approach is create a function called New

func New(l, b float32) *Rect { // this is a like a constructor
	return &Rect{L: l, B: b}
}

// func Default() *Rect { // this is a like a constructor
// 	return &Rect{L: 1, B: 1}
// }

func (r Rect) Area() float32 { // pass by value
	r.A = r.L * r.B
	return r.A
}

func (r Rect) Perimeter() float32 { // pass by value
	r.P = 2 * (r.L + r.B)
	return r.P
}

func (r *Rect) AreaP() float32 { // pass/call by reference
	r.A = r.L * r.B
	return r.A
}

func (r *Rect) PerimeterP() float32 { // pass/call by reference
	r.P = 2 * (r.L + r.B)
	return r.P
}
