package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	nums     []int
	Expected []int
}

func TestWordBreak(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with [1,2,3]", nums: []int{1, 2, 3}, Expected: []int{1, 3, 2}},
		{Name: "case with [3,2,1]", nums: []int{3, 2, 1}, Expected: []int{1, 2, 3}},
		{Name: "case with [1,1,5]", nums: []int{1, 1, 5}, Expected: []int{1, 5, 1}},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := nextPermutation(cse.nums)
			assert.Equalf(t, cse.Expected, result, "for %d expected %t, got %t", cse.nums, cse.Expected, result)
		})
	}
}
