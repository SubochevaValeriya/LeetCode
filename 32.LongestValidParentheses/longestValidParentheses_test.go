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

func TestWordBreak(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with (()", s: "(()", Expected: 2},
		{Name: "case with )()())", s: ")()())", Expected: 4},
		{Name: "case with ()((()((()", s: "()()(((()((", Expected: 4},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := longestValidParentheses(cse.s)
			assert.Equalf(t, cse.Expected, result, "for %d expected %t, got %t", cse.s, cse.Expected, result)
		})
	}
}
