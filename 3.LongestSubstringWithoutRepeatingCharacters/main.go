package main

import "fmt"

func main() {
	var s string
	s = "cdd"
	//fmt.Scanf(&s)
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {

	if s == "" {
		return 0
	}

	long := 1
	answer := 1
	runeS := []rune(s)
	runeM := map[rune]int{}

	//runeAnswer := []rune{}

	//for i := 1; i < len(runeS)-1; i++ {
	//	for j := 0; j < len(runeS)-1; j += i {
	//		for n := j + 1; n <= j+i; n++ {
	//			if runeS[j] == runeS[n] {
	//				long = answer
	//			}
	//		}
	//	}
	//	if long != answer {
	//		answer = i
	//	}
	//}

	for i := range runeS {
		for j := i + 1; j < len(runeS); j++ {
			if runeS[i] != runeS[j] {
				if _, ok := runeM[runeS[j]]; ok == true {
					break
				} else {
					long++
					runeM[runeS[j]] = 1
				}
			} else {
				break
			}
		}
		if long > answer {
			answer = long
		}
		long = 1
		runeM = map[rune]int{}
	}
	return answer
}
