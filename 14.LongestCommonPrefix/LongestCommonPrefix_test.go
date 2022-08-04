package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	InputS   []string
	Expected string
}

func TestWordBreak(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with fl", InputS: []string{"flower", "flow", "flight"}, Expected: "fl"},
		{Name: "case with empty", InputS: []string{"dog", "racecar", "car"}, Expected: ""},
		{Name: "case with a", InputS: []string{"a"}, Expected: "a"},
		{Name: "case with b", InputS: []string{"", "b"}, Expected: ""},
		{Name: "case with aa", InputS: []string{"aa", "aa"}, Expected: "aa"},
		{Name: "case with aaa", InputS: []string{"aaa", "aa", "aaa"}, Expected: "aa"},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := LongestCommonPrefix(cse.InputS)
			assert.Equalf(t, cse.Expected, result, "for %d expected %t, got %t", cse.InputS, cse.Expected, result)
		})
	}
}
