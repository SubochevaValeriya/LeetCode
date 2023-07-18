package main

import "fmt"

func main() {
	l := Constructor(2)
	l.Put(1, 1)
	l.Put(2, 2)
	l.Get(1)
	fmt.Println(l)
	l.Put(3, 3)
	fmt.Println(l.Get(2))
	fmt.Println(l)
}

type LRUCache struct {
	capacity   int
	keyValue   map[int]int
	keyTime    map[int]int
	operations []int
	leastUsed  int
}

func Constructor(capacity int) LRUCache {

	return LRUCache{
		capacity:   capacity,
		keyValue:   make(map[int]int, 3),
		keyTime:    make(map[int]int, 3),
		operations: []int{},
		leastUsed:  0,
	}
}

func (this *LRUCache) Get(key int) int {

	value, ok := this.keyValue[key]
	if !ok {
		return -1
	}
	this.operations = append(this.operations, key)
	this.keyTime[key] = len(this.operations) - 1
	return value
}

func (this *LRUCache) Put(key int, value int) {
	_, ok := this.keyValue[key]
	if len(this.keyValue) == this.capacity && !ok {
		for i := this.leastUsed; i < len(this.operations); i++ {
			if this.keyTime[this.operations[i]] == i {
				this.leastUsed = i + 1
				delete(this.keyValue, this.operations[i])
				delete(this.keyTime, this.operations[i])
				break
			}
		}
	}

	this.keyValue[key] = value
	this.operations = append(this.operations, key)
	this.keyTime[key] = len(this.operations) - 1
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
