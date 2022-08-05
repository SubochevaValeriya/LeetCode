package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	words    []string
	s        string
	Expected int
}

func TestWordBreak(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with abc", words: []string{"a", "b", "c", "ab", "bc", "abc"}, s: "abc", Expected: 3},
		{Name: "case with vomy", words: []string{"ycwj", "hak", "v", "kjg", "zw", "vtes", "vom", "ii", "as", "v", "vo", "v", "w", "vomy", "loa", "fbm", "j", "v", "vo", "e", "y", "t", "eh", "yk", "bt", "x", "vomy", "vom", "yab", "v", "sydi", "wnb", "z", "v", "iygp", "vomy"}, s: "vomy", Expected: 13},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := countPrefixes(cse.words, cse.s)
			assert.Equalf(t, cse.Expected, result, "for %d & %d expected %t, got %t", cse.words, cse.s, cse.Expected, result)
		})
	}
}
