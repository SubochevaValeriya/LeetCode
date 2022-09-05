package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name      string
	candies   int
	numPeople int
	Expected  []int
}

func TestWordBreak(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with 7, 4", candies: 7, numPeople: 4, Expected: []int{1, 2, 3, 1}},
		{Name: "case with 10, 3", candies: 10, numPeople: 3, Expected: []int{5, 2, 3}},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := distributeCandies(cse.candies, cse.numPeople)
			assert.Equalf(t, cse.Expected, result, "for %d & %d expected %t, got %t", cse.candies, cse.numPeople, cse.Expected, result)
		})
	}
}
