// https://leetcode.com/problems/minimum-number-of-moves-to-make-palindrome/
package main

func minMovesToMakePalindrome(s string) int {

	if len(s) == 1 {
		return 0
	}
	sRune := []rune(s)
	left, right := 0, len(s)-1
	minSwap := len(s)
	minSwapI, minSwapJ := left, right
	if sRune[left] != sRune[right] {
		for i := 0; i < len(s)-1; i++ {
			if i >= minSwap {
				break
			}
			for j := len(s) - 1; j > 0; j-- {
				if i+right-j >= minSwap {
					break
				}
				if sRune[i] == sRune[j] {
					minSwap = i + right - j
					minSwapI = i
					minSwapJ = j
				}
			}
		}
	}
	for i := minSwapI; i > left; i-- {
		sRune[i-1], sRune[i] = sRune[i], sRune[i-1]
	}

	for i := minSwapJ; i < right; i++ {
		sRune[i+1], sRune[i] = sRune[i], sRune[i+1]
	}

	if minSwap == len(s) {
		minSwap = 0
	}

	if len(sRune) == 4 {
		return minSwap
	}
	return minSwap + minMovesToMakePalindrome(string(sRune[1:len(s)-1]))
}
