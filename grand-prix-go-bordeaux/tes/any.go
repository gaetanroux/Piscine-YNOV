package piscine

//Any returns true, for a string slice
func Any(f func(string) bool, a []string) bool {

	for _, number := range a {

		answer := f(number)
		if answer {

			return true
		}
	}

	return false
}
