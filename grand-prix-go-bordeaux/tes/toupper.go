package piscine

// ToUpper function capitalizes each letter of string.
func ToUpper(s string) string {
	runes := []rune(s)

	var newr rune
	for i, r := range runes {
		if r >= 'a' && r <= 'z' {
			newr = rune(r - 32)
			runes[i] = newr
		}
	}

	return string(runes)
}
