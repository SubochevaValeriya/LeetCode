package main

func main() {

}

type candidate struct {
	x int
	y int
}

func numSpecial(mat [][]int) int {
	onesInColumns := map[int]int{}
	onesInRows := map[int]int{}

	for i, ints := range mat {
		for j, num := range ints {
			if num == 1 {
				onesInRows[i]++
				onesInColumns[j]++
			}
		}
	}

	countSpecials := 0
	for i, ints := range mat {
		for j, num := range ints {
			if num == 1 && onesInRows[i] == 1 && onesInColumns[j] == 1 {
				countSpecials++
			}
		}
	}

	return countSpecials
}
