package kyu8

import "strings"

func DNAtoRNA(dna string) string {
	var str strings.Builder
	for _, char := range dna {
		switch char {
		case 'T':
			str.WriteString("U")
		default:
			str.WriteString(string(char))
		}
	}
	return str.String()
}

func DNAtoRNA1(dna string) string {
	return strings.Replace(dna, "T", "U", -1)
}

func DNAtoRNA2(dna string) string {
	var response string
	for _, letter := range dna {
		if letter == 'T' {
			letter = 'U'
		}
		response += string(letter)
	}
	return response
}
