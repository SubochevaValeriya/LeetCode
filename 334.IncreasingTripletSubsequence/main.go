package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(increasingTriplet([]int{1, 1, 1}))
}

func increasingTriplet(nums []int) bool {
	a, b := math.MaxInt64, math.MaxInt64
	for _, num := range nums {
		if num <= a {
			a = num
		} else if num <= b {
			b = num
		} else {
			fmt.Println(a, b)
			return true
		}
	}
	return false
}

func increasingTriplet3(nums []int) bool {
	a, b, c := 0, 1, 2

	if len(nums) < 3 {
		return false
	}

	for i := 1; i < len(nums); i++ {
		fmt.Println(nums[a], nums[b], nums[c], "ANSWER")
		if nums[a] < nums[b] && nums[b] < nums[c] {
			return true
		}

		if i > a && i > c && nums[i] > nums[c] {
			b = c
			c = i
		}

		if nums[a] >= nums[b] {
			if nums[i] < nums[a] {
				if i+1 < len(nums)-1 {
					if nums[i+1] > nums[i] {
						a = i
						i++
					}
				}
			}
		}

		if a == b {
			b++
		}

		if c == b {
			c++
		}

		if c > len(nums)-1 {
			return false
		}

		if nums[b] >= nums[c] {
			if c+1 > len(nums)-1 {
				if nums[i+1] <= nums[c] {
					b = i + 1
				}
			} else if nums[b] < nums[c+1] {

			} else if nums[i+1] <= nums[c] {
				b = i + 1
			}
		}

		if c == b {
			c++
		}

		if c > len(nums)-1 {
			return false
		}

		if nums[c] <= nums[b] {
			c++
			if c > len(nums)-1 {
				return false
			}
		}
		if nums[a] < nums[b] && nums[b] < nums[c] {
			if a < b && b < c {
				return true
			}
		}
	}

	return false
}

func increasingTriplet2(nums []int) bool {
	a, b, c := 0, 1, 2

	for i := 0; i <= len(nums)-2; i++ {
		fmt.Println(nums[a], nums[b], nums[c], "ANSWER")
		if c > len(nums)-1 {
			return false
		}
		if nums[a] < nums[b] && nums[b] < nums[c] {
			fmt.Println("DSD")
			return true
		}

		if nums[a] >= nums[b] {
			if nums[i+1] < nums[a] {
				a = i + 1
				i++
			}
		}

		if a == b {
			b++
			if b > len(nums)-1 {
				return false
			}
		}
		fmt.Println(nums[a], nums[b], nums[c], "S")
		if b == c {
			c++
			if c > len(nums)-1 {
				return false
			}
		}
		fmt.Println(nums[a], nums[b], nums[c], "S")

		if nums[a] < nums[b] && nums[b] < nums[c] {
			if a < b && b < c {
				return true
			}
		}
		if nums[b] >= nums[c] {
			fmt.Println(nums[i+2])
			fmt.Println(c)
			//fmt.Println(a, b, c)
			//fmt.Println(i + 2)
			//fmt.Println(nums[i+2], nums[c])
			if c+1 > len(nums)-1 {
				if nums[i+2] <= nums[c] {
					fmt.Println("d")
					b = i + 2
				}
			} else if nums[b] < nums[c+1] {

			} else if nums[i+2] <= nums[c] {
				fmt.Println("d")
				b = i + 2
			}
			fmt.Println("sd", nums[a], nums[b], nums[c])
		}

		if b == c {
			c++
		}

		if c > len(nums)-1 {
			return false
		}
		if nums[c] < nums[b] {
			c++
			if c > len(nums)-1 {
				return false
			}
		}
		fmt.Println(a, b, c, len(nums), "SEE")
	}
	//fmt.Println(nums[a], nums[b], nums[c])
	return false
}

//func increasingTriplet(nums []int) bool {
//
//	for i := 0; i < len(nums)-2; i++ {
//	loop:
//		for j := i + 1; j < len(nums)-1; j++ {
//			for k := j + 1; k < len(nums); k++ {
//				if nums[k] > nums[j] && nums[j] > nums[i] {
//					return true
//				}
//				if nums[j] <= nums[i] {
//					continue loop
//				}
//				if nums[k] <= nums[j] {
//					continue
//				}
//			}
//		}
//	}
//	return false
//}
