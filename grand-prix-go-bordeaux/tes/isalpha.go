package piscine

func IsAlpha(s string) bool {

	runes := []rune(s)

	var isAlpha bool

	for _, r := range runes {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')) {
			isAlpha = false
			return isAlpha
		}
		isAlpha = true
	}
	return isAlpha
}

