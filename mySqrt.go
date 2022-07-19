package main

import "fmt"

func mySqrt(x int) int {
	res := 0
	for i := 0; i <= x; i++ {
		if i*i > x {
			break
		}
		res = i
	}
	return res
}

func main() {
	x := 1
	fmt.Println(mySqrt(x))
}
