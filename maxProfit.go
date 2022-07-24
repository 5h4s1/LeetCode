package main

import "fmt"

func maxProfit(prices []int) int {
	l := len(prices)
	min := prices[0]
	max := prices[1]
	res := max - min
	for i := 1; i < l; i++ {
		if min > prices[i] {
			min = prices[i]
			max = prices[i]
		}
		if max < prices[i] {
			max = prices[i]
		}
		if max-min > res {
			res = max - min
		}
	}
	return res
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
