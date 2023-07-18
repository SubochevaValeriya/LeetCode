package main

import (
	"fmt"
	"strconv"
)

func main() {
	//chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	chars := []byte{'a', 'a', 'a', 'a', 'a', 'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c'}
	fmt.Println(compress(chars))
	//fmt.Println(chars)
}

func compress(chars []byte) int {
	start := 0
	//initialLenght := len(chars)
	fmt.Println(len(chars))
	if len(chars) == 0 {
		return 0
	}

	currentCharacter := chars[0]
	count := 1

	for i := 1; i < len(chars); i++ {
		if chars[i] == currentCharacter {
			count++
			fmt.Printf("%c %v %v \n", chars[i], i, count)
		} else {
			chars[start] = currentCharacter
			start++
			currentCharacter = chars[i]
			if count != 1 {
				sliceLeft := append(chars[:start], []byte(strconv.Itoa(count))...)
				chars = append(sliceLeft, chars[len(sliceLeft):]...)
				start += len([]byte(strconv.Itoa(count)))
			}
			count = 1
		}
	}
	chars[start] = currentCharacter
	start++
	if count != 1 {
		chars = append(chars[:start], []byte(strconv.Itoa(count))...)
		start += len([]byte(strconv.Itoa(count)))
	}
	chars = chars[:start]
	fmt.Printf("%c", chars)
	return len(chars[:start])
}
