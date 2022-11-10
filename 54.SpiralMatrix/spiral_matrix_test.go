package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	matrix   [][]int
	Expected []int
}

func TestLeetCode(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with [[1,2,3],[4,5,6],[7,8,9]]", matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, Expected: []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		//		{Name: "case with [5,5,10,10,20]", bills: []int{5, 5, 10, 10, 20}, Expected: false},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := spiralOrder(cse.matrix)
			assert.Equalf(t, cse.Expected, result, "for %d expected %t, got %t", cse.matrix, cse.Expected, result)
		})
	}
}
