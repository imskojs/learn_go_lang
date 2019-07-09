package kyu8

func BoolToWord(word bool) string {
	switch {
	case word == true:
		return "Yes"
	default:
		return "No"
	}
}

func BoolToWord1(word bool) string {
	return map[bool]string{false: "No", true: "Yes"}[word]
}
