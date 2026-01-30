package main

func main() {

	func() {
		println("Hello Tata Elxsi minds!")
	}() // Executer

	var fn1 func()

	fn1 = func() {
		println("I am learing Golang")
	}

	fn1()

	r1 := func(a, b int) int {
		return a + b
	} // There is no executor since the func itself is stored in a variable

	r := r1(10, 20)

	println(r)

	r2 := func(a, b int) int {
		return a - b
	}(100, 20) // execute the function since the result is int

	println(r2)

	maxFn := func(slice []int) int {
		m := slice[0]
		for _, v := range slice {
			m = max(m, v)
		}
		return m
	}

	mx := maxFn([]int{32, 2, 4, 56, 65, 534, 34, 34, 65, 673, 23})
	println("max", mx)

	var fn2 Fn = func() {
		mx := maxFn([]int{32, 2, 4, 56, 65, 534, 34, 34, 65, 673, 23}) // calling another function from the scope
		println("max", mx)
	}
	if fn2 != nil {
		fn2()       // executing a func
		fn2.Greet() // calling a methof from a variable called fn2

	}

	var fn3 Fn

	if fn3 == nil {
		println("yes fn3 is nil, not assigned any func to it")
	}

	fn3 = fn2.Greet // ultimately fn3 requires a function with the sig func() , it can be a pure func, a method with same sig or it can be another funcs varialbe
	fn3()

	fn3 = Greet
	fn3()

}

type Fn func() // Fn is a type not a variable

func (f Fn) Greet() {
	println("Hello This is a Fn type, Greet Method")
}

func Greet() {
	println("Hello Everyone , I am Greet function")
}

// a function without a name is called as anonymous function
// is anonymous function a closure ? yes
// can a func be nil? Yes .. Why? a function has a simple pointer internally , since there is a pointer , it could be nil
// a func can be a type, and also can be stored in a variable.
// a func can be used as a variable in Go

// tasks
// create a min function
// create a anonymous function with []any that should give max and min when any is a number type
