// https://leetcode.com/problems/snapshot-array/submissions/968779780/
package main

import "fmt"

func main() {
	this := Constructor(3)
	this.Set(0, 5)
	this.Snap()
	this.Set(0, 6)
	fmt.Println(this.Get(0, 0))

	//this := Constructor(4)
	//this.Snap()
	//this.Snap()
	//fmt.Println(this.Get(3, 1))
	//this.Set(2, 4)
	//this.Snap()
	//this.Set(1, 4)

	this = Constructor(1)
	this.Set(0, 4)
	this.Set(0, 16)
	this.Set(0, 13)
	this.Snap()
	fmt.Println(this.Get(0, 0))
	this.Snap()

	this = Constructor(2)
	this.Snap()
	fmt.Println(this.Get(1, 0))
	fmt.Println(this.Get(0, 0))
	this.Set(1, 8)
	fmt.Println(this.Get(1, 0))
	this.Set(0, 20)
	fmt.Println(this.Get(0, 0))
	this.Set(0, 7)
	fmt.Println(this)
}

type SnapshotArray struct {
	array     []int
	snapshots [][][2]int
	snaps     int
}

func Constructor(length int) SnapshotArray {
	arr := SnapshotArray{
		array:     make([]int, length),
		snapshots: make([][][2]int, length),
		snaps:     -1,
	}

	for i := range arr.snapshots {
		arr.snapshots[i] = make([][2]int, 1)
		arr.snapshots[i][0] = [2]int{-1, 0}
	}
	return arr
}

func (this *SnapshotArray) Set(index int, val int) {
	this.array[index] = val
	if this.snapshots[index][len(this.snapshots[index])-1][0] == this.snaps+1 {
		this.snapshots[index][len(this.snapshots[index])-1][1] = val
		return
	}
	this.snapshots[index] = append(this.snapshots[index], [2]int{this.snaps + 1, val})
}

func (this *SnapshotArray) Snap() int {
	this.snaps++
	return this.snaps
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	left, right := 0, len(this.snapshots[index])-1
	for left <= right {
		median := left + (right-left)/2
		if this.snapshots[index][median][0] > snap_id {
			right--
			continue
		}

		if this.snapshots[index][median][0] == snap_id {
			return this.snapshots[index][median][1]
		}

		if this.snapshots[index][median][0] < snap_id {
			if median == len(this.snapshots[index])-1 {
				return this.snapshots[index][median][1]
			}
		}

		if this.snapshots[index][median+1][0] > snap_id {
			return this.snapshots[index][median][1]
		}

		left++
	}
	return this.array[index]
}
