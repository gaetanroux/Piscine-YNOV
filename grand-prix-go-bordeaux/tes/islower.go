package piscine

func IsLower(s string) bool {

	runes := []rune(s)
	var isLowercase bool

	for _, r := range runes {
		if !(r >= 'a' && r <= 'z') {
			isLowercase = false
			return isLowercase
		}
		isLowercase = true
	}
	return isLowercase
}
