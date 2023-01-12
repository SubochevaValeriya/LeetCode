package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(reorganizeString("aab"))
	fmt.Println(reorganizeString("aaab"))
	fmt.Println(reorganizeString("vvvlo"))
	fmt.Println(reorganizeString("aabbcc"))
}

func reorganizeString(s string) string {
	m := map[string]int{}
	answer := make([]rune, len(s))
	alphabet := []rune{}

	for i := 0; i < len(s); i++ {
		m[string(s[i])]++

	}

	for letter := 'a'; letter <= 'z'; letter++ {
		alphabet = append(alphabet, letter)
	}

	sort.SliceStable(alphabet, func(i, j int) bool {
		return m[string(alphabet[i])] > m[string(alphabet[j])]
	})

	for i := 0; i < len(alphabet); i++ {
		count := m[string(alphabet[i])]
		if count != 0 {
			for j := 0; count > 0; {
				if j >= len(s) {
					return ""
				}
				if answer[j] != 0 {
					j++
					continue
				} else {
					answer[j] = alphabet[i]
					j += 2
					count--
				}
			}
			m[string(alphabet[i])] = 0
		} else {
			if i == len(alphabet)-1 {
				i = -1
			}
			continue
		}

		if len(answer) > 1 {
			if answer[len(answer)-1] == answer[len(answer)-2] && answer[len(answer)-1] != 0 {
				return ""
			}
		}

		if answer[len(answer)-1] != 0 && answer[len(answer)-2] != 0 {
			return string(answer)
		}

		if i == len(alphabet)-1 {
			i = -1
		}
	}

	return string(answer)
}

//needToReorganize := false
//sRune := []rune(s)
//for j := 1; j < len(sRune); j++ {
//	if sRune[i] == sRune[j] {
//		needToReorganize = true
//		continue
//	}
//	if sRune[i] != sRune[j] && needToReorganize == true {
//		sRune[i], sRune[j] = sRune[j], sRune[i]
//	} else {
//		break
//	}
//}

//	return string(sRune)

func reorganizeString0(s string) string {
	m := map[string]int{}
	var answer string
	alphabet := []string{}

	for i := 0; i < len(s); i++ {
		m[string(s[i])]++

	}

	for letter := 'a'; letter <= 'z'; letter++ {
		alphabet = append(alphabet, string(letter))
	}

	for i := 0; i < len(alphabet); i++ {

		count := m[alphabet[i]]
		if count != 0 {
			answer += alphabet[i]
			m[alphabet[i]]--
		} else {
			if i == len(alphabet)-1 {
				i = -1
			}
			continue
		}

		if len(answer) > 1 {
			if answer[len(answer)-1] == answer[len(answer)-2] {
				return ""
			}
		}
		if len(answer) == len(s) {
			return answer
		}

		if i == len(alphabet)-1 {
			i = -1
		}
	}

	return ""
}
