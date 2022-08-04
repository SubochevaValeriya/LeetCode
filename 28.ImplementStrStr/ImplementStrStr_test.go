package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	haystack string
	needle   string
	Expected int
}

func TestWordBreak(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with hello", haystack: "hello", needle: "ll", Expected: 2},
		{Name: "case with aaaaa", haystack: "aaaaa", needle: "bba", Expected: -1},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := StrStr(cse.haystack, cse.needle)
			assert.Equalf(t, cse.Expected, result, "for %d and %d expected %t, got %t", cse.haystack, cse.needle, cse.Expected, result)
		})
	}
}
