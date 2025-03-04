package main

func main() {

}
func findErrorNums(nums []int) []int {
	m := map[int]int{}
	answer := make([]int, 2)

	for _, num := range nums {
		m[num]++
	}

	for i := 1; i <= len(nums); i++ {
		if m[i] > 1 {
			answer[0] = i
		}
		if m[i] == 0 {
			answer[1] = i
		}
	}

	return answer
}
