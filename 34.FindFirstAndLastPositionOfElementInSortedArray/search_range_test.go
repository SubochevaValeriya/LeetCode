package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	nums     []int
	target   int
	Expected []int
}

func TestLeetCode(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with [5,7,7,8,8,10]", nums: []int{5, 7, 7, 8, 8, 10}, target: 8, Expected: []int{3, 4}},
		{Name: "case with [5,7,7,8,8,10]", nums: []int{5, 7, 7, 8, 8, 10}, target: 6, Expected: []int{-1, -1}},
		{Name: "case with []", nums: []int{}, target: 0, Expected: []int{-1, -1}},
		{Name: "case with [1]", nums: []int{1}, target: 1, Expected: []int{0, 0}},
		{Name: "case with [2,2]", nums: []int{2, 2}, target: 2, Expected: []int{0, 1}},
	}
	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := searchRange(cse.nums, cse.target)
			assert.Equalf(t, cse.Expected, result, "for %d expected %t, got %t", cse.nums, cse.Expected, result)
		})
	}
}
