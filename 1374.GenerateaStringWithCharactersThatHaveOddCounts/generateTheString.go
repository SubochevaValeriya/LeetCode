package main

import (
	"strings"
)

func main() {

}

func generateTheString(n int) string {

	//source := make([]string, 0, 26)
	//
	//for i := 'a'; i <= 'z'; i++ {
	//	source = append(source, string(i))
	//}

	if n%2 != 0 {
		return strings.Repeat("a", n)
	}

	return strings.Repeat("a", n-1) + strings.Repeat("b", 1)
}
