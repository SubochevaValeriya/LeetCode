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

func TestWordBreak(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with level", s: "level", Expected: "l"},
		{Name: "case with ababab", s: "ababab", Expected: "abab"},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := longestPrefix(cse.s)
			assert.Equalf(t, cse.Expected, result, "for %d expected %t, got %t", cse.s, cse.Expected, result)
		})
	}
}
