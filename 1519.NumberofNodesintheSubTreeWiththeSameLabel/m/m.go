package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := make([]int, len(a)+1)
	copy(b, a)
	fmt.Println(b)
}
