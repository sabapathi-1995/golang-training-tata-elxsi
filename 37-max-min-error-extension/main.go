package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	println(max(10, 20))
	fmt.Println(max(12.2, 12.1)) // takes both int and float bcz 1.18 onwards go supports generics

	var slice []int // slice is nil

	for i := 0; i < 10; i++ {
		slice = append(slice, rand.IntN(100)) // using append . when slice is nil ,and using append would initilize the slice and then appends it
	}

	fmt.Println(slice)
	mx, mi, _ := GetMax(slice)

	fmt.Println("max:", mx, "min:", mi)

	// Working with user defined error

	if mx, mi, err := GetMax([]int{}); err != nil {
		switch err.(type) {
		case *SlcieError:
			se := err.(*SlcieError)
			fmt.Println("ID:", se.Id)
			fmt.Println("Msg:", se.Msg)

		default:
			println(err.Error())
		}

	} else {
		fmt.Println("max:", mx, "min:", mi)
	}
}

// What should be an error ?
// Generally errors.New , fmt.Errorf()
// can creat any time that implements the interface called error is a valid type to be used as error

func GetMax(slice []int) (int, int, error) {
	if slice == nil {
		//return 0, 0, errors.New("nil slice")
		return 0, 0, NewSliceError(101, "nil slice")
	}
	if len(slice) == 0 {
		return 0, 0, NewSliceError(102, "slice with no length")
	}
	mx := slice[0]
	mi := slice[0]

	for _, v := range slice {
		mx = max(v, mx)
		mi = min(v, mi)
	}

	return mx, mi, nil
}

type SlcieError struct {
	Id  uint8
	Msg string
}

func NewSliceError(id uint8, msg string) *SlcieError {
	return &SlcieError{id, msg}
}

func (s *SlcieError) Error() string {
	return fmt.Sprintf("Error Id:%d Error Message:%s", s.Id, s.Msg)
}

// Create Rect with L and B
// get Area , make sure area also returns error
// that error must be a customised error
// RectError Code, Msg
// if rect nil return error
// if l or b is 0 return error saying that length or brewdth is zero something like
