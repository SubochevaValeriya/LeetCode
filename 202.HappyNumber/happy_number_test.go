package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	n        int
	Expected bool
}

func TestLeetCode(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with 19", n: 19, Expected: true},
		{Name: "case with 2", n: 2, Expected: false},
		{Name: "case with 7", n: 7, Expected: true},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := isHappy(cse.n)
			assert.Equalf(t, cse.Expected, result, "for %d expected %t, got %t", cse.n, cse.Expected, result)
		})
	}
}
