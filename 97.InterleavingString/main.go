package main

import "fmt"

func main() {

	//s1 := "aabcc"
	//s2 := "dbbca"
	//s3 := "aadbbcbcac"

	s1 := "abababababababababababababababababababababababababababababababababababababababababababababababababbb"
	s2 := "babababababababababababababababababababababababababababababababababababababababababababababababaaaba"
	s3 := "abababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababbb"

	fmt.Println(isInterleave(s1, s2, s3))
}

//
//type step struct {
//	is      bool
//	s1Index int
//	s2Index int
//}

func isInterleave(s1 string, s2 string, s3 string) bool {
	fmt.Println(s1, s2, s3)
	if len(s1) == 0 && len(s2) == 0 && len(s3) == 0 {
		return true
	}
	if len(s3) == 0 || len(s3) < len(s2)+len(s1) {
		return false
	}

	if len(s1) != 0 {
		if s3[0] == s1[0] {
			if isInterleave(s1[1:], s2, s3[1:]) {
				return true
			}
		}
	}

	if len(s2) != 0 {
		if s3[0] == s2[0] {
			if isInterleave(s1, s2[1:], s3[1:]) == true {
				return true
			}
		}
	}

	return false
}

//dp1 := make([]int, len(s3))
//dp2 := make([]int, len(s3))
//dp := make([]bool, len(s3))
//
//if s3[0] != s1[0] {
//dp1[0] = -1
//}
//
//if s3[0] != s2[0] {
//dp2[0] = -1
//}
//
//if dp1[0] == 0 || dp2[0] == 0 {
//dp[0] = true
//} else {
//return false
//}
//
//for i := 1; i < len(s3); i++ {
//dp1[i] = dp1[i-1]
//dp2[i] = dp2[i-1]
//
//if s3[i] == s1[dp1[i]+1] {
//dp1[i]++
//}
//
//if s3[i] == s2[dp2[i]+1] {
//dp2[i]++
//}
//
//if dp1[0] == 0 || dp2[0] == 0 {
//dp[0] = true
//} else {
//return false
//}
//}
