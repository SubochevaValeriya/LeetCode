package main

func main() {

}

func minSteps(s string, t string) int {
	frequencyMapT := map[rune]int{}

	for _, letter := range t {
		frequencyMapT[letter]++
	}

	var count int

	for _, letter := range s {
		if frequencyMapT[letter] > 0 {
			frequencyMapT[letter]--
		} else {
			count++
		}
	}

	return count
}
