package kyu6

import "math"

func Parse(data string) []int {
	var result []int
	var total int
	for _, char := range data {
		switch char {
		case 'i':
			total++
		case 'd':
			total--
		case 's':
			total = total * total
		case 'o':
			result = append(result, total)
		}
	}
	return result
}

func ParseWrong(data string) []int {
	var result []int
	var total float64
	for _, char := range data {
		switch char {
		case 'i':
			total++
		case 'd':
			total--
		case 's':
			total = math.Pow(total, 2)
		case 'o':
			// Converting big float64 to int results in int overflow. So don't.
			result = append(result, int(total))
		}
	}
	return result
}
