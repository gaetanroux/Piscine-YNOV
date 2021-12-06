package piscine

// IsPrintable function returns true if the string passed in parameter
// only contains printable characters, and returns false otherwise.
func IsPrintable(s string) bool {
	runes := []rune(s)
	var isOnlyPrint bool

	for _, r := range runes {
		if r >= 0 && r <= 31 {
			isOnlyPrint = false
			return isOnlyPrint
		}
		isOnlyPrint = true
	}
	return isOnlyPrint
}
