package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	prices   []int
	Expected int
}

func TestWordBreak(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with [3,3,5,0,0,3,1,4]", prices: []int{3, 3, 5, 0, 0, 3, 1, 4}, Expected: 6},
		{Name: "case with [3,2,6,5,0,3]", prices: []int{3, 2, 6, 5, 0, 3}, Expected: 7},
		{Name: "case with [2,1,4,5,2,9,7]", prices: []int{2, 1, 4, 5, 2, 9, 7}, Expected: 11},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := maxProfit(cse.prices)
			assert.Equalf(t, cse.Expected, result, "for %d expected %t, got %t", cse.prices, cse.Expected, result)
		})
	}
}
