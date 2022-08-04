package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	InputS   []int
	Expected int
}

func TestWordBreak(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with [1,1,2]", InputS: []int{1, 1, 2}, Expected: 2},
		{Name: "case with [0,0,1,1,1,2,2,3,3,4]", InputS: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, Expected: 5},
		{Name: "case with [1,1]", InputS: []int{1, 1}, Expected: 1},
		{Name: "case with [1,2]", InputS: []int{1, 2}, Expected: 2},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := RemoveDuplicates(cse.InputS)
			assert.Equalf(t, cse.Expected, result, "for %d expected %t, got %t", cse.InputS, cse.Expected, result)
		})
	}
}
