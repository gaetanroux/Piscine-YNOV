package piscine

// TrimAtoi function transforms a number within a string
// into a number represented as an int.
func TrimAtoi(s string) int {
	finalslice := AtoiToSlice(s)

	var result []int
	isNegative := false

	for i := range finalslice {
		for j, k := 0, '0'; j <= 9; j, k = j+1, k+1 {
			if finalslice[i] == k {
				result = append(result, j)
			}
			if finalslice[0] == '-' {
				isNegative = true
			}
		}
	}

	finalint := 0

	lenres := 0
	for i := range result {
		lenres = i + 1
	}

	for i := 0; i < lenres; i++ {
		finalint = finalint + result[i]*Power(10, lenres-i-1)
	}

	if isNegative {
		finalint = finalint * (-1)
	}
	return finalint
}

// AtoiToSlice function transforms a number
// represented as a string to a number
// represented as an int.
// The handling of the signs
// + or - is taken into account.
func AtoiToSlice(s string) []rune {
	runes := []rune(s)
	var res []rune

	for i := range runes {
		if runes[i] >= '0' && runes[i] <= '9' || runes[i] == '-' {
			res = append(res, runes[i])
		}
	}
	return res
}

// Power function
func Power(nb int, power int) int {
	result := 1
	if power >= 0 {
		for i := 1; i < (power + 1); i++ {
			result = result * nb
		}
	} else {
		result = 0
	}
	return result
}
