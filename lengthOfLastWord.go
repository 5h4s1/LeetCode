package main

import "fmt"

func lengthOfLastWord(s string) int {
	k := 0
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		if s[i] != ' ' {
			for j := i; j >= 0; j-- {
				if s[j] == ' ' {
					return k
				}
				k++
			}
			break
		}

	}
	return k
}
func main() {
	s := "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(s))
}
