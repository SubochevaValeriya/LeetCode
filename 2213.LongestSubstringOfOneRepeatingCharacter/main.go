package main

import "math"

func main() {

}

type substringsInfo struct {
	substrings   [][2]int
	maxLength    int
	maxSubstring [2]int
}

func longestRepeating(s string, queryCharacters string, queryIndices []int) []int {

	letterSubstrings := map[rune]substringsInfo{}
	sRune := []rune(s)
	var start, end = 0, 0
	maxLength := 0
	maxSubstring := [2]int{}

	for i, letter := range sRune {
		if i == 0 {
			continue
		}
		if letter != sRune[i-1] {
			entry, ok := letterSubstrings[sRune[i-1]]
			if !ok {
				entry = substringsInfo{
					substrings:   [][2]int{},
					maxLength:    0,
					maxSubstring: [2]int{},
				}
			}
			entry.substrings = append(letterSubstrings[sRune[i-1]].substrings, [2]int{start, end})
			if end-start+1 > entry.maxLength {
				entry.maxLength = end - start + 1
				entry.maxSubstring = [2]int{start, end}
				if entry.maxLength > maxLength {
					maxLength = entry.maxLength
					maxSubstring = entry.maxSubstring
				}
			}
			letterSubstrings[sRune[i-1]] = entry
			start, end = i, i
		} else {
			end++
		}
		if i == len(s)-1 {
			entry, ok := letterSubstrings[sRune[i]]
			if !ok {
				entry = substringsInfo{
					substrings:   [][2]int{},
					maxLength:    0,
					maxSubstring: [2]int{},
				}
			}
			entry.substrings = append(letterSubstrings[sRune[i]].substrings, [2]int{start, end})
			if end-start+1 > entry.maxLength {
				entry.maxLength = end - start + 1
				entry.maxSubstring = [2]int{start, end}
				if entry.maxLength > maxLength {
					maxLength = entry.maxLength
					maxSubstring = entry.maxSubstring
				}
			}
			letterSubstrings[sRune[i]] = entry
		}
	}

	answer := make([]int, len(queryCharacters))

	for i, character := range queryCharacters {
		if sRune[queryIndices[i]] == character {
			answer[i] = maxLength
			continue
		}
		//sRune[i] = character
		entry, ok := letterSubstrings[character]
		if !ok {
			entry = substringsInfo{
				substrings:   [][2]int{{queryIndices[i],queryIndices[i]}},
				maxLength:    1,
				maxSubstring: [2]int{queryIndices[i],queryIndices[i]},
			}
			letterSubstrings[character] = entry
		} else{
		for j, substring := range entry.substrings {
			var needToRecalculateMax bool
			if substring[0] < queryIndices[i] && substring[1] > queryIndices[i] {
				if letterSubstrings[character].maxSubstring == substring {
					needToRecalculateMax = true
				}


			}

			abc d b d
			abc c b c
			cbc c b c
			if substring[1] < queryIndices[i] {
				continue
			}
		}
		}
		for i, character := range letterSubstrings[sRune[queryIndices[i]-1]].substrings {

		}
	}
}
