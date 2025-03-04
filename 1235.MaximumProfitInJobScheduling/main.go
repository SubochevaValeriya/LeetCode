package main

import (
	"fmt"
	"sort"
)

func main() {
	//startTime := []int{1, 2, 3, 4, 6}
	//endTime := []int{3, 5, 10, 6, 9}
	//profit := []int{20, 20, 100, 70, 60}

	startTime := []int{4, 3, 1, 2, 4, 8, 3, 3, 3, 9}
	endTime := []int{5, 6, 3, 5, 9, 9, 8, 5, 7, 10}
	profit := []int{9, 9, 5, 12, 10, 11, 10, 4, 14, 7}

	//startTime := []int{1, 1, 1}
	//endTime := []int{2, 3, 4}
	//profit := []int{5, 6, 4}

	fmt.Println(jobScheduling(startTime, endTime, profit))
}

type job struct {
	startTime int
	endTime   int
	profit    int
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	jobs := make([]job, len(profit))

	for i := range jobs {
		jobs[i].startTime = startTime[i]
		jobs[i].endTime = endTime[i]
		jobs[i].profit = profit[i]
	}

	//sorting
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].endTime < jobs[j].endTime
	})
	sort.Ints(endTime)

	dp := make([]int, len(jobs))
	dp[0] = jobs[0].profit
	for i := 1; i < len(jobs); i++ {
		lastNonConflictingJob := lastNonConflictingJob(endTime, jobs[i].startTime)
		var lastNonConflictingJobProfit int
		if lastNonConflictingJob != -1 {
			lastNonConflictingJobProfit = dp[lastNonConflictingJob]
		}
		dp[i] = max(dp[i-1], jobs[i].profit+lastNonConflictingJobProfit)
	}

	return dp[len(dp)-1]
}

func lastNonConflictingJob(endTime []int, starTimeI int) int {
	//lastNonConflJob, _ := slices.BinarySearch(endTime, starTimeI)
	lastNonConflJob := sort.Search(len(endTime), func(i int) bool {
		return endTime[i] > starTimeI
	})

	if endTime[lastNonConflJob] <= starTimeI {
		return lastNonConflJob
	}
	if lastNonConflJob == 0 {
		return -1
	}
	if endTime[lastNonConflJob-1] <= starTimeI {
		return lastNonConflJob - 1
	}

	return -1
}
