package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	height   []int
	Expected int
}

func TestWordBreak(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with [1,8,6,2,5,4,8,3,7]", height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, Expected: 49},
		{Name: "case with [1,1]", height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, Expected: 49},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := MaxArea(cse.height)
			assert.Equalf(t, cse.Expected, result, "for %d expected %t, got %t", cse.height, cse.Expected, result)
		})
	}
}
