package main

//RecursivePower returns the power of the int passed as parameter.
func RecursivePower(nb int, power int) int {

	if power == 1 {

		return nb
	}

	if power < 0 {

		return 0
	}

	if power > 0 {

		return nb * RecursivePower(nb, power-1)
	}

	return 1
}
