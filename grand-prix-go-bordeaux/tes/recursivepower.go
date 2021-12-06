package piscine

// RecursivePower function returns the power
// of the int passed as parameter.
// Negative powers will return 0.
// Overflows do not have to be dealt with.
func RecursivePower(nb int, power int) int {
	result := nb
	if power > 0 {
		result = result * RecursivePower(nb, power-1)
	} else if power == 0 {
		result = 1
	} else {
		result = 0
	}
	return result
}
