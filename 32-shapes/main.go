package main

import (
	"fmt"
	"myshapes/shapes/rect"
	"myshapes/shapes/square"
)

func main() {

	// r1 := NewRect(10.4, 11.34)

	// a1, p1 := r1.Area(), r1.Perimeter()

	// fmt.Printf("Area of Rect r1:%.2f\nPerimeter of Rect r1:%.2f\n", a1, p1)

	// s1 := NewSquare(23.45)

	// println()
	// a2, p2 := s1.Area(), s1.Perimeter()
	// fmt.Printf("Area of Square s1:%.2f\nPerimeter of Square s1:%.2f\n", a2, p2)

	r1 := rect.New(10.4, 11.34)

	a1, p1 := r1.Area(), r1.Perimeter()

	fmt.Printf("Area of Rect r1:%.2f\nPerimeter of Rect r1:%.2f\n", a1, p1)

	s1 := square.New(23.45)

	println()
	a2, p2 := s1.Area(), s1.Perimeter()
	fmt.Printf("Area of Square s1:%.2f\nPerimeter of Square s1:%.2f\n", a2, p2)

	//rand.IntN(100)

}

// To create a package, every package must have a directory
// the immediate directory ideally the name of the package
// you can create different name but not an ideal or idiomatic way
// The root path of the package is the name of the module which is given in go.mod file
