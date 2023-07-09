package main

func main() {

}

func canPlaceFlowers(flowerbed []int, n int) bool {
	if flowerbed[0] == 0 {
		if len(flowerbed) == 1 {
			if n <= 1 {
				return true
			}
			return false
		}
		if flowerbed[1] == 0 {
			flowerbed[0] = 1
			n--
		}
	}

	for i := 1; i < len(flowerbed); i++ {
		if n <= 0 {
			return true
		}
		if i == len(flowerbed)-1 {
			if flowerbed[i-1] == 0 && flowerbed[i] == 0 {
				flowerbed[i] = 1
				n--
				continue
			}
		}
		if flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			n--
		}
	}
	if n <= 0 {
		return true
	}
	return false
}
