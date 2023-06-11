package main

import (
	"fmt"
	"time"
)

func main() {
	letters := []byte{'c', 'f', 'j'}
	var target byte = 'j'
	fmt.Println(string(nextGreatestLetter(letters, target)))
}

func nextGreatestLetter(letters []byte, target byte) byte {

	left, right := 0, len(letters)-1
	for right >= left {
		fmt.Println(left, right)
		median := left + (right-left)/2
		if letters[median] <= target {
			fmt.Println(left, right, median)
			left++
		}
		if letters[median] > target {
			if median == 0 {
				return letters[median]
			}
			if letters[median-1] <= target {
				return letters[median]
			}
			right--
		}

		if median == 0 && median == len(letters)-1 {
			return letters[0]
		}

		time.Sleep(2 * time.Second)
	}

	return letters[0]
}
