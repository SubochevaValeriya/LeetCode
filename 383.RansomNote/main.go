package main

import "fmt"

func main() {
	ransomNote := "a"
	magazine := "ba"
	fmt.Println(canConstruct(ransomNote, magazine))
}

func canConstruct(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}
	dict := map[byte]int{}

	for i := 0; i < len(magazine); i++ {
		dict[magazine[i]] += 1
	}

	for i := 0; i < len(ransomNote); i++ {
		if _, ok := dict[ransomNote[i]]; ok && dict[ransomNote[i]] > 0 {
			dict[ransomNote[i]]--
			continue
		}
		return false
	}

	return true
}

func canConstruct(ransomNote string, magazine string) bool {
	dict := map[string]int{}

	for _, val := range magazine {
		dict[string(val)]++
	}

	for i := 0; i < len(ransomNote); i++ {
		val := string(ransomNote[i])
		if _, ok := dict[val]; ok && dict[val] > 0 {
			dict[val]--
			ransomNote = ransomNote[0:i] + ransomNote[i+1:]
			i--
			continue
		}
	}

	if len(ransomNote) > 0 {
		return false
	}

	return true
}

//func canConstruct(ransomNote string, magazine string) bool {
//	for i := 0; i < len(ransomNote); i++ {
//		for j := 0; j < len(magazine); j++ {
//			if ransomNote[i] == magazine[j] {
//				ransomNote = ransomNote[0:i] + ransomNote[i+1:]
//				magazine = magazine[0:j] + magazine[j+1:]
//				i--
//				j--
//				break
//			}
//		}
//	}
//	if len(ransomNote) > 0 {
//		return false
//	}
//
//	return true
//}
