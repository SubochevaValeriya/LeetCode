package main

import "sort"

func main() {

}

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	m := map[uint8][]uint8{}

	for i := range s1 {
		for

		if s1[i] != s2[i] {
			sort.Slice()
			m[s1[i]] = append(m[s1[i]], s2[i])
			m[s2[i]] = append(m[s2[i]], s1[i])
		}
	}


	parents := make([]int, 26)

	ASCIISymbol := 97
	for i := range parents {
		parents[i] = ASCIISymbol
		ASCIISymbol += 1
	}



	m := map[string]string{}


}
