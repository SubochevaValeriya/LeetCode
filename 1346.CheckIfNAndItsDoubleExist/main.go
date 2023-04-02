package main

func main() {

}

func checkIfExist(arr []int) bool {

	m := map[int]struct{}{}

	for _, num := range arr {

		_, ok := m[num*2]
		if ok {
			return true
		}
		if num%2 == 0 {
			_, ok = m[num/2]
			if ok {
				return true
			}
		}

		m[num] = struct{}{}

	}

	return false

}
