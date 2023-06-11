package main

import "sort"

func main() {

}

func canMakeArithmeticProgression(arr []int) bool {

	sort.Ints(arr)
	baseDiff := arr[1] - arr[0]

	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != baseDiff {
			return false
		}
	}
	return true
}
