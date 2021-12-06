package piscine

//isPrime says if a number is prime
func isPrime(a int) bool {

	if a == 1 {

		return false
	}

	if a == 2 {

		return false
	}

	for i := 3; i < a; i++ {

		if a%i != 0 {

			return false
		}
	}

	return true
}

//Map for an int slice, applies a function of this type func(int) bool on each elements of that slice and returns a slice of all the return values.
func Map(f func(int) bool, a []int) []bool {

	var tabbool []bool
	for _, number := range a {

		tabbool = append(tabbool, f(number))
	}

	return tabbool
}
