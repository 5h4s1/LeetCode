package main

import "fmt"

func singleNumber(nums []int) int {
	res := 0
	for i := range nums {
		res = res ^ nums[i]
	}

	return res
}

func main() {
	arr := []int{1, 3, 3, 1, 2}
	fmt.Println(singleNumber(arr))
}
