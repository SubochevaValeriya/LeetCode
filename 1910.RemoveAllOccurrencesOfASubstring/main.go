package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeOccurrences("daabcbaabcbc", "abc"))
}

func removeOccurrences(s string, part string) string {
	for i := 0; i < len(s); i++ {
		if i+len(part) > len(s) {
			fmt.Println(i, s)
			return s
		}
		if s[i:i+len(part)] == part {
			s = s[:i] + s[i+len(part):]
			i = i - len(part)
			if i < -1 {
				i = -1
			}
		}
	}
	return s
}
