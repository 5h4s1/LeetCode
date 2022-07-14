package main

import "fmt"

func plusOne(digits []int) []int {
	l := len(digits)
	for i := l - 1; i >= 0; i-- {
		digits[i] = (digits[i] + 1) % 10
		if digits[i] != 0 {
			break
		} else if digits[i] == 0 && i == 0 {
			digits = append([]int{1}, digits...)
		}
	}

	return digits
}

func main() {
	arr := []int{9, 9, 9, 9}

	fmt.Println(plusOne(arr))
}
