package piscine

// IsNumeric oui
func IsNumeric(s string) bool {
	runes := []rune(s)
	var isNum bool

	for _, r := range runes {
		if !(r >= '0' && r <= '9') {
			isNum = false
			return isNum
		}
		isNum = true
	}
	return isNum
}
