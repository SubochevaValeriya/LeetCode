package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	s        string
	Expected string
}

func TestLeetCode(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with aabbcc", s: "aabbcc", Expected: "abcabc"},
		{Name: "case with vvvlo", s: "vvvlo", Expected: "vlvov"},
		{Name: "case with eqmeyggvp", s: "eqmeyggvp", Expected: "egmpqvyeg"},
		{Name: "case with hhhhhjjjjjfff", s: "hhhhhjjjjjfff", Expected: "fhjfhjfhjhjhj"},
		{Name: "case with hhhhhjjj", s: "hhhhjjj", Expected: ""},
		{Name: "case with aaab", s: "aaab", Expected: ""},
		{Name: "case with baabb", s: "baabb", Expected: "babab"},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := reorganizeStringInAlphabetOrder(cse.s)
			assert.Equalf(t, cse.Expected, result, "for %d expected %t, got %t", cse.s, cse.Expected, result)
		})
	}
}
