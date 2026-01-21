package main

import (
	"errors"
	"fmt"
)

func main() {
	println(IsNumber(100.1232))
	println(IsNumber(true))

	if s, err := GetSq(10); err != nil {
		println(err.Error())
	} else {
		fmt.Println("Square:", s)
	}

	if s, err := GetSq(10.1231); err != nil {
		println(err.Error())
	} else {
		fmt.Println("Square:", s)
	}

	if s, err := GetSq(true); err != nil {
		println(err.Error())
	} else {
		fmt.Println("Square:", s)
	}

	if s, err := GetSq("Hello"); err != nil {
		println(err.Error())
	} else {
		fmt.Println("Square:", s)
	}

}

func IsNumber(val any) bool {
	switch val.(type) {
	case int, uint, uint8, uint16, uint32, uint64, int8, int16, int32, int64, float32, float64:
		return true
	default:
		return false
	}
}

func GetSq(val any) (float64, error) {

	if IsNumber(val) {
		switch v := val.(type) {
		case uint:
			return float64(v * v), nil
		case int:
			return float64(v * v), nil
		case int8:
			return float64(v * v), nil
		case uint8:
			return float64(v * v), nil
		case uint16:
			return float64(v * v), nil
		case int16:
			return float64(v * v), nil
		case uint32:
			return float64(v * v), nil
		case int32:
			return float64(v * v), nil
		case uint64:
			return float64(v * v), nil
		case int64:
			return float64(val.(float32) * val.(float32)), nil // instead of asserting, we are using v every where
		case float32:
			return float64(v * v), nil
		case float64:
			return v * v, nil
		}

	} else {
		//return 0, fmt.Errorf("invalid number")
		return 0, errors.New("invalid number")
	}

	return 0, nil
}

// type switch --> only any or interfaces
// To find the type of a variable at runtime
