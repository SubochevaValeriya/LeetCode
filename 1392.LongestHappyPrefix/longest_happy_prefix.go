package main

func longestPrefix(s string) string {
	endPrfx, startSfx := len(s)-2, 1 // end/start of potential longest prefix & suffix

	for endPrfx >= 0 {
		if s[:endPrfx+1] == s[startSfx:] {
			return s[:endPrfx+1]
		}
		endPrfx--
		startSfx++

	}
	return ""
}

func longestPrefix2(s string) string {
	left, right := 0, len(s)-1
	var answer string

	for right > 0 {
		if s[:left+1] == s[right:] {
			answer = s[:left+1]
		}
		left++
		right--

	}
	return answer
}
