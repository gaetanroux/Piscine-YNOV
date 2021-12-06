package piscine

// ConcatParams oui
func ConcatParams(args []string) string {

	var result string
	lengthargs := 0
	for i := range args {
		lengthargs = i
	}

	for j, word := range args {
		if j == lengthargs {
			result += word
		} else {
			result += word + "\n"
		}

	}
	return result
}
