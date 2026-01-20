package main

import (
	"fmt"
	"strconv"
)

func main() {

	var any1 any // any

	any1 = 100 // --? The pointer to 100 and the pointer of int

	// var num1 int = int(any1) // cant do that , type casting is not possible with any, instead do type assertion

	var num1 int = any1.(int) // any1 has a value of type int so , assert it to int and store in a int variable
	println(num1)

	var ok1 bool = true
	any1 = ok1

	num1, ok := any1.(int) // ok is bool either true or false
	if ok {                // if ok==true{}
		println(num1)
	} else {
		ok1, ok = any1.(bool)
		if ok {
			println(ok1)
		} else {
			println("unable to assert to int")
		}
	}

	num1, num2, num3, num4, num5, float1, float2, any2, any3, str1, bool1 := int(988), uint8(123), int16(-12312), int64(-1231232131123), uint32(2331231231), float32(787.34), 2423422323232.2342, any(4435354), any(545345.345345), "312312", true

	//var sum float64 = float64(num1) + float64(num2) + float64(num3) + float64(num4) + float64(num5) + float64(float1) + float2 + float64(any2.(int)) + any3.(float64)

	var sum float64 = float64(num1) + float64(num2) + float64(num3) + float64(num4) + float64(num5) + float64(float1) + float2

	sum1, ok1 := any2.(int)
	if ok1 {
		sum += float64(sum1)
	}

	sum2, ok1 := any3.(float64)
	if ok1 {
		sum += sum2
	}

	sum3, err := strconv.Atoi(str1)
	if err == nil {
		sum += float64(sum3)
	}

	//val := 1

	if bool1 {
		sum += 1 //float64(val) // compiler infers it as 1.0
	}

	fmt.Printf("%.4f", sum)

	//sum = sum + 100.0 //
}
