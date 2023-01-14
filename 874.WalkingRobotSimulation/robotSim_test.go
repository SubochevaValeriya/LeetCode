package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name      string
	commands  []int
	obstacles [][]int
	Expected  int
}

func TestLeetCode(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with [4, -1, 3]", commands: []int{4, -1, 3}, obstacles: [][]int{}, Expected: 25},
		{Name: "case with [4, -1, 4, -2, 4]", commands: []int{4, -1, 4, -2, 4}, obstacles: [][]int{{2, 4}}, Expected: 65},
		{Name: "case with [6, -1, -1, 6]", commands: []int{6, -1, -1, 6}, obstacles: [][]int{}, Expected: 36},
		{Name: "case with [-2,-1,8,9,6]", commands: []int{-2, -1, 8, 9, 6}, obstacles: [][]int{{-1, 3}, {0, 1}, {-1, 5}, {-2, -4}, {5, 4}, {-2, -3}, {5, -1}, {1, -1}, {5, 5}, {5, 2}}, Expected: 0},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := robotSim(cse.commands, cse.obstacles)
			assert.Equalf(t, cse.Expected, result, "for %d expected %t, got %t", cse.commands, cse.Expected, result)
		})
	}
}
