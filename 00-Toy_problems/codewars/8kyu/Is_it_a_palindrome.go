package kyu8

import (
	"math"
	"strings"
)

func IsPalindrome(str string) bool {
	if len(str) < 2 {
		return true
	}
	str = strings.ToLower(str)
	mid := int(math.Floor(float64(len(str))))
	for currentIndex := 0; currentIndex < mid; currentIndex++ {
		lastIndex := len(str) - 1 - currentIndex
		currentElement := str[currentIndex]
		lastElement := str[lastIndex]
		if currentElement != lastElement {
			return false
		}
	}
	return true
}
