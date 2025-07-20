package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	houses := []int{282475249, 622650073, 984943658, 144108930, 470211272, 101027544, 457850878, 458777923}
	heaters := []int{823564440, 115438165, 784484492, 74243042, 114807987, 137522503, 441282327, 16531729, 823378840, 143542612}

	fmt.Println(findRadius(houses, heaters))
}

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)

	maxRadius := 0
	i := 0

	for _, house := range houses {
		for i < len(heaters)-1 && heaters[i+1] <= house {
			i++
		}

		minDist := int(math.Abs(float64(house - heaters[i])))
		if i < len(heaters)-1 {
			minDist = min(minDist, heaters[i+1]-house)
		}

		maxRadius = max(maxRadius, minDist)
	}

	return maxRadius
}
