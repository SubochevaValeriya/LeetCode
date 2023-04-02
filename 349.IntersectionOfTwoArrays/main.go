package main

func main() {

}

func intersection(nums1 []int, nums2 []int) []int {
	m := map[int]bool{}
	result := []int{}

	for _, num := range nums1 {
		m[num] = false
	}

	for _, num := range nums2 {
		added, ok := m[num]
		if ok && !added {
			result = append(result, num)
			m[num] = true
		}
	}
	return result
}
