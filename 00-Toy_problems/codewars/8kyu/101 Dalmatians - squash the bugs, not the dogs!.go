package kyu8

func HowManyDalmatians(number int) string {
	switch {
	case number == 101:
		return "101 DALMATIONS!!!"
	case number > 50:
		return "Woah that's a lot of dogs!"
	case number > 10:
		return "More than a handful!"
	default:
		return "Hardly any"
	}
}
