package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(listToSlice(addTwoNumbers(makeList([]int{5}), makeList([]int{5}))))

	//	fmt.Println(listToSlice(addTwoNumbers(makeList([]int{7, 2, 4, 3}), makeList([]int{5, 6, 4}))))

	//fmt.Println(listToSlice(addTwoNumbers(makeList([]int{2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 9}), makeList([]int{5, 6, 4, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 9, 9, 9, 9}))))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {

	first, second := "", ""
	for {
		if l1 == nil && l2 == nil {
			break
		}
		if l1 != nil {
			first += strconv.Itoa(l1.Val)
			l1 = l1.Next
		}
		if l2 != nil {
			second += strconv.Itoa(l2.Val)
			//fmt.Println(second)
			l2 = l2.Next
		}
	}
	fmt.Println(first)
	firstNum, _ := strconv.Atoi(first)
	fmt.Println(firstNum)
	secondNum, _ := strconv.Atoi(second)
	fmt.Println(secondNum)
	sum := strconv.Itoa(firstNum + secondNum)
	fmt.Println(sum)
	answer := &ListNode{
		Val:  0,
		Next: nil,
	}

	ans := answer

	for i, numStr := range sum {
		num, _ := strconv.Atoi(string(numStr))
		answer.Next = &ListNode{
			Val:  0,
			Next: nil,
		}
		answer.Val = num
		if i == len(sum)-1 {
			answer.Next = nil
		}
		answer = answer.Next
	}
	return ans
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	links := map[*ListNode]*ListNode{}

	for {
		if l1.Next == nil && l2.Next == nil {
			break
		}
		if l1.Next != nil {
			links[l1.Next] = l1
			l1 = l1.Next
		}
		if l2.Next != nil {
			links[l2.Next] = l2
			l2 = l2.Next
		}
	}

	var remainder int
	sum := &ListNode{
		Val:  0,
		Next: nil,
	}

	if l1.Val+l2.Val > 9 {
		sum.Val = (l1.Val + l2.Val + remainder) % 10
		remainder = (l1.Val + l2.Val + remainder) / 10
	} else {
		remainder = 0
		sum.Val = l1.Val + l2.Val + remainder
	}
	if links[l1] == nil && links[l2] == nil {
		if remainder == 0 {
			return sum
		}
	}
	return dfs(links[l1], links[l2], sum, links, remainder)
}

func dfs(l1, l2, sum *ListNode, links map[*ListNode]*ListNode, remainder int) *ListNode {
	newSum := ListNode{
		Val:  0,
		Next: sum,
	}

	var first, second int
	if l1 != nil {
		first = l1.Val
	}
	if l2 != nil {
		second = l2.Val
	}
	if first+second+remainder > 9 {
		newSum.Val = (first + second + remainder) % 10
		remainder = (first + second + remainder) / 10
	} else {
		newSum.Val = first + second + remainder
		remainder = 0
	}

	_, ok := links[l1]
	_, ok2 := links[l2]
	if !ok && !ok2 {
		if remainder != 0 {
			return &ListNode{
				Val:  remainder,
				Next: &newSum,
			}
		}
		return &newSum
	}
	if !ok {
		return dfs(&ListNode{
			Val:  0,
			Next: nil,
		}, links[l2], &newSum, links, remainder)
	}
	if !ok2 {
		return dfs(links[l1], &ListNode{
			Val:  0,
			Next: nil,
		}, &newSum, links, remainder)
	}

	return dfs(links[l1], links[l2], &newSum, links, remainder)
}

func listToSlice(list *ListNode) []int {

	slice := []int{}

	for i := list; i != nil; i = i.Next {
		slice = append(slice, i.Val)
	}

	return slice
}

func makeList(slice []int) *ListNode {

	list := &ListNode{
		Val:  0,
		Next: nil,
	}

	var last = list
	for _, num := range slice {
		last.Next = &ListNode{
			Val:  num,
			Next: nil,
		}

		last = last.Next
	}

	return list.Next
}
