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
