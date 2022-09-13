package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	bills    []int
	Expected bool
}

func TestLeetCode(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with [5,5,5,10,20]", bills: []int{5, 5, 5, 10, 20}, Expected: true},
		{Name: "case with [5,5,10,10,20]", bills: []int{5, 5, 10, 10, 20}, Expected: false},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := lemonadeChange(cse.bills)
			assert.Equalf(t, cse.Expected, result, "for %d expected %t, got %t", cse.bills, cse.Expected, result)
		})
	}
}
