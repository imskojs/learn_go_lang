package kyu8

func BoolToWord(word bool) string {
	switch {
	case word == true:
		return "Yes"
	default:
		return "No"
	}
}
