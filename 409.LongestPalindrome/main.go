package main

import (
	"fmt"
)

func main() {
	//s := "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"
	s := "bb"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) int {

	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	m := map[byte]int{}
	lenghtOfThePalindrome := 0
	oddCount := 0

	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	for _, val := range m {
		if val%2 == 0 {
			lenghtOfThePalindrome += val
		} else {
			lenghtOfThePalindrome += val
			oddCount += 1
		}
	}

	if oddCount == 0 {
		return lenghtOfThePalindrome
	}

	return lenghtOfThePalindrome + 1 - oddCount
}
