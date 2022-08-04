package main

func StrStr(haystack string, needle string) int {
	end := len(needle)

	for start := 0; start < len(haystack)-end; start++ {
		if haystack[start:start+end] == needle {
			return start
		}
	}

	return -1
}
