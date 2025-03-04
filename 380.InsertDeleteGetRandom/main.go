package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := Constructor()

	//fmt.Println(r.Remove(0))
	//fmt.Println(r.Remove(0))
	fmt.Println(r.Insert(0))
	fmt.Println(r.Remove(0))
	fmt.Println(r)
	fmt.Println(r.Insert(0))
	fmt.Println(r)

	//fmt.Println(r.Insert(0))
	//fmt.Println(r.Insert(1))
	//fmt.Println(r.Remove(0))
	//fmt.Println(r.Insert(2))
	//fmt.Println(r.Remove(1))
	//fmt.Println(r)
	//fmt.Println(r.GetRandom())
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

type RandomizedSet struct {
	m    map[int]int
	nums []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{m: map[int]int{},
		nums: []int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}

	this.nums = append(this.nums, val)
	this.m[val] = len(this.nums) - 1

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	indx, ok := this.m[val]
	if !ok {
		return false
	}
	delete(this.m, val)
	lastIndx := len(this.nums) - 1
	if this.nums[lastIndx] != val {
		this.m[this.nums[lastIndx]] = indx
		this.nums[indx] = this.nums[lastIndx]
	}
	this.nums = this.nums[:lastIndx]

	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
