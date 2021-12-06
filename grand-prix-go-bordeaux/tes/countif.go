package piscine

// CountIf that returns the number of elements of a string slice for which the f function returns true
func CountIf(f func(string) bool, tab []string) int {
	count := 0

	for _, number := range tab {

		answer := f(number)
		if answer {

			count++
		}
	}

	return count
}
