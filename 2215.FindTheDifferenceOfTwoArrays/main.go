package main

func main() {

}

func findDifference(nums1 []int, nums2 []int) [][]int {
	mNum1 := map[int]int{}
	mNum2 := map[int]int{}

	for _, num := range nums1 {
		mNum1[num]++
	}

	distinct1 := []int{}
	distinct2 := []int{}

	for _, num := range nums2 {
		mNum2[num]++
	}

	for _, num := range nums1 {
		_, ok := mNum2[num]
		if !ok {
			if mNum1[num] != -1 {
				distinct1 = append(distinct1, num)
				mNum1[num] = -1
			}
		}
	}

	for _, num := range nums2 {
		_, ok := mNum1[num]
		if !ok {
			if mNum2[num] != -1 {
				distinct2 = append(distinct2, num)
				mNum2[num] = -1
			}
		}
	}

	return [][]int{distinct1, distinct2}
}
