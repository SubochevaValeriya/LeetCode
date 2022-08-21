package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	board    [][]string
	word     string
	Expected bool
}

func TestWordSearch(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with [0,1,0,3,2,3]", board: [][]string{{"A", "B", "C", "E"}, {"S", "F", "C", "S"}, {"A", "D", "E", "E"}}, word: "ABCEED", Expected: true},
		//{Name: "case with [7,7,7,7,7,7,7]", nums: []int{7, 7, 7, 7, 7, 7, 7}, Expected: 1},
		//{Name: "case with [-2, -1]", nums: []int{-2, -1}, Expected: 2},
		//{Name: "case with [1,3,6,7,9,4,10,5,6]", nums: []int{1, 3, 6, 7, 9, 4, 10, 5, 6}, Expected: 6},
		//	{Name: "case with [5,7,-24,12,13,2,3,12,5,6,35]", nums: []int{5, 7, -24, 12, 13, 2, 3, 12, 5, 6, 35}, Expected: 6},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := exist(cse.board, cse.Name)
			assert.Equalf(t, cse.Expected, result, "for %d expected %t, got %t", cse.board, cse.Name, cse.Expected, result)
		})
	}
}
