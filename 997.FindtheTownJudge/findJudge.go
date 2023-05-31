package main

func main() {

}

func findJudge(n int, trust [][]int) int {

	trustSomebody := map[int]struct{}{}
	howManyPeopleTrust := map[int]int{}

	for i := 0; i < len(trust); i++ {
		trustSomebody[trust[i][0]] = struct{}{}
		howManyPeopleTrust[trust[i][1]]++
	}

	contenders := []int{}

	for i := 1; i <= n; i++ {
		if _, ok := trustSomebody[i]; !ok {
			contenders = append(contenders, i)
		}
	}

	if len(contenders) == 0 {
		return -1
	}

	for _, contender := range contenders {
		if howManyPeopleTrust[contender] != n-1 {
			return contender
		}
	}

	return -1
}
