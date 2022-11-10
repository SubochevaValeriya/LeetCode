package main

func characterReplacement(s string, k int) int {
	maxLenght := k

	for i := 0; i < len(s); i++ {
		if i+maxLenght+1 < len(s) {
			if s[i] == s[i+maxLenght+1] {
				maxLenght = i + maxLenght + 1 - i
				for j := i + maxLenght + 1; j < len(s); j++ {
					if s[i] == s[j] {
						maxLenght = j - i
					}
				}
			}
		}
	}

	return maxLenght
}
