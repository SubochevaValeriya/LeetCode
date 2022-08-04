package main

import "fmt"

func main() {
	nums := []int{
		1, 2, 5, 6, 7, 2, 4,
	}
	k := 4
	t := 0

	fmt.Println(ContainsNearbyAlmostDuplicate(nums, k, t))
}

//func containsNearbyAlmostDuplicatee(nums []int, k int, t int) bool {
//	a := 0
//	lastValue := nums[0]
//
//	for i, value := range nums {
//		if i <= a+k {
//			if value <= lastValue+t && value <= lastVal	ue-t {
//				return true
//			}
//
//			lastValue = value
//		}
//		a = i
//	}
//
//	//dict := map[int]int{}
//	//	dict := map[int][]int{}
//	dict := map[[2]int]int{}
//	//type d struct {
//	//	start int
//	//	end int
//	//	index int
//	//}
//	//
//	//dict := []d{}
//	for i, value := range nums {
//		ar := [2]int{value - t, value + t}
//		////val1 = value - t
//		////val2 = value + t
//		//if _, found := dict[ar]; found {
//		//	if i-dict[ar] <= k {
//		//		return true
//		//	}
//		//}
//
//		dict[ar] = i
//	}
//	return false
//}

func numInArr(ar [2]int, num int) {

}

//func abs(n int) int {
//	if n < 0 {
//		return -n
//	}
//
//	return n
//}

func ContainsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if len(nums) == 1 {
		return false
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= i+k && j < len(nums); j++ {
			if abs(nums[i]-nums[j]) <= t {
				return true
			}
		}
	}

	return false
	//dict := map[int]int{}
	//for i, value := range nums {
	//	for j := range dict {
	//		if t >= abs(value-j) && dict[j] != i {
	//			if i-dict[j] <= k {
	//				return true
	//			}
	//		}
	//	}
	//	dict[value] = i
	//}
	//
	//return false
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

//func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
//	dict := map[int]int{}
//	for i, value := range nums {
//		for j := range dict {
//			if t >= abs(value-j) && dict[j] != i {
//				if i-dict[j] <= k {
//					return true
//				}
//			}
//		}
//		dict[value] = i
//	}
//
//	return false
//}
//
//func abs(n int) int {
//	if n < 0 {
//		return -n
//	}
//
//	return n
//}
