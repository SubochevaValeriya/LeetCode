package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	strs     []string
	Expected [][]string
}

func TestWordBreak(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with [\"eat\",\"tea\",\"tan\",\"ate\",\"nat\",\"bat\"]", strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}, Expected: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := GroupAnagrams(cse.strs)
			assert.Equalf(t, cse.Expected, result, "for %d expected %t, got %t", cse.strs, cse.Expected, result)
		})
	}
}
