package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/excel-sheet-column-number/description/

func main() {
	fmt.Println(titleToNumber("BBB"))
}

//2
//26*3 + 2
//26*2 + 2
//2
//
//26 * 2
//2

func titleToNumber(columnTitle string) int {

	number := 0
	power := 0
	for i := len(columnTitle) - 1; i >= 0; i-- {
		currentNumber := columnTitle[i] - 64
		number += int(math.Pow(26, float64(power))) * int(currentNumber)
		power++
	}
	return number
	// Z - 26
	// AA - 27
	// AAA -
	// ZZ - 701

	// ZZZ - 26*26*26

}
