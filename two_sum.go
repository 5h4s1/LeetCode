package main

import "fmt"

//func twoSum(nums []int, target int) []int {
//	var res []int
//	for i := 0; i < len(nums); i++ {
//		for j := i + 1; j < len(nums); j++ {
//			if nums[i]+nums[j] == target {
//				res = append(res, i, j)
//				return res
//			}
//		}
//	}
//	return res
//}

func twoSum(nums []int, target int) []int {
	var res []int
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[target-nums[i]] += 1

	}
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] != 2 && nums[i]*2 == target {
			continue
		}

		if m[nums[i]] != 0 {
			res = append(res, i)
		}
	}
	fmt.Println(m)
	return res
}

func main() {
	nums := []int{0, 3, -3, 4, -1}
	fmt.Println(twoSum(nums, -1))
}
