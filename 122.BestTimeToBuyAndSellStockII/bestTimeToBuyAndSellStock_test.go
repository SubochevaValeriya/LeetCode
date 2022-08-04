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
		{Name: "case with 7,1,5,3,6,4", prices: []int{7, 1, 5, 3, 6, 4}, Expected: 7},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := maxProfit(cse.prices)
			assert.Equalf(t, cse.Expected, result, "for %d expected %t, got %t", cse.prices, cse.Expected, result)
		})
	}
}
