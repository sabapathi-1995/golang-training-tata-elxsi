package main

func main() {

	var a int = 100

	defer println("6:", a) // 6 100

	defer func() {
		a++
		println("5:", a) // 5
	}()

	defer func(a int) {
		a++
		println("4:", a) // 4
	}(a)

	defer func(a *int) {
		*a++
		println("\n3:", *a)
	}(&a) // 3

	println("1:", a) // 1 100

	a++

	println("2:", a) // 2 101

	str := "Hello World"
	//println()
	for _, v := range str {
		defer print(string(v))
	} // here is the area it pops out

}
