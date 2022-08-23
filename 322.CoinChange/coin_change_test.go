package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	coins    []int
	amount   int
	Expected int
}

func TestLeetCode(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with 11", coins: []int{1, 2, 5}, amount: 11, Expected: 3},
		{Name: "case with 3", coins: []int{2}, amount: 3, Expected: -1},
		{Name: "case with 0", coins: []int{1}, amount: 0, Expected: 0},
		{Name: "case with 6249", coins: []int{186, 419, 83, 408}, amount: 6249, Expected: 20},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := coinChange(cse.coins, cse.amount)
			assert.Equalf(t, cse.Expected, result, "for %d and %d expected %t, got %t", cse.coins, cse.amount, cse.Expected, result)
		})
	}
}
