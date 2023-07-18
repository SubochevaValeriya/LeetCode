package main

func main() {

}

func maxOperations(nums []int, k int) int {
	operations := 0

	diffWithK := map[int]int{}

	for _, num := range nums {
		if diffWithK[num] > 0 {
			operations++
			diffWithK[num]--
			continue
		}
		diffWithK[k-num]++
	}

	return operations
}
