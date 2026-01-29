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

}

// fluent api
// chain of actions

//change icalc and calc both from int to any
// and implement Add, Sub, Mul, Div for all number types
