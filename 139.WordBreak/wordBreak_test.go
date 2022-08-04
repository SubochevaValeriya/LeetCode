package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name      string
	InputS    string
	InputDict []string
	Expected  bool
}

func TestWordBreak(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with leetcode", InputS: "leetcode", InputDict: []string{"leet", "code"}, Expected: true},
		{Name: "case with applepenapple", InputS: "applepenapple", InputDict: []string{"apple", "pen"}, Expected: true},
		{Name: "case with catsandog", InputS: "catsandog", InputDict: []string{"cats", "dog", "sand", "and", "cat"}, Expected: false},
		{Name: "case with empty", InputS: "", InputDict: []string{""}, Expected: true},
		{Name: "case with cars", InputS: "cars", InputDict: []string{"car", "ca", "rs"}, Expected: true},
		{Name: "case with a", InputS: "a", InputDict: []string{"b"}, Expected: false},
		{Name: "case with b", InputS: "b", InputDict: []string{"c", "a"}, Expected: false},
		{Name: "case with aaaaaaa", InputS: "aaaaaaa", InputDict: []string{"aaaa", "aaa"}, Expected: true},
		{Name: "case with goalspecial", InputS: "goalspecial", InputDict: []string{"go", "goal", "goals", "special"}, Expected: true},
		{Name: "case with cat", InputS: "catdog", InputDict: []string{"cat", "catd", "jk"}, Expected: false},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := WordBreak(cse.InputS, cse.InputDict)
			assert.Equalf(t, cse.Expected, result, "for %d and %d expected %t, got %t", cse.InputS, cse.InputDict, cse.Expected, result)
		})
	}
}
