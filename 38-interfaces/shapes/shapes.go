package shapes

import "fmt"

type IShape interface {
	IArea
	IPerimeter
	IWhat
}

type IArea interface {
	Area() float64
}

type IPerimeter interface {
	Perimeter() float64
}

type IWhat interface {
	What() string
}

func Shape(ishape IShape) {
	fmt.Printf("Area of %s: %.2f\n", ishape.What(), ishape.Area())
	fmt.Printf("Perimeter of %s: %.2f\n", ishape.What(), ishape.Perimeter())
}
