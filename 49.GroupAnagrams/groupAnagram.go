package main

import "sort"

func GroupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	for _, str := range strs {
		s := []rune(str)
		sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
		m[string(s)] = append(m[string(s)], str)
	}

	answer := [][]string{}

	for _, strings := range m {
		answer = append(answer, strings)
	}

	return answer
}
