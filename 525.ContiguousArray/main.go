package main

func main() {

}

func findMaxLength(nums []int) int {
	maxLength := 0
	prefixSumOnIndex := map[int]int{}
	prefixSum := 0
	for i, num := range nums {
		switch num {
		case 0:
			prefixSum -= 1
		case 1:
			prefixSum += 1
		}
		if prefixSum == 0 {
			if i+1 > maxLength {
				maxLength = i + 1
			}
			continue
		}

		indx, ok := prefixSumOnIndex[prefixSum]
		if !ok {
			prefixSumOnIndex[prefixSum] = i
		} else {
			if i-indx > maxLength {
				maxLength = i - indx
			}
		}
	}

	return maxLength
}
