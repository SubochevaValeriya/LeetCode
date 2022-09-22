package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	s        string
	Expected string
}

func TestLeetCode(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with Let's take LeetCode contest", s: "Let's take LeetCode contest", Expected: "s'teL ekat edoCteeL tsetnoc"},
		{Name: "case with God Ding", s: "God Ding", Expected: "doG gniD"},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := reverseWords(cse.s)
			assert.Equalf(t, cse.Expected, result, "for %d expected %t, got %t", cse.s, cse.Expected, result)
		})
	}
}
