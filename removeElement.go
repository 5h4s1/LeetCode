package main

import "fmt"

func removeElement(nums []int, val int) int {
	r := len(nums) - 1
	l := 0
	k := 0
	for r >= l {
		if nums[r] == val {
			r--
			k++
			continue
		} else if nums[l] == val {
			nums[l] = nums[r]
			k++
			r--
		}
		l++
	}
	return len(nums) - k
}

func main() {
	a := []int{2, 2, 2}
	fmt.Println(removeElement(a, 2))
}
