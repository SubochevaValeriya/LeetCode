package main

func main() {

}

func minimumRecolors(blocks string, k int) int {
	var recolors, minRecolors int
	for i := 0; i < len(blocks); i++ {
		if blocks[i] == 'W' {
			recolors++
		}
		if i < k-1 {
			continue
		}
		if i == k-1 || recolors < minRecolors {
			minRecolors = recolors
		}
		if blocks[i-(k-1)] == 'W' {
			recolors--
		}
	}

	return minRecolors
}
