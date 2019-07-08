package kyu8

func PositiveSum(numbers []int) int {
	var total int
	for _, num := range numbers {
		if num > 0 {
			total += num
		}
	}
	return total
}

func PositiveSum1(numbers []int) (sum int) {
	for _, num := range numbers {
		if num > 0 {
			sum += num
		}
	}
	return
}
