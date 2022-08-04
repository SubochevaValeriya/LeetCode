package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name      string
	InputNums []int
	InputK    int
	InputT    int
	Expected  bool
}

func TestContainsDuplicates(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with 1,2,5,6,7,2,4", InputNums: []int{1, 2, 5, 6, 7, 2, 4}, InputK: 4, InputT: 0, Expected: true},
		{Name: "case with 1,2,5,6,7,2,4", InputNums: []int{1, 2, 5, 6, 7, 2, 4}, InputK: 4, InputT: 0, Expected: true},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := ContainsNearbyAlmostDuplicate(cse.InputNums, cse.InputK, cse.InputT)
			assert.Equalf(t, cse.Expected, result, "for %d and %d and %d expected %t, got %t", cse.InputNums, cse.InputK, cse.InputT, cse.Expected, result)
		})
	}
}
