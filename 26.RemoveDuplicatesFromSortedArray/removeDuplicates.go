package main

func RemoveDuplicates2(nums []int) int {
	answer := []int{}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] != nums[j] {
				answer = append(answer, nums[i])
				i = j - 1
				break
			}
		}
	}
	answer = append(answer, nums[len(nums)-1])
	nums = append(nums[:0], answer...)
	return len(nums)
}

func removeDuplicates3(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)
}

func removeDuplicates4(nums []int) int {
	var answer int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] != 1 {
			for j := i + 1; j < len(nums); j++ {
				if nums[i] < nums[j] {
					nums[i+1] = nums[j]
					break
				}
			}
		}
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			answer = i - 1
			break
		} else if i == len(nums)-1 {
			answer = i
		}
	}

	return answer + 1
}
