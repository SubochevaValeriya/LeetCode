package main

import "fmt"

func main() {

	//answerKey := "TFFT"
	//k := 1

	//answerKey := "TTFF"
	//k := 2
	//fmt.Println(maxConsecutiveAnswers(answerKey, k))

	//answerKey := "FFFTTFTTFT"
	//k := 3

	answerKey := "FTFFTFTFTTFTTFTTFFTTFFTTTTTFTTTFTFFTTFFFFFTTTTFTTTTTTTTTFTTFFTTFTFFTTTFFFFFTTTFFTTTTFTFTFFTTFTTTTTTF"
	fmt.Println(len(answerKey))
	k := 32
	fmt.Println(maxConsecutiveAnswers(answerKey, k))
}

func maxConsecutiveAnswers(answerKey string, k int) int {
	// T case
	maxConsecutiveQuestions := maxConsecutive(answerKey, 0, 'F', k)
	// F case
	maxConsecutiveQuestions = maxConsecutive(answerKey, maxConsecutiveQuestions, 'T', k)
	return maxConsecutiveQuestions
}

func maxConsecutive(answerKey string, maxConsecutiveQuestions int, wrongAnswer uint8, leftReplacements int) int {
	left := 0
	for right := 0; right < len(answerKey); right++ {
		if answerKey[right] == wrongAnswer {
			if leftReplacements == 0 {
				if right-left > maxConsecutiveQuestions {
					maxConsecutiveQuestions = right - left
				}
				for left <= right {
					if answerKey[left] == wrongAnswer {
						left++
						leftReplacements++
						break
					}
					left++
				}
			}
			leftReplacements--
		}
		if right == len(answerKey)-1 && leftReplacements >= 0 {
			if right-left+1 > maxConsecutiveQuestions {
				maxConsecutiveQuestions = right - left + 1
			}
		}
	}
	return maxConsecutiveQuestions
}
