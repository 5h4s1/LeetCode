package main

func longestCommonPrefix(strs []string) string {
	start := strs[0]
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		ch := ""
		l := len(start)
		if len(start) > len(strs[i]) {
			l = len(strs[i])
		}
		for j := 0; j < l; j++ {
			if start[j] != strs[i][j] {
				break
			}
			ch += string(strs[i][j])
		}
		prefix = ch
		start = prefix

	}
	return prefix
}

func main() {
	a := []string{"f1lowdsadsad", "f1lower", "f1light"}
	print(longestCommonPrefix(a))
}
