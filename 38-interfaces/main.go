package main

import (
	"demo-shapes/shapes"
	"demo-shapes/shapes/rect"
	"demo-shapes/shapes/square"
)

func main() {

	// 1. create objects and call Shape function

	r1 := rect.New(10.2, 12.4)
	r2 := rect.New(12.3, 10.5)

	s1 := square.New(12.34)
	s2 := square.New(25.45)

	shapes.Shape(r1)
	println()
	shapes.Shape(r2)
	println()
	shapes.Shape(s1)
	println()
	shapes.Shape(s2)

	// 2. create a slice of IShape append all the objects and call the Shape function

	var shapeSlice []shapes.IShape // it is nil

	//shapeSlice = make([]shapes.IShape, 10)

	shapeSlice = append(shapeSlice, r1, r2, s1, s2) // how can I pass all 4 objects at once ?

	for _, s := range shapeSlice {
		shapes.Shape(s)
	}

	// 3. keep the slice as any and type assert with shape.IShape

	shapeAnySlice := make([]any, 0)
	shapeAnySlice = append(shapeAnySlice, r1, r2, s1, s2, true, "Hello World")
	println()
	for _, s := range shapeAnySlice {
		switch s := s.(type) {
		case shapes.IShape:
			shapes.Shape(s)
		default:
			println("cannot call shape becase of a different type")
		}
	}
}

// create new shapes called Cuboid and also Circle
// make sure that these new shapes implement the interface called IShape

// in the main create those objects and execute shapes.Shape
