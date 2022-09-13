package main

// https://leetcode.com/problems/lemonade-change/

func lemonadeChange(bills []int) bool {
	m := map[int]int{}

	for _, bill := range bills {
		switch bill {
		case 10:
			if m[5] == 0 {
				return false
			}
			m[5]--
		case 20:
			if m[5] == 0 {
				return false
			}
			if m[10] > 0 {
				m[5]--
				m[10]--
			} else if m[5] < 3 {
				return false
			} else {
				m[5] -= 3
			}
		}

		m[bill]++
	}
	return true
}
