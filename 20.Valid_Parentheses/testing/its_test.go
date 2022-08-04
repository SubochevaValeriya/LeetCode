package main

import (
	"testing"
)

func BenchmarkSample(b *testing.B) {
	s := "[]()[()()][{}]()"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isValid(s)
	}
}
func isValid(s string) bool {

	//if len(s) < 2 - always false

	if len(s) < 2 {
		return false
	}

	//Creating of variables

	r := []rune(s) // slice of runes

	return isValidTech(r, variablesForIsValid())

	// map for Parentheses:

	// base cases:

	//if len(r) == 2 {
	//	if string(r[1]) == m[string(r[0])] {
	//		return true
	//	} else {
	//		return false
	//	}
	//}

	// three true cases with a recursion

	//if string(r[1]) == m[string(r[0])] {
	//	r = r[2:]
	//	s = string(r)
	//	return isValid(s)
	//}
	//
	//if string(r[len(r)-1]) == m[string(r[0])] {
	//	r = r[1 : len(r)-1]
	//	s = string(r)
	//	return isValid(s)
	//}
	//
	//if string(r[len(r)-1]) == m[string(r[len(r)-2])] {
	//	r = r[:len(r)-2]
	//	s = string(r)
	//	return isValid(s)
	//}

	// go through slice

	//for i := 0; i < len(r)-2; i++ {
	//	j := i + 1
	//	if string(r[j]) == m[string(r[i])] {
	//		newr := []rune{}
	//		newr = append(newr, r[:i]...)
	//		newr = append(newr, r[j+1:]...)
	//		s = string(newr)
	//		return isValid(s)
	//	}
	//}
	//
	//return false
}

func isValidTech(r []rune, m map[string]string) bool {

	if len(r) == 2 {
		if string(r[1]) == m[string(r[0])] {
			return true
		} else {
			return false
		}
	}

	for i := 0; i < len(r)-2; i++ {
		j := i + 1
		if string(r[j]) == m[string(r[i])] {
			newr := []rune{}
			newr = append(newr, r[:i]...)
			newr = append(newr, r[j+1:]...)
			return isValidTech(newr, variablesForIsValid())
		}
	}

	return false
}

func variablesForIsValid() map[string]string {

	m := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}
	return m
}

// проверяем количество отк и закр одинаково
// ([)]
// если нашли отк ищем закр
// если находим отк другого типа - должен быть закр такого же типа
