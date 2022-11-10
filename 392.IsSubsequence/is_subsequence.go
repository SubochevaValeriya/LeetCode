package main

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}

	first := 0
	n := 0
	for i := 0; i < len(t); i++ {
		if n == 0 {
			first = i
		}
		if t[i] == s[n] {
			if n == len(s)-1 {
				return true
			}
			n++
		} else if first != len(t)-len(s) && i == len(t)-1 {
			n = 0
			i = first
		}
	}
	return false
}
