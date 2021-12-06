package piscine

//IterativePower returns the power of the int passed as parameter
func IterativePower(nb int, power int) int {

	pwr := 1

	if power == 0 {

		return 1

	} else if power > 0 {

		for i := 1; i <= power; i++ {

			pwr = pwr * nb

		}

		return pwr

	} else {

		return 0

	}

}
