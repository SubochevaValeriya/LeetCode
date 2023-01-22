package main

import "fmt"

func main() {
	fmt.Println(missingRolls([]int{4, 2, 2, 5, 4, 5, 4, 5, 3, 3, 6, 1, 2, 4, 2, 1, 6, 5, 4, 2, 3, 4, 2, 3, 3, 5, 4, 1, 4, 4, 5, 3, 6, 1, 5, 2, 3, 3, 6, 1, 6, 4, 1, 3}, 2, 53))
}

func missingRolls(rolls []int, mean int, n int) []int {

	sum := (n + len(rolls)) * mean
	var knownSum int
	missing := []int{}

	for _, roll := range rolls {
		knownSum += roll
	}

	missingSum := sum - knownSum

	if missingSum/6 > n {
		return missing
	}

	for {
		missingMean := missingSum / n
		if missingMean > 6 || missingMean < 1 {
			return []int{}
		}
		if missingSum-missingMean < 0 {
			break
		}
		missingSum -= missingMean
		missing = append(missing, missingMean)
		n--
		if n == 0 {
			return missing
		}
	}

	return missing

}
