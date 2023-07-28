package main

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)

	defer cancel()

	for {
		select {
		case <-ctx.Done():
			fmt.Println(errors.Is(ctx.Err(), context.DeadlineExceeded))
			return
		default:
			fmt.Println("ok")
		}
	}
	//nums := []int{2, 2, 2, 2, 2}
	////nums := []int{2, 2, 2, 2, 2}
	////nums := []int{1, 3, 5, 4, 7}
	//fmt.Println(findNumberOfLIS(nums))

}

func findNumberOfLIS(nums []int) int {
	length := make([]int, len(nums))
	count := make([]int, len(nums))
	maxLen := 1
	for i := 0; i < len(nums); i++ {
		length[i] = 1
		count[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if length[j]+1 > length[i] {
					length[i] = length[j] + 1
					count[i] = 0
				}
				if length[j]+1 == length[i] {
					count[i] += count[j]
				}
			}
		}
		if length[i] > maxLen {
			maxLen = length[i]
		}
	}

	var result int
	for i := 0; i < len(nums); i++ {
		if length[i] == maxLen {
			result += count[i]
		}
	}

	return result
}

func findNumberOfLISOld(nums []int) int {
	length := make([]int, len(nums))
	count := make([]int, len(nums))

	length[0] = 1
	count[0] = 1
	maxLen := 1

	mapLenghtToCount := map[int]int{}

	for i := 1; i < len(nums); i++ {
		length[i] = 1
		count[i] = count[i-1]
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				length[i] = int(math.Max(float64(length[i]), float64(length[j]+1)))
				if length[i] == maxLen {
					mapLenghtToCount[maxLen]++
					count[i] = mapLenghtToCount[length[i]] * mapLenghtToCount[length[j]]
					fmt.Println(mapLenghtToCount, i, j)
				} else if length[i] > maxLen {
					maxLen = length[i]
					if length[j] != 1 {
						count[i] = count[i-1]
						break
					} else {
						count[i] = 1
					}
				}
			}
		}
		if maxLen == 1 {
			count[i]++
			mapLenghtToCount[maxLen]++
		}
	}

	fmt.Println(length, count)
	return count[len(count)-1]
}
