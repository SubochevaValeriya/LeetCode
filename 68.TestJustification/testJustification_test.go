package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	words    []string
	maxWidth int
	Expected []string
}

func TestWordBreak(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with [\"This\", \"is\", \"an\", \"example\", \"of\", \"text\", \"justification.\"]",
			words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
			maxWidth: 16,
			Expected: []string{"This    is    an", "example  of text", "justification.  "}},
		{Name: "case with [\"What\",\"must\",\"be\",\"acknowledgment\",\"shall\",\"be\"]", words: []string{"What", "must", "be", "acknowledgment", "shall", "be"},
			maxWidth: 16,
			Expected: []string{"What   must   be", "acknowledgment  ", "shall be        "}},
		{Name: "case with \"Science\",\"is\",\"what\",\"we\",\"understand\",\"well\",\"enough\",\"to\",\"explain\",\"to\",\"a\",\"computer.\",\"Art\",\"is\",\"everything\",\"else\",\"we\",\"do\"", words: []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			maxWidth: 20,
			Expected: []string{"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  "}},
	}
	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := fullJustify(cse.words, cse.maxWidth)
			assert.Equalf(t, cse.Expected, result, "for %d and %d expected %t, got %t", cse.words, cse.maxWidth, cse.Expected, result)
		})
	}
}
