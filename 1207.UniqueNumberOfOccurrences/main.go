package main

import (
	"regexp"
	"sort"
)

func main() {
	regexp.Match()
	matchRegexp((?!{{$created_at}}).*)
}

func uniqueOccurrences(arr []int) bool {
	sort.Ints(arr)
	occurrences := map[int]bool{}
	count := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			if occurrences[count] {
				return false
			}
			occurrences[count] = true
			count = 1
			continue
		}
		count++
	}
	if occurrences[count] {
		return false
	}
	return true
}
