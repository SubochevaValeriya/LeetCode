package main

func firstMissingPositive(nums []int) int {
	var numsMin int
	if nums[0] > 0 {
		numsMin = nums[0]
	}
	numsMax := nums[0]

	for _, num := range nums {
		if numsMin > num && num >= 0 {
			numsMin = num
		}

		if num > numsMin {
			if numsMax < num {
				numsMax = num
			}
		}
	}

	if numsMin > 1 || numsMax < 0 {
		return 1
	}

	numsMin2 := numsMax
	for _, num := range nums {
		if num > numsMin && num < numsMin2 {
			numsMin2 = num
		}
	}

	if numsMin2 - numsMin

	for i := numsMin + 1; i <= numsMax-1; i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}

	return numsMax + 1
}

func firstMissingPositive2(nums []int) int {
	var numsMin int
	if nums[0] > 0 {
		numsMin = nums[0]
	}
	numsMax := nums[0]
	m := map[int]int{}
	m[numsMin]++

	for _, num := range nums {
		if numsMin > num && num >= 0 {
			numsMin = num
			m[num]++
		}

		if num > numsMin {
			if numsMax < num {
				numsMax = num
			}
			m[num]++
		}
	}

	if numsMin > 1 || numsMax < 0 {
		return 1
	}

	for i := numsMin + 1; i <= numsMax-1; i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}

	return numsMax + 1
}
