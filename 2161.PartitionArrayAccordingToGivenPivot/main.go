package main

func main() {

}

func pivotArray(nums []int, pivot int) []int {
	less := make([]int, 0, len(nums))
	great := make([]int, 0, len(nums))
	middle := make([]int, 0, len(nums))

	for _, num := range nums {
		if num < pivot {
			less = append(less, num)
		} else if num > pivot {
			great = append(great, num)
		} else {
			middle = append(middle, num)
		}
	}

	return append(append(less, middle...), great...)
}
