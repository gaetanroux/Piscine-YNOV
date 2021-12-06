package piscine

// ForEach oui
func ForEach(f func(int), arr []int) {
	for _, v := range arr {
		f(v)
	}

}
