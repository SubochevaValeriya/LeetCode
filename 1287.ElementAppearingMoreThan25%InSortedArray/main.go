package main

func main() {

}

func findSpecialInteger(arr []int) int {
	countCurrent := 0
	currentNumber := arr[0]
	for _, num := range arr {
		if num == currentNumber {
			countCurrent++
			if float64(countCurrent)/float64(len(arr)) > 0.25 {
				return currentNumber
			}
			continue
		}

		countCurrent = 1
		currentNumber = num
	}

	return currentNumber
}
