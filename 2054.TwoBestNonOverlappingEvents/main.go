package main

import (
	"fmt"
	"sort"
)

func main() {
	events := [][]int{{1, 3, 2}, {4, 5, 2}, {2, 4, 3}}
	//events := [][]int{{1, 5, 3}, {1, 5, 1}, {6, 6, 5}}

	result := maxTwoEvents(events)
	fmt.Println(result)
}

func maxTwoEvents(events [][]int) int {
	n := len(events)
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	maxSuffixSum := make([]int, n)
	maxSuffixSum[n-1] = events[n-1][2]
	for i := n - 2; i >= 0; i-- {
		maxSuffixSum[i] = max(events[i][2], maxSuffixSum[i+1])
	}

	var maxSum int

	for _, event := range events {
		firstNonOverlappingEvent := sort.Search(n, func(i int) bool {
			return events[i][0] > event[1]
		})

		maxCurrentSum := event[2]

		if firstNonOverlappingEvent != n {
			maxCurrentSum += maxSuffixSum[firstNonOverlappingEvent]
		}

		if maxCurrentSum > maxSum {
			maxSum = maxCurrentSum
		}
	}

	return maxSum
}

//
//func maxTwoEvents(events [][]int) int {
//	sort.Slice(events, func(i, j int) bool {
//		if events[i][1] == events[j][1] {
//			if events[i][0] == events[j][0] {
//				return events[i][2] < events[j][2]
//			}
//
//			return events[i][0] < events[j][0]
//		}
//
//		return events[i][1] < events[j][1]
//	})
//	maxValue := events[0][2]
//	//for i, event := range events {
//	//	value := event[2]
//	//	for j := i - 1; j >= 0; j-- {
//	//		if events[j][1] < event[0] && event[2]+events[j][2] > value {
//	//			value = event[2] + events[j][2]
//	//		}
//	//	}
//	//
//	//	if value > maxValue {
//	//		maxValue = value
//	//	}
//	//}
//
//	maxLastValue := 0
//
//	for i, event := range events {
//		if i == 0 {
//			maxLastValue = event[2]
//			continue
//		}
//
//		if event[2] > maxLastValue{
//			maxLastValue = event[2]
//		}
//
//		if
//		//if events[i-1][1] < event[0] {
//		//	maxValueBeforeEndTime[i] = max(maxValueBeforeEndTime[i-1], events[i-1][2])
//		//	continue
//		//}
//
//		for j := i - 1; j >= 0; j-- {
//			if events[j][1] < event[0] && (j == 0 || events[j][1] != event[j-1]) {
//				maxValueAfterEndTime[i] = max(maxValueAfterEndTime[j], event[2])
//			}
//		}
//	}
//
//
//	maxValueAfterEndTime := make([]int, len(events))
//	for i, event := range events {
//		maxValueAfterEndTime[i] = event[2]
//		if i == 0 {
//			continue
//		}
//		//if events[i-1][1] < event[0] {
//		//	maxValueBeforeEndTime[i] = max(maxValueBeforeEndTime[i-1], events[i-1][2])
//		//	continue
//		//}
//
//		for j := i - 1; j >= 0; j-- {
//			if events[j][1] < event[0] && (j == 0 || events[j][1] != event[j-1]) {
//				maxValueAfterEndTime[i] = max(maxValueAfterEndTime[j], event[2])
//			}
//		}
//	}
//	fmt.Println(events, maxValueAfterEndTime)
//	for i, event := range events {
//		if i == 0 {
//			continue
//		}
//		value := event[2] + maxValueAfterEndTime[i-1]
//		if value > maxValue {
//			maxValue = value
//		}
//	}
//
//	return maxValue
//}
