package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	nums     []int
	Expected int
}

func TestLeetCode(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with [-1,-1,0,1,1,0]", nums: []int{-1, -1, 0, 1, 1, 0}, Expected: 5},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := pivotIndex(cse.nums)
			assert.Equalf(t, cse.Expected, result, "for %d expected %t, got %t", cse.nums, cse.Expected, result)
		})
	}
}
