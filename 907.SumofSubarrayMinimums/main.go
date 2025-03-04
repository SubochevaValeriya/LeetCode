package main

import (
	"fmt"
)

func main() {
	//arr := []int{3, 1, 2, 4}
	arr := []int{11, 81, 94, 43, 3}
	//arr := []int{26, 23, 65, 36, 19, 51, 4, 11, 99, 20}
	//arr := []int{19, 19, 62, 66}
	//arr := []int{36, 74, 84, 92, 17, 68, 97, 6, 68, 85}
	//arr := []int{5, 26, 42, 49, 67, 71, 53, 45, 41, 16}
	//arr := []int{85, 93, 93, 90}
	//arr := []int{92, 80, 9, 62, 49}

	fmt.Println(sumSubarrayMins(arr))
}

type num struct {
	value int
	index int
}

func sumSubarrayMins(arr []int) int {
	minimum := []num{{value: arr[0],
		index: 0,
	}}
	dp := make([]int, len(arr))
	dp[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] <= minimum[0].value {
			dp[i] = arr[i] * (i + 1)
			minimum = []num{{value: arr[i],
				index: i,
			}}
			continue
		}

		if arr[i] > minimum[len(minimum)-1].value {
			minimum = append(minimum, num{
				value: arr[i],
				index: i,
			})
			dp[i] = dp[i-1] + arr[i]
			continue
		}

		for j := len(minimum) - 1; j > 0; j-- {
			fmt.Println(arr[i])
			if arr[i] <= minimum[j].value && arr[i] > minimum[j-1].value {
				fmt.Println(dp[minimum[j-1].index])
				dp[i] = dp[minimum[j-1].index] + arr[i]*(i-minimum[j-1].index)
				minimum = append(minimum[:j], num{
					value: arr[i],
					index: i,
				})
				break
			}
		}

	}
	var sum int
	for _, v := range dp {
		sum += v
	}

	fmt.Println(dp, sum)
	return sum
}
