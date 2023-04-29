package main

//func WordBreak(s string, wordDict []string) bool {
//	m := map[string]int{}
//
//	for i, word := range wordDict {
//		m[word] = len(word)
//	}
//
//	for i := 0; i < len(s); i++ {
//		for j := 0; j < len(wordDict); j++ {
//			if i > len(s)-1 {
//				return true
//			}
//			if i+len(wordDict[j])-1 > len(s)-1 {
//				if j == len(wordDict)-1 {
//					return false
//				}
//				continue
//			}
//			if s[i:i+len(wordDict[j])] == wordDict[j] {
//				i = i + len(wordDict[j])
//				j = -1
//			} else if j == len(wordDict)-1 && s[i:i+len(wordDict[j])] != wordDict[j] {
//				return false
//			}
//		}
//	}
//	return true
//}

func wordBreak(s string, wordDict []string) bool {
	m := map[string]struct{}{}
	for _, word := range wordDict {
		m[word] = struct{}{}
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			_, ok := m[s[j:i]]
			if ok && dp[j] {
				dp[i] = true
			}

			_, ok = m[s[j:i]]
		}
	}

	return dp[len(dp)-1]
}

//
//func WordBreakOld(s string, wordDict []string) bool {
//	dp := make([]bool, len(s))
//	dp[0] = true
//
//	for i := 0; i < len(s); i++ {
//		for _, word := range wordDict {
//			if i-len(word) > 0 {
//				if word == s[i:i+len(word)] && dp[i-len(word)] {
//					dp[i] = true
//					i = i + len(word) - 1
//				} else if word == s[i:i+len(word)] {
//					dp[i] = true
//					i = i + len(word) - 1
//				}
//
//				//} else {
//				//	dp[i] = false
//				//}
//			}
//		}
//	}
//	return dp[len(s)]
//}
//
//func WordBreakRecursion(s string, wordDict []string) bool {
//	m := map[string]int{}
//	for _, word := range wordDict {
//		m[word] = len(word)
//	}
//
//	answer, _ := wordBreakHelper(s, m)
//	return answer
//
//	//for word, l := range m {
//	//	if s == word {
//	//		return true
//	//	}
//	//	if len(s) >= l {
//	//		if word == s[:l] {
//	//			if WordBreak(s[l:], wordDict) {
//	//				m[s[l:]] = len(s[l:])
//	//				return true
//	//			}
//	//		}
//	//	}
//	//}
//	//return false
//}
//
//func wordBreakHelper(s string, m map[string]int) (bool, map[string]int) {
//	for word, l := range m {
//		if s == word {
//			return true, m
//		}
//		if len(s) >= l {
//			if word == s[:l] {
//				if answer, m := wordBreakHelper(s[l:], m); answer {
//					m[s[l:]] = len(s[l:])
//					return true, m
//				}
//			}
//		}
//	}
//	return false, m
//}
//
//func WordBreak3(s string, wordDict []string) bool {
//	for _, word := range wordDict {
//		if s == word {
//			return true
//		}
//		if len(s) >= len(word) {
//			if word == s[:len(word)] {
//				if WordBreak(s[len(word):], wordDict) {
//					return true
//				}
//			}
//		}
//	}
//
//	return false
//}
//
//func WordBreak2(s string, wordDict []string) bool {
//	m := map[string]int{}
//	lastLenght := 0
//	var checked bool
//
//	for _, word := range wordDict {
//		m[word] = 0
//	}
//
//	if len(s) == 1 {
//		for word := range m {
//			if word == s {
//				return true
//			} else {
//				return false
//			}
//		}
//	}
//
//	for i := 0; i < len(s); i++ {
//		for j := i; j <= len(s); j++ {
//			if j == len(s) {
//				for _, val := range m {
//					if val == 0 {
//						checked = false
//					}
//				}
//
//				if checked == false {
//					i = i - lastLenght
//					j = i + lastLenght
//					if i < 0 {
//						i = 0
//						j = lastLenght
//					}
//					fmt.Println(i, j, "d")
//				} else {
//					return false
//				}
//
//				//if _, ok := m[s[i-lastLenght:j]]; ok {
//				//	i = j
//				//	break
//				//}
//				//return false
//			}
//			if _, ok := m[s[i:j+1]]; ok {
//				fmt.Println(i, j, "tr")
//				m[s[i:j+1]] = 1
//				lastLenght = len(s[i : j+1])
//				fmt.Println(lastLenght, "ll")
//				i = j
//				//	checked = s[i : j+1]
//				break
//			}
//		}
//	}
//
//	return true
//}
