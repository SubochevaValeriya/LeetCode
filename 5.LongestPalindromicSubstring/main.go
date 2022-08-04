package main

import (
	"fmt"
)

func main() {
	//s := "aacabdkacaa"
	//s := "bb"
	//s := "ivilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreat"
	//s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabcaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	//s := "ababababa"
	//s := "dabbcd"
	//s := "jrjnbctoqgzimtoklkxcknwmhiztomaofwwzjnhrijwkgmwwuazcowskjhitejnvtblqyepxispasrgvgzqlvrmvhxusiqqzzibcyhpnruhrgbzsmlsuacwptmzxuewnjzmwxbdzqyvsjzxiecsnkdibudtvthzlizralpaowsbakzconeuwwpsqynaxqmgngzpovauxsqgypinywwtmekzhhlzaeatbzryreuttgwfqmmpeywtvpssznkwhzuqewuqtfuflttjcxrhwexvtxjihunpywerkktbvlsyomkxuwrqqmbmzjbfytdddnkasmdyukawrzrnhdmaefzltddipcrhuchvdcoegamlfifzistnplqabtazunlelslicrkuuhosoyduhootlwsbtxautewkvnvlbtixkmxhngidxecehslqjpcdrtlqswmyghmwlttjecvbueswsixoxmymcepbmuwtzanmvujmalyghzkvtoxynyusbpzpolaplsgrunpfgdbbtvtkahqmmlbxzcfznvhxsiytlsxmmtqiudyjlnbkzvtbqdsknsrknsykqzucevgmmcoanilsyyklpbxqosoquolvytefhvozwtwcrmbnyijbammlzrgalrymyfpysbqpjwzirsfknnyseiujadovngogvptphuyzkrwgjqwdhtvgxnmxuheofplizpxijfytfabx"
	s := "reifadyqgztixemwswtccodfnchcovrmiooffbbijkecuvlvukecutasfxqcqygltrogrdxlrslbnzktlanycgtniprjlospzhhgdrqcwlukbpsrumxguskubokxcmswjnssbkutdhppsdckuckcbwbxpmcmdicfjxaanoxndlfpqwneytatcbyjmimyawevmgirunvmdvxwdjbiqszwhfhjmrpexfwrbzkipxfowcbqjckaotmmgkrbjvhihgwuszdrdiijkgjoljjdubcbowvxslctleblfmdzmvdkqdxtiylabrwaccikkpnpsgcotxoggdydqnuogmxttcycjorzrtwtcchxrbbknfmxnonbhgbjjypqhbftceduxgrnaswtbytrhuiqnxkivevhprcvhggugrmmxolvfzwadlnzdwbtqbaveoongezoymdrhywxcxvggsewsxckucmncbrljskgsgtehortuvbtrsfisyewchxlmxqccoplhlzwutoqoctgfnrzhqctxaqacmirrqdwsbdpqttmyrmxxawgtjzqjgffqwlxqxwxrkgtzqkgdulbxmfcvxcwoswystiyittdjaqvaijwscqobqlhskhvoktksvmguzfankdigqlegrxxqpoitdtykfltohnzrcgmlnhddcfmawiriiiblwrttveedkxzzagdzpwvriuctvtrvdpqzcdnrkgcnpwjlraaaaskgguxzljktqvzzmruqqslutiipladbcxdwxhmvevsjrdkhdpxcyjkidkoznuagshnvccnkyeflpyjzlcbmhbytxnfzcrnmkyknbmtzwtaceajmnuyjblmdlbjdjxctvqcoqkbaszvrqvjgzdqpvmucerumskjrwhywjkwgligkectzboqbanrsvynxscpxqxtqhthdytfvhzjdcxgckvgfbldsfzxqdozxicrwqyprgnadfxsionkzzegmeynye"
	//	fmt.Println(longestPalindrome(s))
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}
	var answer, maxAnsw string
	var maxLenght int
	m := map[string]int{}

	m[string(s[0])] = 1
	maxAnsw = string(s[0])
	maxLenght = 1

	for i := 0; i < len(s); i++ {
		for j := i + maxLenght; j < len(s); j++ {
			if s[i] != s[j] {
				continue
			}

			if _, ok := m[s[i+1:j]]; ok || Reverse(s[i:j+1]) {
				m[s[i:j+1]] = len(s[i : j+1])
				answer = s[i : j+1]
				//} else {
				//	if Reverse(s[i : j+1]) {
				//		m[s[i:j+1]] = len(s[i : j+1])
				//		answer = s[i : j+1]
				//	}
			}

			if len(answer) > maxLenght {
				maxLenght = len(answer)
				maxAnsw = answer
			}
		}
	}

	return maxAnsw
}

func Reverse(s string) bool {
	j := len(s) - 1
	for i := 0; i < len(s); i++ {
		if s[i] != s[j] {
			return false
		}
		j--
	}
	return true
}

//
//func longestPalindrome(s string) string {
//	m := map[string]int{}
//
//	start := 0
//	end := len(s) - 1
//	maxLength := 0
//	var maxAnsw string
//	var answer string
//
//	if len(s) == 1 {
//		m[s] = 1
//		return s
//	}
//
//	//if len(s) == 2 && s[0] == s[1] {
//	//	m[s] = 2
//	//	return s
//	//}
//	//
//	//if len(s) == 3 && s[0] == s[2] {
//	//	m[s] = 3
//	//	return s
//	//}
//
//	for ; start <= end; end-- {
//		if s[start] != s[end] {
//			answer = ""
//		}
//		if start == end && len(s)-1-start > maxLength {
//			end = len(s) - 1
//			start++
//		}
//
//		if s[start] == s[end] && start != end {
//			if len(s[start:end]) < 4 {
//				answer = longestPalindrome(s[start:end])
//			} else {
//				if _, ok := m[s[start+1:end]]; ok {
//					answer = string(s[start]) + s[start+1:end] + string(s[end])
//				} else if !ok {
//				} else {
//					if longestPalindrome(s[start+1:end]) != "" || s[start+1:end] == "" {
//						answer = string(s[start]) + longestPalindrome(s[start+1:end]) + string(s[end])
//					}
//				}
//			}
//		}
//	}
//
//	m[answer] = len(answer)
//	if m[answer] > maxLength {
//		maxAnsw = answer
//		maxLength = m[answer]
//
//	}
//
//	fmt.Println(maxAnsw)
//	return answer
//}

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

//
//func longestPalindrome(s string) string {
//	m := map[string]int{}
//
//	start := 0
//	end := len(s) - 1
//	//var lenghtOfSubstring int
//	maxLength := 0
//	var maxAnsw string
//	var answer string
//
//	if len(s) == 1 {
//		return s
//	}
//
//	if len(s) == 2 && s[0] == s[1] {
//		return s
//	} else if len(s) == 2 && s[0] != s[1] {
//		return string(s[0])
//	}
//
//	for ; start < len(s); start++ {
//		if s[start] == s[end] && start != end {
//			if s[start] == s[end] && start != end {
//				if _, ok := m[s[start+1:end]]; ok {
//					answer = string(s[start]) + longestPalindrome(s[start+1:end]) + string(s[end])
//					m[s] = len(s)
//					break
//				} else {
//					if longestPalindrome(s[start+1:end]) != "" {
//						answer = string(s[start]) + longestPalindrome(s[start+1:end]) + string(s[end])
//						m[answer] = len(answer)
//						if m[answer] > maxLength {
//							maxAnsw = answer
//							maxLength = m[answer]
//						}
//						break
//					}
//				}
//			}
//		} else {
//			answer = ""
//			for end > start {
//				if s[start] == s[end] && start != end {
//					if _, ok := m[s[start+1:end]]; ok {
//						answer = string(s[start]) + longestPalindrome(s[start+1:end]) + string(s[end])
//						m[s] = len(s)
//						break
//					} else {
//						if longestPalindrome(s[start+1:end]) != "" {
//							answer = string(s[start]) + longestPalindrome(s[start+1:end]) + string(s[end])
//							m[answer] = len(answer)
//							if m[answer] > maxLength {
//								maxAnsw = answer
//								maxLength = m[answer]
//							}
//							break
//						}
//					}
//				} else {
//					end--
//				}
//			}
//			end = len(s) - 1
//		}
//	}
//
//	fmt.Println(answer)
//	return maxAnsw
//}

//func longestPalindrome(s string) string {
//	start := 0
//	end := len(s) - 1
//	lenghtOfSubstring := 0
//	maxLength := 0
//	var answer string
//
//
//
//
//	for start <= end {
//		if start == end {
//			lenghtOfSubstring++
//			answer += string(s[start])
//			if maxLength < lenghtOfSubstring {
//				maxLength = lenghtOfSubstring
//			}
//			break
//		} else {
//			if s[start] == s[end] {
//				lenghtOfSubstring++
//				answer += string(s[start])
//				start++
//				end--
//			} else {
//				lenghtOfSubstring = 0
//				end--
//			}
//		}
//	}
//	return answer
//}

//func longestPalindrome(s string) string {
//
//
//	lenghtOfSubstring := 0
//	maxLength := 0
//	var answer string
//
//	for start := 0; start < len(s); start++ {
//		for end := len(s) - 1;
//
//		if start == end {
//			lenghtOfSubstring++
//			answer += string(s[start])
//			if maxLength < lenghtOfSubstring {
//				maxLength = lenghtOfSubstring
//			}
//		} else {
//			if s[start] == s[end] {
//				end--
//				lenghtOfSubstring++
//				answer += string(s[start])
//			} else {
//				lenghtOfSubstring = 0
//				end--
//			}
//		}
//	}
//
//	return answer
//}

//func longestPalindrome(s string) string {
//
//	lenghtOfSubstring := 0
//	maxLength := 0
//	var answer string
//
//	for start := 0; start < len(s); start++ {
//		for end := len(s) - 1; end > 0; end-- {
//			if start == end {
//				lenghtOfSubstring++
//				answer += string(s[start])
//				if maxLength < lenghtOfSubstring {
//					maxLength = lenghtOfSubstring
//				}
//			} else {
//				if s[start] == s[end] {
//					lenghtOfSubstring++
//					answer += string(s[start])
//				} else {
//					lenghtOfSubstring = 0
//				}
//			}
//		}
//	}
//
//	return answer
//}
