package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	return n + fibonacci(n-1)
}

func main() {
	var n int
	fmt.Scanln(&n)
	fmt.Println(fibonacci(n))
}
