package main

func checkString(s string) bool {
	checkForA := false

	for _, letter := range s {

		if checkForA == true {
			if string(letter) == "a" {
				return false
			}
		}
		if string(letter) == "b" {
			checkForA = true
		}
	}

	return true
}
