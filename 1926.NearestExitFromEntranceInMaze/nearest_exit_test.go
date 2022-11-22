package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	maze     [][]string
	entrance []int
	Expected int
}

func TestLeetCode(t *testing.T) {
	testCases := []TestCase{
		{Name: "case 1", maze: [][]string{{"+", "+", ".", "+"}, {".", ".", ".", "+"}, {"+", "+", "+", "."}}, entrance: []int{1, 2}, Expected: 1},
		{Name: "case 2", maze: [][]string{{"+", "+", "+"}, {".", ".", "."}, {"+", "+", "+"}}, entrance: []int{1, 0}, Expected: 2},
		{Name: "case 3", maze: [][]string{{".", "+"}}, entrance: []int{0, 0}, Expected: -1},
		{Name: "case 4", maze: [][]string{{".", "."}}, entrance: []int{0, 1}, Expected: 1},
		{Name: "case 5", maze: [][]string{{"+", ".", "+", "+", "+", "+", "+"}, {"+", ".", "+", ".", ".", ".", "+"}, {"+", ".", "+", ".", "+", ".", "+"}, {"+", ".", ".", ".", "+", ".", "+"}, {"+", "+", "+", "+", "+", ".", "+"}}, entrance: []int{0, 1}, Expected: 12},
		{Name: "case 6", maze: [][]string{{"+", ".", "+", "+", "+", "+", "+"}, {"+", ".", "+", ".", ".", ".", "+"}, {"+", ".", "+", ".", "+", ".", "+"}, {"+", ".", ".", ".", "+", ".", "+"}, {"+", "+", "+", "+", "+", ".", "+"}}, entrance: []int{3, 2}, Expected: 4},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := nearestExitTest(cse.maze, cse.entrance)
			assert.Equalf(t, cse.Expected, result, "for %d expected %t, got %t", cse.maze, cse.Expected, result)
		})
	}
}
