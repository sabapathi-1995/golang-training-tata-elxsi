package main

import (
	"demo-mod/shapes"
	"fmt"
)

func main() {

	sq := shapes.NewSquare(12.4)
	fmt.Println("Area of a square", sq.Area())

	//p := sq.perimeter() // cant use it bcz it is unexportable

	//r := shapes.newRect(10.2, 23.2)

	r := shapes.Rect{L: 12.3, B: 32.23}

	a1 := r.Area()
	//p1 := r.perimeter()

	fmt.Println("Area of rect:", a1)

	//shapes.greet()

	shapes.What()

}
