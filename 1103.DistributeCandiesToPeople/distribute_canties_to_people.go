package main

//https://leetcode.com/problems/distribute-candies-to-people/

func distributeCandies(candies int, numPeople int) []int {
	answer := make([]int, numPeople, numPeople)
	n := 1

	for i := 0; i < numPeople; i++ {
		if candies-n < 0 {
			answer[i] += candies
			break
		}
		answer[i] += n
		candies -= n
		n++

		if i == numPeople-1 {
			i = -1
		}

		//switch i {
		//case 0:
		//	switcher = false
		//case numPeople:
		//	switcher = true
		//}
		//
		//if !switcher {
		//	i++
		//} else {
		//	i--
		//}
	}

	return answer
}
