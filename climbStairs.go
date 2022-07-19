package main

import "fmt"

func climbStairs(n int) int {
	num1 := 0
	num2 := 1
	series := 0
	for i := 0; i <= n; i++ {
		num1 = num2
		num2 = series
		series = num1 + num2
	}
	return series
}

func main() {
	n := 4
	fmt.Println(climbStairs(n))
}
