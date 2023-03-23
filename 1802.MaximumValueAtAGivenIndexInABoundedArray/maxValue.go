package main

import (
	"errors"
	"fmt"
)

func main() {
	//fmt.Println(maxValue(3, 2, 18)) //7
	//fmt.Println(maxValue(4, 2, 6))
	//fmt.Println(maxValue(8, 6, 20)) //4
	//fmt.Println(maxValue(6, 1, 10)) //3

	//0 1 2 3 4 5
	////
	//2 3 2 1 1 1

	//	fmt.Println(maxValue(8, 7, 14))
}

func maxValue(n int, index int, maxSum int) int {
	if n == 1 {
		return maxSum
	}

	maxInd := 0
	right := maxSum
	left := 0
	//lastPretend := 0
	pretendents := map[int]int{}
	for pretendent := left + (right-left)/2; right >= left; pretendent = left + (right-left)/2 {
		fmt.Println("JJJJJ", pretendent, left, right)
		if pretendent+(n-1) > maxSum {
			right = pretendent
			continue
		}

		mS, ok := pretendents[pretendent]
		if ok {
			if mS < maxSum {
				mS2, ok := pretendents[pretendent+1]
				if ok {
					if mS2 > maxSum {
						return pretendent
					}
				} else {
					pretendent = pretendent + 1
				}
			}
		}
		//if lastPretend == pretendent {
		//	fmt.Println("ds", pretendent)
		//	return pretendent
		//}

		sum := pretendent
		leftSum, err := calcSum(pretendent-1, index, sum, maxSum)
		if err != nil {
			right = pretendent
			if maxInd != 0 {
				return maxInd
			}
			//lastPretend = pretendent
			continue
		}

		fmt.Println(pretendent, sum, leftSum, "jj")
		sum, err = calcSum(pretendent-1, n-index-1, leftSum, maxSum)
		fmt.Println(sum, pretendent, "ddsasda")
		//if err != nil {
		//	fmt.Println("s", pretendent)
		//	right = pretendent
		//	if maxInd != 0 {
		//		return maxInd
		//	}
		//	lastPretend = pretendent
		//	continue
		//}
		fmt.Println(pretendent, sum, "KKK")
		fmt.Println(pretendent, sum)
		fmt.Println(left, right)
		if sum == maxSum {
			return pretendent
		}

		if sum > maxSum {
			pretendents[pretendent] = sum
			right = pretendent
			if maxInd != 0 {
				return maxInd
			}
			//lastPretend = pretendent
			continue
		}

		if sum < maxSum {
			pretendents[pretendent] = sum
			fmt.Println(pretendent, "MIN")
			left = pretendent
			//lastPretend = pretendent
			continue
			//fmt.Println(pretendent, sum)
			//fmt.Println((maxSum - sum) / n)
			//return pretendent + (maxSum-sum)/n
		}
		//nums := make([]int, n)
		//
		//if lastPretend == pretendent {
		//	return pretendent
		//}
		//
		//nums[index] = pretendent
		//sum := pretendent
		//
		//iLeft := index - 1
		//iRight := index + 1
		//for iLeft > -1 || iRight < n {
		//	if iRight < n {
		//		if nums[iRight-1] == 1 {
		//			sum += n - iRight
		//			iRight = n
		//		} else {
		//			nums[iRight] = nums[iRight-1] - 1
		//			sum += nums[iRight]
		//			iRight++
		//		}
		//	}
		//
		//	if iLeft > -1 {
		//		if nums[iLeft+1] == 1 {
		//			sum += iLeft + 1
		//			iLeft = -1
		//		} else {
		//			nums[iLeft] = nums[iLeft+1] - 1
		//			sum += nums[iLeft]
		//			iLeft--
		//		}
		//	}
		//
		//	if sum > maxSum {
		//		break
		//	}
		//}
		//
		////for i := index + 1; i < n; i++ {
		////	if nums[i-1] == 1 {
		////		nums[i] = nums[i-1]
		////	} else {
		////		nums[i] = nums[i-1] - 1
		////	}
		////	sum += nums[i]
		////	if sum > maxSum {
		////		break
		////	}
		////}
		////
		////for i := index - 1; i >= 0; i-- {
		////	if nums[i+1] == 1 {
		////		nums[i] = nums[i+1]
		////	} else {
		////		nums[i] = nums[i+1] - 1
		////	}
		////	sum += nums[i]
		////	if sum > maxSum {
		////		break
		////	}
		////}
		//
		//fmt.Println(pretendent, sum, nums)

		//if sum > maxSum {
		//	right = pretendent
		//	if maxInd != 0 {
		//		return maxInd
		//	}
		//	lastPretend = pretendent
		//	continue
		//}
		//
		//if sum < maxSum {
		//	left = pretendent
		//	lastPretend = pretendent
		//	continue
		//	//fmt.Println(pretendent, sum)
		//	//fmt.Println((maxSum - sum) / n)
		//	//return pretendent + (maxSum-sum)/n
		//}
		//
		//if sum == maxSum {
		//	return pretendent
		//}
	}

	return -1
}

func calcSum(pretendent int, index int, sum int, maxSum int) (int, error) {
	if sum > maxSum {
		return sum, errors.New("err")
	}
	if pretendent == 1 {
		sum += index
		return sum, nil
	}

	if index == 0 {
		return sum, nil
	}
	sum += pretendent
	fmt.Println(pretendent, sum, "fddf")
	return calcSum(pretendent-1, index-1, sum, maxSum)
}
