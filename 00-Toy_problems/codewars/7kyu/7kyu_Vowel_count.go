package kyu7

import "strings"

func GetCount(str string) (count int) {
	for _, char := range str {
		switch string(char) {
		case "a":
			count++
		case "e":
			count++
		case "i":
			count++
		case "o":
			count++
		case "u":
			count++
		}
	}
	return
}

func GetCount1(str string) (count int) {
	for _, char := range str {
		switch char {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return
}

func GetCount2(str string) (count int) {
	vowels := []string{"a", "e", "i", "o", "u"}
	for _, vowel := range vowels {
		count += strings.Count(str, vowel)
	}
	return
}
