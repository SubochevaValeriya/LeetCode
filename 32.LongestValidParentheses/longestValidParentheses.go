package main

func longestValidParentheses2(s string) int {
	for len(s) < 2 {
		return 0
	}

	var answer, longestAns int

	left := 0
	right := 0

	for i := range s {
		if string(s[i]) == "(" {
			left++
		} else {
			right++
		}

		if left == right {
			answer = left + right
		}
		if right > left {
			left = 0
			right = 0
		}

		if answer > longestAns {
			longestAns = answer
		}
	}

	left, right, answer = 0, 0, 0

	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == "(" {
			left++
		} else {
			right++
		}

		if left == right {
			answer = left + right
		}

		if left > right {
			left = 0
			right = 0
		}

		if answer > longestAns {
			longestAns = answer
		}
	}

	return longestAns
}

//for right >= left {
//	fmt.Println(string(s[left]), left, right, s)
//	if left == 0 {
//		left++
//	}
//	if val, ok := m[string(s[left])]; ok {
//		if val == string(s[right]) && len(s[left+1:right])%2 == 0 {
//			answer += 2
//			if right == left-1 && answer > longestAns {
//				longestAns = answer
//				return longestAns
//			} else {
//				left++
//				right--
//			}
//		} else {
//			answer = 0
//			right--
//		}
//	} else {
//		left++
//	}
//}

func longestValidParentheses(s string) int {
	for len(s) < 2 {
		return 0
	}

	m := map[string]string{
		"(": ")",
	}

	stack := []string{}
	var answer, longestAns int

	for i, _ := range s {
		if _, ok := m[string(s[i])]; ok {
			stack = append(stack, string(s[i]))
		} else if len(stack) > 0 {
			if m[stack[len(stack)-1]] == string(s[i]) {
				stack = stack[:len(stack)-1]
				answer += 2
			}
			if answer > longestAns {
				longestAns = answer
			}
		} else {
			stack = []string{}
			answer = 0
		}
	}

	return longestAns
}
