package main

import "strings"

func main() {

}

func numberOfBeams(bank []string) int {
	var beams int
	var lastLaserNumber int
	for _, row := range bank {
		currentLaserNumber := strings.Count(row, "1")
		if currentLaserNumber == 0 {
			continue
		}
		if lastLaserNumber != 0 {
			beams += lastLaserNumber * currentLaserNumber
		}
		lastLaserNumber = currentLaserNumber
	}

	return beams
}
