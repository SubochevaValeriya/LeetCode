package main

import "fmt"

func main() {

	gain := []int{44, 32, -9, 52, 23, -50, 50, 33, -84, 47, -14, 84, 36, -62, 37, 81, -36, -85, -39, 67, -63, 64, -47, 95, 91, -40, 65, 67, 92, -28, 97, 100, 81}

	fmt.Println(largestAltitude(gain))

}

func largestAltitude(gain []int) int {
	altitudes := make([]int, len(gain)+1)
	highest := 0

	for i := 1; i < len(altitudes); i++ {
		altitudes[i] = altitudes[i-1] + gain[i-1]
		if altitudes[i] > highest {
			highest = altitudes[i]
		}
	}
	fmt.Println(altitudes)
	return highest
}
