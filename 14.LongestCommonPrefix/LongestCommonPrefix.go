package main

func LongestCommonPrefix2(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	var answer string
	maxLenght := 0

	firstWord := strs[0]

	for j := 1; j < len(strs); j++ {
		if maxLenght+1 <= len(strs[j]) && maxLenght+1 <= len(strs[0]) {
			if firstWord[:maxLenght+1] == strs[j][:maxLenght+1] {
				if j == len(strs)-1 {
					answer = firstWord[:maxLenght+1]
					maxLenght++
					j = 0
				}
			} else {
				break
			}
		} else {
			break
		}
	}

	return answer
}

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	var answer string
	maxLenght := 0

	firstWord := strs[0]

	for j := 1; j < len(strs); j++ {
		if maxLenght+1 > len(strs[j]) || maxLenght+1 > len(strs[0]) {
			return answer
		}

		if firstWord[:maxLenght+1] != strs[j][:maxLenght+1] {
			return answer
		}

		if j == len(strs)-1 {
			answer = firstWord[:maxLenght+1]
			maxLenght++
			j = 0
		}
	}

	return answer
}

//for i, word := range strs {
//	for i, letter := range word {
//		z
//		if string(letter) ==
//	}
//}
//
//for i, word := range strs {
//	for i, letter := range word {
//		if string(letter) ==
//	}
//}
