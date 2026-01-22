package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func main() {

	slice1 := []int{10, 20, 30, 40, 50}
	slice2 := make([]int, 10, 20)
	for i := range slice2 {
		slice2[i] = rand.IntN(100)
	}

	s1 := SumOf(slice1)
	s2 := SumOf(slice2)

	println(s1, s2)

	var slice3 []int // slice is nil
	s3 := SumOf(slice3)
	println(s3)
	if s4, err := SumOfL(slice3); err == nil {
		println(s4)
	} else {
		println(err.Error())
	}

	arr1 := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	//SumOf(arr1) // array can be converted to slice

	slice4 := arr1[:] // The pointer of arr1[0]-> Ptr in slice header, the length of arr1 Len:10, Cap:10
	slice4[0] = 99999
	fmt.Println(arr1)
	slice4 = append(slice4, 8888)
	slice4[1] = 1111
	fmt.Println(arr1)

	s5 := SumOf(arr1[:])
	println(s5)

}

func SumOf(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

// The ideamatic approach of Go is error to be the last return value

func SumOfL(slice []int) (int, error) {
	if slice == nil {
		return 0, errors.New("nil slice")
	}
	sum := 0
	for i := 0; i < len(slice); i++ {
		sum += slice[i]
	}
	return sum, nil
}

func SumOfL1(slice []int) (error, int) { // technically not wrong but not a ideamatic approach of Go
	if slice == nil {
		return errors.New("nil slice"), 0
	}
	sum := 0
	for i := 0; i < len(slice); i++ {
		sum += slice[i]
	}
	return nil, sum
}
