package piscine

// Capitalize function make the first letter
// of each word uppercase and lowercases the rest
// of each word of a string.
// A word is a sequence of alphanumerical characters.
func Capitalize(s string) string {
	runes := []rune(s)

	strlen := 0
	for i := range runes {
		strlen = i + 1
	}

	for i := 0; i < strlen; i++ {
		if i != 0 && (!((runes[i-1] >= 'a' && runes[i-1] <= 'z') || (runes[i-1] >= 'A' && runes[i-1] <= 'Z') || (runes[i-1] >= '0' && runes[i-1] <= '9'))) {
			if runes[i] >= 'a' && runes[i] <= 'z' {
				runes[i] = rune(runes[i] - 32)
			}
		} else if i == 0 {
			if runes[i] >= 'a' && runes[i] <= 'z' {
				runes[i] = rune(runes[i] - 32)
			}
		} else {
			if runes[i] >= 'A' && runes[i] <= 'Z' {
				runes[i] = rune(runes[i] + 32)
			}
		}
	}
	return string(runes)
}
