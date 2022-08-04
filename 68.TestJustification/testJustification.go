package main

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	answer := []string{}
	var line string
	var qWordsInLine int
	for i := 0; i < len(words); i++ {
		word := words[i]
		if len(word)+len(line) <= maxWidth {
			line = line + word + " "
			qWordsInLine++
			if i == len(words)-1 {
				if maxWidth-len(line) >= 0 {
					line = line[:len(line)] + strings.Repeat(" ", maxWidth-len(line))
				} else {
					line = strings.Trim(line, " ")
				}
				answer = append(answer, line)
			}
		} else {
			line = strings.Trim(line, " ")
			addSpaces := maxWidth - len(line)
			if qWordsInLine == 1 {
				line = line + strings.Repeat(" ", addSpaces)
			} else {
				line = strings.Replace(line, " ", strings.Repeat(" ", addSpaces/(qWordsInLine-1)+1), -1)
				if addSpaces%(qWordsInLine-1) != 0 {
					for j := 0; j < len(line)-1; j++ {
						if string(line[j]) == " " && string(line[j+1]) != " " && maxWidth != len(line) {
							line = line[:j] + " " + line[j:]
							j++
						}
					}
				}
			}
			answer = append(answer, line)
			line = ""
			qWordsInLine = 0
			i--
		}
	}
	return answer
}

//func fullJustify(words []string, maxWidth int) []string {
//	answer := []string{}
//	var line string
//	var qWordsInLine int
//	for i := 0; i < len(words); i++ {
//		word := words[i]
//		if len(word)+len(line) <= maxWidth {
//			line = line + word + " "
//			qWordsInLine++
//			if i == len(words)-1 {
//				if maxWidth-len(line) >= 0 {
//					line = line[:len(line)] + strings.Repeat(" ", maxWidth-len(line))
//				} else {
//					line = strings.Trim(line, " ")
//				}
//				answer = append(answer, line)
//			}
//		} else {
//			line = strings.Trim(line, " ")
//			addSpaces := maxWidth - len(line)
//			if qWordsInLine == 1 {
//				line = line + strings.Repeat(" ", addSpaces)
//			} else {
//				line = strings.Replace(line, " ", strings.Repeat(" ", addSpaces/(qWordsInLine-1)+1), -1)
//				if addSpaces%(qWordsInLine-1) != 0 {
//					line = addSpacesToLeftSide(line, maxWidth)
//				}
//			}
//			answer = append(answer, line)
//			line = ""
//			qWordsInLine = 0
//			i--
//		}
//	}
//	return answer
//}
//
//func addSpacesToLeftSide(line string, width int) string {
//	for i := 0; i < len(line)-1; i++ {
//		if string(line[i]) == " " && string(line[i+1]) != " " && width != len(line) {
//			line = line[:i] + " " + line[i:]
//			i++
//		}
//	}
//	return line
//}
