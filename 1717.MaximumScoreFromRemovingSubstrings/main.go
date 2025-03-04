package main

import (
	"fmt"
	"slices"
)

func main() {
	//s := "cdbcbbaaabab"
	//x := 4
	//y := 5
	s := "aabbaaxybbaabb"
	x := 5
	y := 4
	fmt.Println(maximumGain(s, x, y))
}

func maximumGain(s string, x int, y int) int {
	var points, morePoints, lessPoints int
	bytes := []byte(s)
	if x > y {
		bytes, morePoints = replace(bytes, []byte("ab"), x)
		fmt.Println(s, morePoints)
		bytes, lessPoints = replace(bytes, []byte("ba"), y)
	} else {
		bytes, morePoints = replace(bytes, []byte("ba"), y)
		bytes, lessPoints = replace(bytes, []byte("ab"), x)
	}

	points = points + morePoints + lessPoints

	return points
}

func replace(bytes, substr []byte, points int) ([]byte, int) {
	var pointSum, writeIndex int

	for i := 0; i < len(bytes); i++ {
		bytes[writeIndex] = bytes[i]
		if writeIndex > 0 && slices.Equal(bytes[writeIndex-1:writeIndex+1], substr) {
			writeIndex -= 2
			pointSum += points
		}
		writeIndex++
	}

	return bytes[:writeIndex], pointSum
}

//func maximumGain(s string, x int, y int) int {
//	start, end := -1, -1
//	points := 0
//
//	for i, letter := range s {
//		if (letter == 'b' || letter == 'a') && start == -1 {
//			start = i
//		}
//		if (letter != 'b' && letter != 'a') && start != -1 {
//			end = i
//		} else if i == len(s)-1 && start != -1 {
//			end = i + 1
//		}
//
//		if end != -1 {
//			points += maximumGainSubstr(s[start:end], x, y)
//			start = -1
//			end = -1
//		}
//	}
//
//	return points
//}
//
//func maximumGainSubstr(s string, x int, y int) int {
//	if len(s) < 2 {
//		return 0
//	}
//	morePoint := x
//	lessPoint := y
//	morePointsStr := "ab"
//	lessPointsStr := "ba"
//	if y > x {
//		morePoint = y
//		lessPoint = x
//		morePointsStr = "ba"
//		lessPointsStr = "ab"
//	}
//	//109980
//	replaced := true
//	points := 0
//
//	for replaced {
//		newS, morePoints := replace(s, morePointsStr, morePoint, -1)
//		fmt.Println(morePoint, newS)
//		newS, lessPoints := replace(newS, lessPointsStr, lessPoint, -1)
//		fmt.Println(lessPoints, newS)
//
//		s = newS
//		//fmt.Println(morePoints, lessPoints)
//		points = points + morePoints + lessPoints
//		replaced = morePoints != 0 || lessPoints != 0
//	}
//
//	return points
//}
//
//func replace(s, substr string, point, times int) (string, int) {
//	var pointSum, timesReplaced int
//	for i := 1; i < len(s); i++ {
//		if timesReplaced != -1 && times == timesReplaced {
//			return s, pointSum
//		}
//		if i == 0 {
//			i++
//			continue
//		}
//		if s[i-1:i+1] == substr {
//			timesReplaced++
//			pointSum += point
//			sNew := ""
//			if i != 1 {
//				sNew = s[:i-1]
//			}
//			sNew = sNew + s[i+1:]
//			s = sNew
//			i -= 2
//		}
//	}
//
//	return s, pointSum
//}
