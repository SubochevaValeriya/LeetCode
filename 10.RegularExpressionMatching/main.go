package main

import "fmt"

func main() {

}

func match(s string, p string) bool {
	lastJ := 0
	count := 0

	for i := 0; i < len(s); i++ {
		for j := lastJ; j < len(p); j++ {
			if s[i] == p[j] || p[j] == '.' {
				return match(s[1:], p[1:])
				count = 0
				lastJ++
				fmt.Println(i, j, "FF")
				if i != len(s)-1 && lastJ > len(p)-1 {
					return false
				}
				break
			}
			if p[j] == '*' && p[j-1] == s[i] {
				fmt.Println(i, j, lastJ, "DD")
				count++
				lastJ++
				if i != len(s)-1 && lastJ > len(p)-1 {
					return false
				}
				break
			}
			if p[j] == '*' && p[j-1] != s[i] && count == 0 {
				fmt.Println(i, j, lastJ, "FF")
				lastJ = j + 1
				if i != len(s)-1 && lastJ > len(p)-1 {
					return false
				}
				i--
				break
			}

			if p[j] == '.' && j-lastJ == 1 && p[lastJ] == '*' {
				count = 0
				lastJ++
				if i != len(s)-1 && lastJ > len(p)-1 {
					return false
				}
				break
			}
		}
	}
}

func isMatch(s string, p string) bool {

	dp := make([]bool, len(s))
	if s[0] == p[0] || p[0] == '.' {
		dp[0] = true
	} else {
		dp[1] = false
	}

	lastJ := 0
	count := 0
	for i := 0; i < len(s); i++ {
		for j := lastJ; j < len(p); j++ {
			fmt.Println(i, j)
			if s[i] == p[j] || p[j] == '.' {
				count = 0
				lastJ++
				fmt.Println(i, j, "FF")
				if i != len(s)-1 && lastJ > len(p)-1 {
					return false
				}
				break
			}
			if p[j] == '*' && p[j-1] == s[i] {
				fmt.Println(i, j, lastJ, "DD")
				count++
				lastJ++
				if i != len(s)-1 && lastJ > len(p)-1 {
					return false
				}
				break
			}
			if p[j] == '*' && p[j-1] != s[i] && count == 0 {
				fmt.Println(i, j, lastJ, "FF")
				lastJ = j + 1
				if i != len(s)-1 && lastJ > len(p)-1 {
					return false
				}
				i--
				break
			}

			if p[j] == '.' && j-lastJ == 1 && p[lastJ] == '*' {
				count = 0
				lastJ++
				if i != len(s)-1 && lastJ > len(p)-1 {
					return false
				}
				break
			}
		}
	}

	if lastJ < len(p) {
		return false
	}

	return true
	//for i < len(s) && j < len(p) {
	//	if j == len(p)-1 && i != len(s)-1 {
	//		fmt.Println(i, j)
	//		return false
	//	}
	//	if s[i] == p[j] {
	//		i++
	//		j++
	//	} else {
	//		switch string(p[j]) {
	//		case "*":
	//			forCheck = string(p[j-1])
	//			if string(s[i]) == forCheck || forCheck == "." {
	//				i++
	//				count++
	//			} else {
	//				if count > 0 {
	//					return false
	//				}
	//				j++
	//				i++
	//				count = 0
	//			}
	//		case ".":
	//			i++
	//			j++
	//		default:
	//			j++
	//		}
	//	}
	//}
	//
	////fmt.Println(dp)
	//fmt.Println(i, j)
	//if j >= len(p)-1 && i >= len(s)-1 {
	//	return true
	//	//	return dp[len(dp)-1]
	//}
	//return false
}

//func match(s string, p string) bool {
