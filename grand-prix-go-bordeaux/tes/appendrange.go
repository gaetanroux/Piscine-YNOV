package piscine

// AppendRange function returns a slice of int
// with all the values between min and max.
// Min is included, and max is excluded.
// If min is superior or equal to max, a nil slice is returned.
func AppendRange(min, max int) []int {
	var result []int
	if max > min {
		for i := min; i < max; i++ {
			result = append(result, i)
		}
	} else {
		return result
	}
	return result
}
