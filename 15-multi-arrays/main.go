package main

func main() {

	arr2d := [2][2]int{{1, 2}, {3, 4}}
	arr3d := [2][2][3]int{{{1, 2, 3}, {4, 5, 6}}, {{7, 8, 9}, {10, 11, 12}}}

	for _, arr1 := range arr2d {
		for _, v := range arr1 {
			print(v, " ")
		}
		println()
	}

	for i := 0; i < len(arr3d); i++ {
		for j := 0; j < len(arr3d[i]); j++ {
			for k := 0; k < len(arr3d[i][j]); k++ {
				print(arr3d[i][j][k], " ")
			}
			println()
		}
	}

	println("using range loop")
	for _, arr1 := range arr3d {
		for _, arr2 := range arr1 {
			for _, v := range arr2 {
				println(v, " ")
			}
			println()
		}
	}

}
