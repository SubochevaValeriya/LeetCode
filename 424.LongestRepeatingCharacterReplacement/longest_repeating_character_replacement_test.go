package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	s        string
	k        int
	Expected int
}

func TestLeetCode(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with ABAB", s: "ABAB", k: 2, Expected: 4},
		{Name: "case with AABABBA", s: "AABABBA", k: 1, Expected: 4},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := characterReplacement(cse.s, cse.k)
			assert.Equalf(t, cse.Expected, result, "for %s & %d expected %d, got %d", cse.s, cse.k, cse.Expected, result)
		})
	}
}
