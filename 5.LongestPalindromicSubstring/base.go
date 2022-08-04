package main

//
//func main() {
//
//}
//
//func longestPalindrome(s string) string {
//	if len(s) == 0 || len(s) == 1 {
//		return s
//	}
//
//	m := map[string]int{}
//	maxLength := 0
//	var maxAnsw string
//
//	_, dict := listOfPalindroms(s, m)
//
//	for i, _ := range dict {
//		if dict[i] > maxLength {
//			maxLength = dict[i]
//			maxAnsw = i
//		}
//	}
//
//	return maxAnsw
//}
//
//func listOfPalindroms(s string, m map[string]int) (string, map[string]int) {
//	start := 0
//	end := len(s) - 1
//	var answer string
//
//	if len(s) == 0 {
//		return "", m
//	}
//
//	if len(s) == 1 {
//		m[s] = 1
//		return s, m
//	}
//
//	if len(s) == 2 && s[0] == s[1] {
//		m[s] = 2
//		return s, m
//	}
//
//	if len(s) == 3 && s[0] == s[2] {
//		m[s] = 3
//		return s, m
//	}
//
//	m[string(s[0])] = 1
//
//	for ; start <= end; end-- {
//		if start == end {
//			answer = ""
//			end = len(s) - 1
//			start++
//		}
//		if s[start] == s[end] && start != end {
//			if _, ok := m[s[start+1:end]]; ok {
//				answer = string(s[start]) + s[start+1:end] + string(s[end])
//			} else if !ok {
//				if i, _ := listOfPalindroms(s[start+1:end], m); i != "" || s[start+1:end] == "" {
//					answer = string(s[start]) + i + string(s[end])
//				}
//			}
//
//			m[answer] = len(answer)
//		}
//		if _, ok := m[s]; ok {
//			break
//		}
//
//	}
//	return answer, m
//}
//*/
