package main

//
//import (
//	"fmt"
//)
//
//func main() {
//	//countInt := Constructor()
//	//countInt.Add(2, 3)
//	//countInt.Add(7, 10)
//	//fmt.Println(countInt.Count(), "FF")
//	//countInt.Add(5, 8)
//	//fmt.Println(countInt.Count())
//	//fmt.Println(countInt)
//
//	//countInt := Constructor()
//	//countInt.Add(10, 27)
//	//countInt.Add(46, 50)
//	//countInt.Add(15, 35)
//	//countInt.Add(12, 32)
//	//countInt.Add(7, 15)
//	//countInt.Add(49, 49)
//	//fmt.Println(countInt.Count(), "FF")
//	//fmt.Println(countInt)
//
//	//countInt := Constructor()
//	//countInt.Add(5, 10)
//	//countInt.Add(3, 5)
//	//fmt.Println(countInt.Count(), "FF")
//	//fmt.Println(countInt)
//
//	//countInt := Constructor()
//	//countInt.Add(571, 770)
//	//countInt.Add(920, 996)
//	//countInt.Add(65, 512)
//	//countInt.Add(959, 959)
//	//countInt.Add(313, 330)
//	//countInt.Add(473, 928)
//	//fmt.Println(countInt.Count(), "FF")
//	//fmt.Println(countInt)
//
//	countInt := Constructor()
//	countInt.Add(457, 717)
//	countInt.Add(918, 927)
//	countInt.Add(660, 675)
//	countInt.Add(885, 905)
//	countInt.Add(323, 416)
//	countInt.Add(774, 808)
//	fmt.Println(countInt.Count(), "FF")
//	fmt.Println(countInt)
//}
//
//type CountIntervals struct {
//	intervals [][]int
//	count     int
//}
//
//func Constructor() CountIntervals {
//
//	return CountIntervals{
//		intervals: [][]int{},
//		count:     0,
//	}
//}
//
//func (this *CountIntervals) Add(left int, right int) {
//	//sort.Slice(this.intervals, func(i, j int) bool {
//	//	return this.intervals[i][0] < this.intervals[j][1]
//	//})
//	if len(this.intervals) == 0 {
//		this.intervals = append(this.intervals, []int{left, right})
//		this.count += right - left + 1
//		return
//	}
//
//	if this.intervals[len(this.intervals)-1][1] < left {
//		this.intervals = append(this.intervals, []int{left, right})
//		this.count += right - left + 1
//		return
//	}
//	for i := 0; i < len(this.intervals); i++ {
//		if i != 0 {
//			if this.intervals[i][0] > right && this.intervals[i-1][1] < left {
//				newInterval := append([][]int{{left, right}}, this.intervals[i:]...)
//				this.intervals = append(this.intervals[0:i], newInterval...)
//				this.count += right - left + 1
//				continue
//			}
//			if this.intervals[i][1] <= this.intervals[i-1][1] {
//				this.count = this.count - (this.intervals[i][1] - this.intervals[i][0] + 1)
//				this.intervals = append(this.intervals[0:i], this.intervals[i+1:]...)
//				i--
//				continue
//			}
//			if this.intervals[i][0] <= this.intervals[i-1][1] {
//				this.count = this.count - (this.intervals[i-1][1] - this.intervals[i][0] + 1)
//				this.intervals[i-1][1] = this.intervals[i][1]
//				this.intervals = append(this.intervals[0:i], this.intervals[i+1:]...)
//				i--
//				continue
//			}
//		}
//		if i == 0 && this.intervals[i][0] > right {
//			this.intervals = append([][]int{{left, right}}, this.intervals...)
//			this.count += right - left + 1
//			return
//		}
//		if left > this.intervals[i][1] {
//			continue
//		}
//
//		if left >= this.intervals[i][0] && right <= this.intervals[i][1] {
//			return
//		}
//		this.count = this.count - (this.intervals[i][1] - this.intervals[i][0] + 1)
//		if right >= this.intervals[i][1] {
//			this.intervals[i][1] = right
//		}
//		if left <= this.intervals[i][0] && right >= this.intervals[i][0] {
//			this.intervals[i][0] = left
//		}
//		this.count = this.count + (this.intervals[i][1] - this.intervals[i][0] + 1)
//	}
//}
//
//func (this *CountIntervals) Count() int {
//	return this.count
//	//count := 0
//	//for i := 0; i < len(this.intervals); i++ {
//	//	count += this.intervals[i][1] - this.intervals[i][0] + 1
//	//}
//	//this.count = count
//	//return count
//}
//
///**
// * Your CountIntervals object will be instantiated and called as such:
// * obj := Constructor();
// * obj.Add(left,right);
// * param_2 := obj.Count();
// */
