package main

// AdvancedSortWordArr OK
func AdvancedSortWordArr(array []string, f func(a, b string) int) {
	size := 0
	for i := range array {
		size = i + 1
	}
	var temp string
	for i := 0; i < size-1; i++ {
		for j := 0; j < (size - i - 1); j++ {
			if f(array[j], array[j+1]) > 0 {
				temp = array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
	}
}

// Compare OK
func Compare(a, b string) int {
	if a < b {
		return -1
	} else if a == b {
		return 0
	}
	return 1
}

// Compare1 OK
func Compare1(b, a string) int {
	if a > b {
		return 1
	} else if b > a {
		return -1
	}

	return 0
}
