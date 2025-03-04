package main

import "fmt"

func main() {
	nums := []int{1, 3, 4, 1, 2, 3, 1}
	fmt.Println(findMatrix(nums))

}

func findMatrix(nums []int) [][]int {
	result := [][]int{}
	frequencyMaps := []map[int]struct{}{}
	frequencyMaps = append(frequencyMaps, map[int]struct{}{})
	result = append(result, []int{})

	for _, num := range nums {
		for j, frequencyMap := range frequencyMaps {
			_, ok := frequencyMap[num]
			if ok && j != len(frequencyMaps)-1 {
				continue
			}
			if ok {
				frequencyMaps = append(frequencyMaps, map[int]struct{}{})
				result = append(result, []int{})
				j++
			}
			frequencyMaps[j][num] = struct{}{}
			result[j] = append(result[j], num)
			break
		}
	}

	return result
}
