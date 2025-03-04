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
		{Name: "case with [1,6,10,8,7,3,2]", ratings: []int{1, 6, 10, 8, 7, 3, 2}, Expected: 18},
		{Name: "case with [1,3,4,5,2]", ratings: []int{1, 3, 4, 5, 2}, Expected: 11},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := candy(cse.ratings)
			assert.Equalf(t, cse.Expected, result, "for %d expected %t, got %t", cse.ratings, cse.Expected, result)
		})
	}
}

//0 0 1 1 1 1 0
//0 1 1 0 0 0 0
//
//1 2 2 2 2 2 1
//1 2 5 4 3 2 1
//
//1 6 10 8
//
//1 2 2 1
//
//
//
//1 2 3 2
//
//8 4 6
//
//7
//0 1 2 1 1 1 0
//
//1 2 4 3 2 2 1
//
//
//1 6 2 1
//
//1 3 2
//
//1 2 2
