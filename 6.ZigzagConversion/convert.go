package main

import "strings"

func main() {

}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	lines := make([]strings.Builder, numRows)

	lineN := 0
	upDownSwitcher := false

	for _, letter := range s {
		lines[lineN].WriteRune(letter)
		if lineN == numRows-1 {
			upDownSwitcher = true
		}
		if lineN == 0 {
			upDownSwitcher = false
		}

		if upDownSwitcher == false {
			lineN++
		} else {
			lineN--
		}
	}

	answer := strings.Builder{}

	for _, line := range lines {
		answer.WriteString(line.String())
	}

	return answer.String()
}
