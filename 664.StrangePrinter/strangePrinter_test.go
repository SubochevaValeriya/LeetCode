package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	s        string
	Expected int
}

func TestLeetCode(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with aaabb", s: "aaabb", Expected: 2},
		{Name: "case with abcabc", s: "abcabc", Expected: 5},
		{Name: "case with tbgtgb", s: "tbgtgb", Expected: 4},
		{Name: "case with abcabcabc", s: "abcabcabc", Expected: 7},
		{Name: "case with baacdddaaddaaaaccbddbcabdaabdbbcdcbbbacbddcabcaaa", s: "baacdddaaddaaaaccbddbcabdaabdbbcdcbbbacbddcabcaaa", Expected: 19},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := strangePrinter(cse.s)
			assert.Equalf(t, cse.Expected, result, "for %d expected %t, got %t", cse.s, cse.Expected, result)
		})
	}
}
