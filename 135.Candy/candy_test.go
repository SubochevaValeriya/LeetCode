package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	ratings  []int
	Expected int
}

func TestWordBreak(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with [1,0,2]", ratings: []int{1, 0, 2}, Expected: 5},
		{Name: "case with [1,2,2]", ratings: []int{1, 2, 2}, Expected: 4},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := candy(cse.ratings)
			assert.Equalf(t, cse.Expected, result, "for %d expected %t, got %t", cse.ratings, cse.Expected, result)
		})
	}
}
