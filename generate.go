package main

import "fmt"

func generate(numRows int) [][]int {
	var res [][]int
	for row := 0; row < numRows; row++ {
		var arr []int
		for column := 0; column <= row; column++ {
			if column == 0 || column == row {
				arr = append(arr, 1)
				continue
			}
			arr = append(arr, res[row-1][column]+res[row-1][column-1])
		}
		res = append(res, arr)
	}
	return res
}

func main() {
	n := 5
	fmt.Println(generate(n))
}
