package main

func isIsomorphic(s string, t string) bool {

	m := map[byte]byte{}
	m2 := map[byte]byte{}

	for i := range s {
		if val, ok := m[s[i]]; ok {
			if val != t[i] {
				return false
			}
		} else if _, ok := m2[t[i]]; ok {
			return false
		} else {
			m[s[i]] = t[i]
			m2[t[i]] = s[i]
		}
	}

	return true
}
