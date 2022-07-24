package main

import "fmt"

func getRow(rowIndex int) []int {
	var res [][]int
	for row := 0; row <= rowIndex; row++ {
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
	return res[rowIndex]
}

func main() {
	n := 3
	fmt.Println(getRow(n))
}
