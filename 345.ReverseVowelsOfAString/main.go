package main

func main() {

}

func reverseVowels(s string) string {

	sRune := []rune(s)

	vowels := map[rune]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
		'A': {},
		'E': {},
		'I': {},
		'O': {},
		'U': {},
	}

	left, right := 0, len(s)-1

	for left <= right {
		_, leftVowel := vowels[sRune[left]]
		_, rightVowel := vowels[sRune[right]]
		if leftVowel && rightVowel {
			sRune[left], sRune[right] = sRune[right], sRune[left]
			left++
			right--
		}

		if !leftVowel {
			left++
		}
		if !rightVowel {
			right--
		}
	}

	return string(sRune)
}
