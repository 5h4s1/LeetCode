package main

import "fmt"

func addBinary(s string, t string) string {
	i := len(s) - 1
	j := len(t) - 1
	add := 0
	res := ""
	for i >= 0 || j >= 0 || add > 0 {
		x := 0
		y := 0
		if i >= 0 {
			x = int(s[i]) - 48
		}
		if j >= 0 {
			y = int(t[j]) - 48
		}
		c := x + y + add
		if c == 0 {
			add = 0
			res = "0" + res
		} else if c == 1 {
			add = 0
			res = "1" + res
		} else if c == 2 {
			add = 1
			res = "0" + res
		} else if c == 3 {
			add = 1
			res = "1" + res
		}
		i--
		j--
	}
	return res
}

func main() {
	s := "1101"
	b := "1100"
	fmt.Println(addBinary(s, b))
}
