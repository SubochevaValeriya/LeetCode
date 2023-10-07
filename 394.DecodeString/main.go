package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	s := "2[abc]3[cd]ef"
	fmt.Println(decodeString(s))
}

type toRepeat struct {
	str   string
	times int
}

func decodeString(s string) string {
	toRepeatInString := []toRepeat{}
	answer := ""
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			times := string(s[i])
			str := ""
			openBrackets, closeBrackets := 0, 0
			for j := i + 1; j < len(s); j++ {
				if unicode.IsDigit(rune(s[j])) && openBrackets == 0 {
					times += string(s[j])
					continue
				}
				if openBrackets != 0 {
					if s[j] == ']' && closeBrackets == openBrackets-1 {
					} else {
						str += string(s[j])
					}
				}
				if s[j] == '[' {
					openBrackets++
				}
				if s[j] == ']' {
					closeBrackets++
				}
				if openBrackets != 0 && openBrackets == closeBrackets {
					timesInt, _ := strconv.Atoi(times)
					toRepeatInString = append(toRepeatInString, toRepeat{
						str:   str,
						times: timesInt,
					})
					i = j
					break
				}
			}
		} else {
			toRepeatInString = append(toRepeatInString, toRepeat{
				str:   string(s[i]),
				times: 1,
			})
		}
	}

	for _, repeat := range toRepeatInString {
		part := ""
		for i, ch := range repeat.str {
			if unicode.IsDigit(ch) {
				part += decodeString(repeat.str[i:])
				break
			}
			part += string(ch)
		}
		answer += strings.Repeat(part, repeat.times)
	}

	return answer
}
