package main

import "strings"

// https://leetcode.com/problems/make-the-string-great/description/

func makeGood(s string) string {

	for i := 0; i < len(s)-1; i++ {
		if s[i] != s[i+1] && strings.ToLower(string(s[i])) == strings.ToLower(string(s[i+1])) {
			s = s[:i] + s[i+2:]
			i = -1
		}
	}

	return s
}
