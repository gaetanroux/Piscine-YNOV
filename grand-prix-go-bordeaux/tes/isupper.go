package piscine

func IsUpper(s string) bool {

	runes := []rune(s)
	var isUppercase bool

	for _, r := range runes {
		if !(r >= 'A' && r <= 'Z') {
			isUppercase = false
			return isUppercase
		}
		isUppercase = true
	}
	return isUppercase
}
