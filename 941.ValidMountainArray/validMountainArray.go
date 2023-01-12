package main

func validMountainArray(arr []int) bool {
	peakPassed := false
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			return false //not strictly
		}

		switch peakPassed {
		case false:
			if arr[i] < arr[i-1] {
				if i == 1 {
					return false // only decreasing (peak not passed)
				}
				peakPassed = true
			}
		case true:
			if peakPassed == true {
				if arr[i] > arr[i-1] {
					return false // increasing again after peak passed
				}
			}
		}
	}

	if peakPassed == true {
		return true
	}

	return false // only increasing (peak not passed)
}
