package main

//
//func longestPalindrome2(s string) string {
//	if len(s) == 0 || len(s) == 1 {
//		return s
//	}
//
//	m := map[string]int{}
//	_, maxAnsw := listOfPalindromss(s, m)
//	return maxAnsw
//}
//
//func listOfPalindromss(s string, m map[string]int) (string, string) {
//	start := 0
//	end := len(s) - 1
//	var answer string
//	maxLength := 0
//	var maxAnsw string
//
//	if len(s) == 0 {
//		return "", ""
//	}
//
//	if len(s) == 1 {
//		m[s] = 1
//		return s, s
//	}
//
//	if len(s) == 2 && s[0] == s[1] {
//		m[s] = 2
//		return s, s
//	}
//
//	if len(s) == 3 && s[0] == s[2] {
//		m[s] = 3
//		return s, s
//	}
//
//	m[string(s[0])] = 1
//	maxLength = 1
//	maxAnsw = string(s[0])
//
//	for ; start <= end; end-- {
//		if _, ok := m[s]; ok {
//			break
//		}
//
//		if s[start] != s[end] {
//			answer = ""
//		}
//
//		if start == end && len(s)-start > maxLength {
//			end = len(s) - 1
//			start++
//		}
//
//		if s[start] == s[end] && start != end {
//			m[string(s[start])+string(s[end])] = 2
//			if len(s[start:end+1]) < 4 {
//				answer, _ = listOfPalindroms(s[start:end+1], m)
//			} else {
//				if _, ok := m[s[start+1:end]]; ok {
//					answer = string(s[start]) + s[start+1:end] + string(s[end])
//				} else if !ok {
//					if i, _ := listOfPalindroms(s[start+1:end], m); i != "" || s[start+1:end] == "" {
//						answer = string(s[start]) + i + string(s[end])
//					}
//				}
//			}
//			m[answer] = len(answer)
//
//			if m[answer] > maxLength {
//				maxAnsw = answer
//				maxLength = m[answer]
//			}
//		}
//	}
//
//	return answer, maxAnsw
//}
//
//func listOfPalindroms(s string, m map[string]int) (string, string) {
//	start := 0
//	end := len(s) - 1
//	var answer string
//	maxLength := 0
//	var maxAnsw string
//
//	if len(s) == 0 {
//		return "", ""
//	}
//
//	if len(s) == 1 {
//		m[s] = 1
//		return s, s
//	}
//
//	if len(s) == 2 && s[0] == s[1] {
//		m[s] = 2
//		return s, s
//	}
//
//	if len(s) == 3 && s[0] == s[2] {
//		m[s] = 3
//		return s, s
//	}
//
//	m[string(s[0])] = 1
//	maxLength = 1
//	maxAnsw = string(s[0])
//
//	for ; start <= end; end-- {
//		if _, ok := m[s]; ok {
//			break
//		}
//
//		if s[start] != s[end] {
//			answer = ""
//			break
//		}
//
//		if start == end && len(s)-start > maxLength {
//			end = len(s) - 1
//			start++
//		}
//
//		if s[start] == s[end] && start != end {
//			m[string(s[start])+string(s[end])] = 2
//			if len(s[start:end+1]) < 4 {
//				answer, _ = listOfPalindroms(s[start:end+1], m)
//			} else {
//				if _, ok := m[s[start+1:end]]; ok {
//					answer = string(s[start]) + s[start+1:end] + string(s[end])
//				} else if !ok {
//					if i, _ := listOfPalindroms(s[start+1:end], m); i != "" || s[start+1:end] == "" {
//						answer = string(s[start]) + i + string(s[end])
//					}
//				}
//			}
//			m[answer] = len(answer)
//
//			if m[answer] > maxLength {
//				maxAnsw = answer
//				maxLength = m[answer]
//			}
//		}
//	}
//
//	return answer, maxAnsw
//}
