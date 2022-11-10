package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	s        string
	t        string
	Expected bool
}

func TestLeetCode(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with ahbgdc", s: "abc", t: "ahbgdc", Expected: true},
		{Name: "case with ahbgdc", s: "axc", t: "ahbgdc", Expected: false},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := isSubsequence(cse.s, cse.t)
			assert.Equalf(t, cse.Expected, result, "for %s & %s expected %t, got %t", cse.s, cse.t, cse.Expected, result)
		})
	}
}
