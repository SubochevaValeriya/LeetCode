package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println([]rune("A"))
	fmt.Println([]rune("B"))
	fmt.Println(convertToTitle(703))
}

//11881376
//2147483647
//308915776
//8031810176

const offset = 64

// (26)^3 * x 26^2 * y +26^1 * z + d

func convertToTitle(columnNumber int) string {
	var title []rune
	power := Log(26, columnNumber)
	if columnNumber < maxColumnNumberForPower(power) {
		power--
	}
	currentNumber := 26
	for remainder := columnNumber; remainder > 0; {
		if currentNumber <= 0 {
			currentNumber = 26
			power--
		}
		if remainder-int(math.Pow(26, float64(power)))*currentNumber >= 0 {
			if remainder-int(math.Pow(26, float64(power)))*currentNumber == 0 && power != 0 {
				currentNumber--
			}
			title = append(title, rune(currentNumber+offset))
			remainder = remainder - int(math.Pow(26, float64(power)))*currentNumber
			power--
			currentNumber = 26
		} else {
			currentNumber--
		}
	}

	return string(title)
}

func Log(base, x int) int {
	return int(math.Log(float64(x)) / math.Log(float64(base)))
}

func maxColumnNumberForPower(power int) int {
	var number int
	for i := 0; i <= power; i++ {
		number += int(math.Pow(26, float64(i)))
	}
	return number
}
