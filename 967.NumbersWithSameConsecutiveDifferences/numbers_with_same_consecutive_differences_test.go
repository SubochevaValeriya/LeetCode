package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	InputN   int
	InputK   int
	Expected []int
}

func TestLeetCode(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with n=3, k=7", InputN: 3, InputK: 7, Expected: []int{181, 292, 707, 818, 929}},
		{Name: "case with n=2, k=1", InputN: 2, InputK: 1, Expected: []int{10, 12, 21, 23, 32, 34, 43, 45, 54, 56, 65, 67, 76, 78, 87, 89, 98}},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := numsSameConsecDiff(cse.InputN, cse.InputK)
			assert.Equalf(t, cse.Expected, result, "for %d and %d expected %t, got %t", cse.InputN, cse.InputK, cse.Expected, result)
		})
	}
}
