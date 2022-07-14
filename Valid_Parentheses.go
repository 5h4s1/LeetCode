package main

import "fmt"

func isValid(strs string) bool {
	var arr []string
	for _, i := range strs {
		if string(i) == "(" || string(i) == "[" || string(i) == "{" {
			arr = append(arr, string(i))
		} else if len(arr) > 0 {
			if string(i) == ")" && arr[len(arr)-1] != "(" {
				return false
			} else if string(i) == "}" && arr[len(arr)-1] != "{" {
				return false
			} else if string(i) == "]" && arr[len(arr)-1] != "[" {
				return false
			}
			arr = arr[:len(arr)-1]
		} else {
			return false
		}

	}
	return len(arr) == 0
}
func main() {
	fmt.Println(isValid("()"))
}
