package main

import "fmt"

func lengthOfLISы(nums []int) int {
	var length, maxLenght, lastnum, diff int
	//m := map[int]int{}
	//for i, num := range nums {
	//	m[i] = num
	//}
	//
	//sort.Ints(nums)
	//newJ := 0
	//for _, num := range nums {
	//	if newJ == len(nums) {
	//		break
	//	}
	//	for j := newJ; j < len(nums); j++ {
	//		if j != 0 && m[j] == m[j-1] {
	//			continue
	//		}
	//		if m[j] == num {
	//			length++
	//			newJ = j + 1
	//			break
	//		}
	//	}
	//}

	for i, num := range nums {
		length = 0
		lastnum = num
		for _, num2 := range nums[i+1:] {
			//fmt.Println(num, num2, lastnum, length, nums, diff)
			if num2 > num && num2 > lastnum {
				fmt.Println(num, num2, lastnum, length, nums, "f")
				length++
				diff = num2 - lastnum
				lastnum = num2
			} else if lastnum-num2 <= diff && num2 > num {
				fmt.Println(num, num2, lastnum, length, nums, "s")
				lastnum = num2
				diff = lastnum - num2
			}
			fmt.Println(num, num2, lastnum, length, nums, "e", diff)
		}

		if length > maxLenght {
			maxLenght = length
		}
	}

	return maxLenght + 1
}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	maxLen := 0

	for i, num := range nums {
		dp[i] = 1
		if i > 0 {
			for j := 0; j < i; j++ {
				if num > nums[j] {
					dp[i] = max(dp[i], dp[j]+1)
				}
			}
		}

		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
