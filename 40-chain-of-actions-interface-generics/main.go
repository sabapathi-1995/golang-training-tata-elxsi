package main

import (
	"demo/calc"
	"demo/icalc"
)

func main() {

	c1 := icalc.New(100)
	r1 := c1.Add(10).Add(20).Sub(10).Mul(3).Div(2).Add(10).Sub(20).Div(2).Get() // These function except Get can be called any number of times bcz it returns the same object here it is an interface
	println("result:", r1)

	c2 := calc.New(100)
	r2 := c2.Add(10).Add(20).Sub(10).Mul(3).Div(2).Add(10).Sub(20).Div(2).Get() // These function except Get can be called any number of times bcz it returns the same object here it is an interface
	println("result:", r2)

	//r := max(10, 20)

	println(add1(10, 20))
	println(add1(10.5, 20.5))
	println(add1(float32(10.5), float32(20.5)))

	// type erasure -->
	// some programming languges use a concept called monomorphization --> The compiler looks at the caller and regenerate a function with the types passed
	// go does not monomophize, go uses a concept called type erasure
	// go does not create a function based on the call
	// other programming languages look at the call and compiler generates a function with params based on the call

	/*

			Go does not do this other programming languages does that
				func addint(a,b int)int{
				return a+b
				}

				func addfloat64(a,b float64)float64{
				return a+b
				}
				func addfloat32(a,b float32)float32{
				return a+b
				}

			Go does type earsure means --> the generic type is replaced by any
		 		add1(a,b any)any
	*/
}

func add1[T int | uint | uint8 | int8 | uint16 | int16 | uint32 | int32 | uint64 | int64 | float32 | float64](a, b T) T {
	return a + b
}

func add2[T INumbers](a, b T) T {
	return a + b
}

type INumbers interface {
	int | uint | uint8 | int8 | uint16 | int16 | uint32 | int32 | uint64 | int64 | float32 | float64
}

//

// fluent api
// chain of actions

//change icalc and calc both from int to any
// and implement Add, Sub, Mul, Div for all number types
