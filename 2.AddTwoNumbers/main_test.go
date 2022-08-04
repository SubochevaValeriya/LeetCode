package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	l1       *ListNode
	Expected int
}

func TestWordBreak(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with [1, 2, 2]", nums: []int{1, 2, 2}, Expected: 2},
		{Name: "case with [3, 1, 3, 4, 2]", nums: []int{3, 1, 3, 4, 2}, Expected: 3},
		{Name: "case with [1,3,4,2,2]", nums: []int{1, 3, 4, 2, 2}, Expected: 2},
		{Name: "case with [2,2,2,2,2]", nums: []int{2, 2, 2, 2, 2}, Expected: 2},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := FindDuplicate(cse.nums)
			assert.Equalf(t, cse.Expected, result, "for %d expected %t, got %t", cse.nums, cse.Expected, result)
		})
	}
}
