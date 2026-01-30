package main

import "math/rand/v2"

func main() {

	fn1 := func() func() {
		println("This is the first func")
		return func() {
			println("This is a return func")
		}
	}

	// fn2 := fn1()
	// fn2()
	fn1()()

	var fn3 Fn = func() func() {
		println("This is the first func")
		return func() {
			println("This is a return func")
		}
	}

	r := fn3.GetRandSq()
	println(r)
	fn4 := fn3()
	fn4()
	fn3()()

	var mymap map[string]any

	mymap["add"] = func(a, b int) int {
		return a + b
	}

	mymap["sub"] = sub

	mymap["fn-return-fn"] = fn3

	mymap["some-int"] = 12312

	var fnmul FnR = func(a, b int) int {
		return a * b
	}
	mymap["mul"] = fnmul

	//mymap["mymap"] = mymap

}

type FnR func(int, int) int

type Fn func() func()

func (f Fn) GetRandSq() int {
	r := rand.IntN(100)
	return r * r
}

func sub(a, b int) int {
	return a - b
}

// Task

// iterate thru mymap
// execute all functions in value of mymap
