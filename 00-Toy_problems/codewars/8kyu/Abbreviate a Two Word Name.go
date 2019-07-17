package kyu8

import "strings"

func AbbrevName(name string) string {
	var parts []string
	stringList := strings.Split(name, " ")
	for _, str := range stringList {
		firstLetter := str[:1]
		parts = append(parts, strings.ToUpper(firstLetter))
	}
	return strings.Join(parts, ".")
}

func AbbrevNameOld(name string) string {
	var result string
	stringArray := strings.Split(name, " ")
	for i, str := range stringArray {
		firstLetter := strings.ToUpper(string(str[0]))
		if i != len(stringArray)-1 {
			firstLetter = firstLetter + "."
		}
		result += string(firstLetter)
	}
	return result

}
