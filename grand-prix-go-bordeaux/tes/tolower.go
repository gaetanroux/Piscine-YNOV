package piscine

// ToLower function capitalizes each letter of string.
func ToLower(s string) string {
	runes := []rune(s)

	var newr rune
	for i, r := range runes {
		if r >= 'A' && r <= 'Z' {
			newr = rune(r + 32)
			runes[i] = newr
		}
	}

	return string(runes)
}
