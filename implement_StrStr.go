package main

import "fmt"

func strStr(haystack string, needle string) int {
	index := 0
	l1 := len(needle) - 1
	l2 := len(haystack) - 1
	if l1 == -1 || (l1 == 0 && haystack == needle) {
		return 0
	}
	for index+l1 <= l2 {
		if haystack[index] == needle[0] && haystack[index+l1] == needle[l1] {
			if haystack[index:index+l1+1] == needle {
				return index
			}
		}
		index++
	}
	return -1
}
func main() {
	haystack := "abbc"
	needle := "c"
	fmt.Println(strStr(haystack, needle))
}
