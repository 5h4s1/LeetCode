package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var a []int
	for x > 0 {
		a = append(a, x%10)
		x = x / 10
	}
	l := len(a)
	for i := 0; i < l; i++ {
		if a[i] != a[l-i-1] {
			return false
		}

	}
	return true
}

func main() {
	fmt.Println(isPalindrome(0))
}
