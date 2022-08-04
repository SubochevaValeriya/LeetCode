package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var max int = 10e9
	var min int = -10e9
	fmt.Println(rand.Intn(max-min) + min)
	initial := []int{3, 2, 4}
	target := 6
	fmt.Println(twoSum(initial, target))
}

//
//func BenchmarkSimplest(b *testing.B) {
//	initial := []int{3, 2, 4}
//	target := 6
//	fmt.Println(twoSum(initial, target))
//}

func twoSum(nums []int, target int) []int {
	var result int
	m := make(map[int]int)
	var answer []int
	for i, integer1 := range nums {
		_, found := m[integer1]
		if found == true {
			answer = append(answer, i, m[integer1])
			return answer
		} else if found == false {
			result = target - integer1
			m[result] = i
		}
	}
	return nil
}

/*func twoSum(nums []int, target int) []int {
	var result int
	var answer []int
	for i := 0; i < len(nums); i++ {
		if result == target && len(answer) == 2 {
			break
		} else if result <= target && len(answer) < 2 {
			result += nums[i]
			answer = append(answer, i)
			if len(answer) == 2 && answer[0] != answer[1] {
				result -= nums[answer[0]]
				answer[0] = answer[1]
				answer = answer[:1]
				if len(answer) < 2 && i == len(nums)-1 {
					i = -1
				}
			}
		} else if result > target && len(answer) == 1 {
			result += nums[i]
			answer = append(answer, i)
			if len(answer) == 2 && result != target && answer[0] != answer[1] {
				result -= nums[answer[0]]
				answer[0] = answer[1]
				answer = answer[:1]
			}
			if len(answer) < 2 && i == len(nums)-1 {
				i = -1
			}
		} else if result > target && len(answer) > 1 {
			result -= nums[answer[0]]
			answer[0] = answer[1]
			answer = answer[:1]
			if len(answer) < 2 && result != target && i == len(nums)-1 {
				i = -1
			}
		}
	}
	return answer
}
*/
//for index, integer := range nums {
//if target == 0 {
//	for i, integer := range nums {
//		if integer == 0 && len(answer) < 2 {
//			result += integer
//			answer = append(answer, i)
//		} else {
//			continue
//		}
//	}
//} else {

//	package main
//
//	import "fmt"
//
//func main() {
//	initial := []int{3, 2, 3}
//	target := 6
//	fmt.Println(twoSum(initial, target))
//}

//func twoSum(nums []int, target int) []int {
//	var result int
//	var answer []int
//	//for index, integer := range nums {
//	if result == target && len(answer) == 2 {
//		break
//	} else if result < target && len(answer) < 2 {
//		result += integer
//		answer = append(answer, index)
//		if len(answer) == 2 && result != target && answer[0] != answer[1] {
//			result -= nums[answer[0]]
//			answer[0] = answer[1]
//			answer = answer[:1]
//		} else if len(answer) < 2 && result != target {
//			index = 0
//		}
//	} else {
//		result -= nums[answer[0]]
//		answer[0] = answer[1]
//		answer = answer[:1]
//	}
//}
//
//if len(answer) < 2 {
//
//}
//return answer
//}
//
