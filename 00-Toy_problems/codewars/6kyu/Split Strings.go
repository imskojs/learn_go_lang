package kyu6

func Solution(str string) []string {
	var result []string
	if len(str)%2 != 0 {
		str += "_"
	}
	for i := 0; i < len(str); i += 2 {
		result = append(result, str[i:i+2])
	}
	return result

}

func SolutionOld(str string) []string {
	var result []string
	var buff string
	var lastLetter string

	if !isEven(len(str)) {
		lastLetter = str[len(str)-1:] + "_"
		str = str[:len(str)-1]
	}

	for i, letter := range str {
		buff += string(letter)
		if isEvenPosition(i) {
			result = append(result, buff)
			buff = ""
		}
	}

	if lastLetter != "" {
		result = append(result, lastLetter)
	}

	return result
}

func isEven(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}

func isEvenPosition(i int) bool {
	return !isEven(i)
}
