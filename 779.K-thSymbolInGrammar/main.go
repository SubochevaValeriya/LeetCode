package main

import (
	"fmt"
	"math"
)

func main() {

	n, k := 2, 2
	fmt.Println(kthGrammar(n, k))
}

func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}

	totalElements := math.Pow(2, float64(n-1))
	halfElements := int(totalElements / 2)
	if k > halfElements {
		return 1 - kthGrammar(n, k-halfElements)
	}

	return kthGrammar(n-1, k)
}

//
//func recursion(n, k int) int {
//	if n == 1 {
//		return 0
//	}
//
//	totalElements := math.Pow(2, float64(n-1))
//	halfElements := int(totalElements / 2)
//	if k > halfElements {
//		return 1 - recursion(n, k-halfElements)
//	}
//
//	return recursion(n-1, k)
//}
//
