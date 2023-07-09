package main

func main() {

}

func kidsWithCandies(candies []int, extraCandies int) []bool {

	maxCandies := candies[0]
	for i := 1; i < len(candies); i++ {
		if candies[i] > maxCandies {
			maxCandies = candies[i]
		}
	}

	kidsWithGreatestNumberOfCandies := make([]bool, len(candies))

	for i := range kidsWithGreatestNumberOfCandies {
		if candies[i]+extraCandies >= maxCandies {
			kidsWithGreatestNumberOfCandies[i] = true
		}
	}
	return kidsWithGreatestNumberOfCandies
}
