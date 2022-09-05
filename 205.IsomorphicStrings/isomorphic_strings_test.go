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
		{Name: "case with egg & add", s: "egg", t: "add", Expected: true},
		{Name: "case with foo & bar ", s: "foo", t: "bar", Expected: false},
		{Name: "case with paper & title", s: "paper", t: "title", Expected: true},
		{Name: "case with badc & baba", s: "badc", t: "baba", Expected: false},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := isIsomorphic(cse.s, cse.t)
			assert.Equalf(t, cse.Expected, result, "for %d and %d expected %t, got %t", cse.s, cse.t, cse.Expected, result)
		})
	}
}
