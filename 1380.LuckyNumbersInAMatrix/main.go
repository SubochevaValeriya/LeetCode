package main

func main() {

}

func luckyNumbers(matrix [][]int) []int {
	minRows := make([]int, len(matrix))
	maxColumn := make([]int, len(matrix[0]))
	result := []int{}

	for i, ints := range matrix {
		for j := range ints {
			number := ints[j]
			if number < minRows[i] || minRows[i] == 0 {
				minRows[i] = number
			}
			if number > maxColumn[j] {
				maxColumn[j] = number
			}
		}
	}

	for _, rowMin := range minRows {
		for _, columnMax := range maxColumn {
			if rowMin == columnMax {
				result = append(result, rowMin)
			}
		}
	}

	return result
}
