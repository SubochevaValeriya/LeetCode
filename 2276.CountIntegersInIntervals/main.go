package main

import (
	"fmt"
	"sort"
)

func main() {
	//countInt := Constructor()
	//countInt.Add(2, 3)
	//countInt.Add(7, 10)
	//fmt.Println(countInt.Count(), "FF")
	//countInt.Add(5, 8)
	//fmt.Println(countInt.Count())
	//fmt.Println(countInt)

	//countInt := Constructor()
	//countInt.Add(10, 27)
	//countInt.Add(46, 50)
	//countInt.Add(15, 35)
	//countInt.Add(12, 32)
	//countInt.Add(7, 15)
	//countInt.Add(49, 49)
	//fmt.Println(countInt.Count(), "FF")
	//fmt.Println(countInt)

	//countInt := Constructor()
	//countInt.Add(5, 10)
	//countInt.Add(3, 5)
	//fmt.Println(countInt.Count(), "FF")
	//fmt.Println(countInt)

	//countInt := Constructor()
	//countInt.Add(881, 918)
	//countInt.Add(492, 775)
	//countInt.Add(611, 999)
	////countInt.Add(959, 959)
	////countInt.Add(313, 330)
	////countInt.Add(473, 928)
	//fmt.Println(countInt.Count(), "FF")
	//fmt.Println(countInt)

	//countInt := Constructor()
	//countInt.Add(571, 770)
	//countInt.Add(920, 996)
	//countInt.Add(65, 512)
	//countInt.Add(959, 959)
	//countInt.Add(313, 330)
	//countInt.Add(473, 928)
	//fmt.Println(countInt.Count(), "FF")
	//fmt.Println(countInt)

	//countInt := Constructor()
	//countInt.Add(457, 717)
	//countInt.Add(920, 996)
	//countInt.Add(65, 512)
	//countInt.Add(959, 959)
	//countInt.Add(313, 330)
	//countInt.Add(473, 928)
	//fmt.Println(countInt.Count(), "FF")
	//fmt.Println(countInt)

	countInt := Constructor()
	countInt.Add(457, 717)
	countInt.Add(918, 927)
	countInt.Add(660, 675)
	countInt.Add(885, 905)
	countInt.Add(323, 416)
	countInt.Add(774, 808)
	countInt.Add(671, 787)
	countInt.Add(133, 963)
	countInt.Add(374, 823)
	fmt.Println(countInt.Count(), "FF")
	fmt.Println(countInt)
}

type CountIntervals struct {
	intervals [][]int
	count     int
}

func Constructor() CountIntervals {

	return CountIntervals{
		intervals: [][]int{},
		count:     0,
	}
}
func (this *CountIntervals) Add(left int, right int) {
	fmt.Println(this.intervals, left, right)
	right++
	if len(this.intervals) == 0 {
		this.intervals = append(this.intervals, []int{left, right})
		this.count += right - left
		return
	}

	if this.intervals[len(this.intervals)-1][1] < left {
		this.intervals = append(this.intervals, []int{left, right})
		this.count += right - left
		return
	}

	if this.intervals[0][0] > right {
		this.intervals = append([][]int{{left, right}}, this.intervals...)
		this.count += right - left
		return
	}

	first := sort.Search(len(this.intervals), func(i int) bool {
		return left <= this.intervals[i][1]
	})

	second := sort.Search(len(this.intervals), func(i int) bool {
		return this.intervals[i][0] > right
	}) - 1

	if right < this.intervals[first][0] {
		this.intervals = append(this.intervals, []int{})
		copy(this.intervals[first+1:], this.intervals[first:])
		this.intervals[first] = []int{left, right}
		this.count += right - left
		return
	}

	if right > this.intervals[second][1] {
		this.count += right - this.intervals[second][1]
		this.intervals[second][1] = right
	}

	if left < this.intervals[first][0] {
		this.count += this.intervals[first][0] - left
		this.intervals[first][0] = left
	}

	for j := first + 1; j <= second; j++ {
		this.count += this.intervals[j][0] - this.intervals[j-1][1]
	}
	this.intervals[second][0] = this.intervals[first][0]
	copy(this.intervals[first:], this.intervals[second:])
	this.intervals = this.intervals[:len(this.intervals)-(second-first)]
}

func (this *CountIntervals) Count() int {
	return this.count
	//count := 0
	//for i := 0; i < len(this.intervals); i++ {
	//	count += this.intervals[i][1] - this.intervals[i][0] + 1
	//}
	//this.count = count
	//return count
}

/**
 * Your CountIntervals object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(left,right);
 * param_2 := obj.Count();
 */
