package main

import (
	"fmt"
)

func main() {
	var s []byte
	var element byte
	//	c, err := fmt.Scanf("%p", s)
	for {
		fmt.Scan(&element)
		if element == 0 {
			break
		} else {
			s = append(s, element)
		}
	}
	fmt.Println(s)
	//for i := range in {
	//	_, err := fmt.Scan(&in[i])
	//	if err != nil {
	//		return in[:i], err
	//		for i :=
	//			fmt.Scan(&s)
	//			fmt.Println(s)
	fmt.Println(reverseString(s))
}

func reverseString(s []byte) []byte {
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}

	//for i := len(s) - 1; i >= 0; i-- {
	//	result = append(result, s[i])
	//}
	//s = result
	return s
}
