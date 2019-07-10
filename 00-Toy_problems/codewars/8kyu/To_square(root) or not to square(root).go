package kyu8

import "math"

func SquareOrSquareRoot(arr []int) []int {
	var result []int
	for _, num := range arr {
		rooted := math.Sqrt(float64(num))
		if float64(int(rooted)) == rooted {
			result = append(result, int(rooted))
		} else {
			result = append(result, int(math.Pow(float64(num), 2)))
		}
	}
	return result
}

func SquareOrSquareRoot1(arr []int) []int {
	result := make([]int, len(arr))
	for index, n := range arr {
		whole, remainder := math.Modf(math.Sqrt(float64(n)))
		if remainder == 0 {
			result[index] = int(whole)
		} else {
			result[index] = n * n
		}
	}
	return result
}
