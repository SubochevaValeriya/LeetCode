package main

import (
	"fmt"
)

//func main() {
//	fmt.Println(reorganizeString("aabbcc"))
//	fmt.Println(reorganizeString("vvvlo"))
//	fmt.Println(reorganizeString("eqmeyggvp"))
//}

//func reorganizeString(s string) string {
//	m := map[string]int{}
//	answer := make([]rune, len(s))
//	alphabet := []rune{}
//
//	for i := 0; i < len(s); i++ {
//		m[string(s[i])]++
//
//	}
//
//	for letter := 'a'; letter <= 'z'; letter++ {
//		alphabet = append(alphabet, letter)
//	}
//
//	sort.SliceStable(alphabet, func(i, j int) bool {
//		return m[string(alphabet[i])] > m[string(alphabet[j])]
//	})
//
//	var theSameCountForLetter = 1
//	var firstJ = 0
//
//	for i := 0; i < len(alphabet); i++ {
//		count := m[string(alphabet[i])]
//		if count != 0 {
//			if i < len(alphabet)-1 {
//				if count == m[string(alphabet[i+1])] && theSameCountForLetter >= 1 {
//					theSameCountForLetter++
//					if m[string(alphabet[i+2])] < count {
//						answerPart := inAlphabetOrder(alphabet[i-theSameCountForLetter+2:i+2], count)
//						answer = append(answerPart, answer[len(answerPart):]...)
//						firstJ += len(answerPart) - 1
//						if m[string(alphabet[i+2])] == 0 {
//							return string(answer)
//						}
//						i++
//						theSameCountForLetter = -1
//					}
//					continue
//				}
//			}
//			theSameCountForLetter = -1
//			for j := firstJ; count > 0; {
//				if j >= len(s) {
//					return ""
//				}
//				if answer[j] != 0 {
//					j++
//					continue
//				} else {
//					answer[j] = alphabet[i]
//					j += 2
//					count--
//				}
//			}
//			m[string(alphabet[i])] = 0
//		} else {
//			if i == len(alphabet)-1 {
//				i = -1
//			}
//			continue
//		}
//
//		if len(answer) > 1 {
//			if answer[len(answer)-1] == answer[len(answer)-2] && answer[len(answer)-1] != 0 {
//				return ""
//			}
//		}
//
//		if answer[len(answer)-1] != 0 && answer[len(answer)-2] != 0 {
//			return string(answer)
//		}
//
//		if i == len(alphabet)-1 {
//			i = -1
//		}
//	}
//
//	return string(answer)
//}
//
//func inAlphabetOrder(alphabet []rune, n int) []rune {
//	answer := []rune{}
//	for i := 0; i < len(alphabet); i++ {
//		if n == 0 {
//			break
//		}
//		answer = append(answer, alphabet[i])
//		if i == len(alphabet)-1 {
//			n--
//			i = -1
//		}
//	}
//	return answer
//}

func reorganizeStringInAlphabetOrder(s string) string {
	if len(s) < 2 {
		return s
	}
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
		if len(s)%2 == 0 {
			if count >= len(s)/2+1 {
				return ""
			}
		} else {
			if count > len(s)/2+1 {
				return ""
			}
		}

		if count != 0 {
			answer += alphabet[i]
			m[alphabet[i]]--
		} else {
			if i == len(alphabet)-1 {
				i = -1
			}
			continue
		}

		if len(answer) == len(s) {
			if answer[len(answer)-1] == answer[len(answer)-2] {
				answerRune := []rune(answer)
				left, right := 0, len(answerRune)-1
				for right > left {
					fmt.Println(string(answerRune))
					if answerRune[right] == answerRune[right-1] {
						if answerRune[left] != answerRune[right] {
							if left == 0 {
								answerRune = append([]rune{answerRune[right]}, answerRune[:right]...)
							} else if answerRune[left+1] != answerRune[right] {
								answerSecondPart := append([]rune{answerRune[right]}, answerRune[left+1:right]...)
								answerRune = append(answerRune[:left+1], answerSecondPart...)
							}
						}
						left++
					} else {
						return string(answerRune)
					}
				}
			}
			return answer
		}

		if i == len(alphabet)-1 {
			i = -1
		}
	}

	return ""
}
