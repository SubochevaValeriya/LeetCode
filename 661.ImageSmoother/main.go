package main

import (
	"fmt"
	"math"
)

func main() {
	img := [][]int{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}}

	fmt.Println(imageSmoother(img))

}

func imageSmoother(img [][]int) [][]int {
	result := make([][]int, len(img))
	sum := 0
	count := 1

	for i := 0; i < len(img); i++ {
		for j := 0; j < len(img[i]); j++ {
			if result[i] == nil {
				result[i] = make([]int, len(img[i]))
			}

			if j == 0 {
				sum = img[i][j]
				count = 1
				if len(img[i]) != 1 {
					sum += img[i][j+1]
					count++
				}
				if len(img) != 1 {
					if i != len(img)-1 {
						sum += img[i+1][j]
						count++
						if len(img[i]) != 1 {
							sum += img[i+1][j+1]
							count++
						}
					}
					if i != 0 {
						sum += img[i-1][j]
						count++
						if len(img[i]) != 1 {
							sum += img[i-1][j+1]
							count++
						}
					}
				}
			}

			if j != 0 {
				if j != len(img[i])-1 {
					sum += img[i][j+1]
					count++
					if i != 0 {
						sum += img[i-1][j+1]
						count++
					}
					if i != len(img)-1 {
						sum += img[i+1][j+1]
						count++
					}
				}
				if j != 1 {
					sum -= img[i][j-2]
					count--
					if i != 0 {
						sum -= img[i-1][j-2]
						count--
					}
					if i != len(img)-1 {
						sum -= img[i+1][j-2]
						count--
					}
				}
			}
			fmt.Println(sum, count)
			result[i][j] = int(math.Floor(float64(sum) / float64(count)))
		}
	}

	return result
}
