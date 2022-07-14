package main

import "fmt"

func romanIsInt(s string) int {
	sum := 0
	m := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	for i := len(s) - 1; i >= 0; i-- {
		if 4*m[string(s[i])] < sum {
			sum -= m[string(s[i])]
		} else {
			sum += m[string(s[i])]
		}
	}
	return sum
}

func main() {
	fmt.Println(romanIsInt("LVIII"))
}
