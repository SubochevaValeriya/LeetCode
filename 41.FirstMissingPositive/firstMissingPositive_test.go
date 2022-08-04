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

func TestWordBreak(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with [1,2,0]", nums: []int{1, 2, 0}, Expected: 3},
		{Name: "case with [3,4,-1,1]", nums: []int{3, 4, -1, 1}, Expected: 2},
		{Name: "case with [7,8,9,11,12]", nums: []int{7, 8, 9, 11, 12}, Expected: 1},
		{Name: "case with [1,2,6,3,5,4]", nums: []int{1, 2, 6, 3, 5, 4}, Expected: 7},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := firstMissingPositive(cse.nums)
			assert.Equalf(t, cse.Expected, result, "for %d expected %t, got %t", cse.nums, cse.Expected, result)
		})
	}
}
