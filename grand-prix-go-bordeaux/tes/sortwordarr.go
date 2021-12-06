package piscine

//SortWordArr sorts by ascii (in ascending order) a string slice.
func SortWordArr(a []string) {
	len := 0

	var s string
	for range a {
		len++
	}
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {

			if a[j] >= a[j+1] {
				s = a[j]
				a[j] = a[j+1]
				a[j+1] = s
			}
		}
	}
}
