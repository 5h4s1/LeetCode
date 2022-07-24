package main

import "fmt"

func isAlphanumeric(b int) int {
	if b > 64 && b < 91 {
		b += 32
	}
	if (b > 47 && b < 58) || (b > 96 && b < 123) {
		return b
	}
	return 0
}
func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left <= right {
		lb, rb := isAlphanumeric(int(s[left])), isAlphanumeric(int(s[right]))
		if lb == 0 {
			left++
		}
		if rb == 0 {
			right--
		}
		if lb == 0 || rb == 0 {
			continue
		}
		if lb != rb {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	s := "race a car"
	fmt.Println(isPalindrome(s))
}
