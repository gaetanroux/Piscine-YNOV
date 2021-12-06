package piscine

// MakeRange function returns a slice of int
// with all the values between min and max.
// Min is included, and max is excluded.
// If min is superior or equal to max, a nil slice is returned.
func MakeRange(min, max int) []int {
	var result []int
	if max > min {
		result = make([]int, max-min)
		for i := 0; i < max-min; i++ {
			result[i] = min + i
		}
	} else {
		return result
	}
	return result
}
