package main

import "fmt"

func main() {
	fmt.Println(mySqrt(8))
}

func mySqrt(x int) int {
	i := x / 2

	for {
		switch {
		case i*i == x || (i*i < x && (i+1)*(i+1) > x):
			return i
		case i*i > x:
			i = i / 2
		case i*i < x:
			i = i++
		}

		//if i*i == x {
		//	return i
		//}
		//if i*i < x {
		//	i = i / 2
		//} else if i*i < x {
		//	i = i * 2
		//}

	}
	return i - 1
}
