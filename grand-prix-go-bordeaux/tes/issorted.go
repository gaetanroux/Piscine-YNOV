package piscine

// TestSort oui
func TestSort(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else if a < b {
		return -1
	}
	return -1
}

// IsSorted oui
func IsSorted(f func(a, b int) int, tab []int) bool {
	tabLength := 0
	for range tab {
		tabLength++
	}
	left := false
	right := false
	for i := 0; i < tabLength-1; i++ {
		if f(tab[i], tab[i+1]) > 0 {
			left = true
		} else if f(tab[i], tab[i+1]) < 0 {
			right = true
		}
	}
	if left && right {
		return false
	} else {
		return true
	}
}
