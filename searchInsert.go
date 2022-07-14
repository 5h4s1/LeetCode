package main

import "fmt"

func searchInsert(num []int, target int) int {
	l := 0
	r := len(num) - 1
	mid := l + (r-l)/2
	if num[0] > target {
		return 0
	}
	if num[r] < target {
		return r + 1
	}
	for l <= r {
		mid = (l + r) / 2
		if num[mid] == target {
			return mid
		} else if num[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
		if num[mid] > target && num[mid-1] < target {
			return mid
		} else if num[mid] < target && num[mid+1] > target {
			return mid + 1
		}
	}
	return -1
}

func main() {
	nums := []int{1, 3}
	target := 3
	fmt.Println(searchInsert(nums, target))
}
