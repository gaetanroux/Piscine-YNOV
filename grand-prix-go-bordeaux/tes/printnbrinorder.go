package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n > 0 {
		finalslice := SortIntTable(ToSlice(n))
		for i := range finalslice {
			for j, k := 0, '0'; j <= 9; j, k = j+1, k+1 {
				if finalslice[i] == j {
					z01.PrintRune(k)
				}
			}
		}
	} else {
		z01.PrintRune('0')
	}
}

func ToSlice(n int) []int {
	var digits []int

	for n > 0 {
		if n == 0 {
			digits = append(digits, 0)
		} else {
			mod := n % 10
			digits = append(digits, mod)
		}
		n = n / 10
	}
	return digits
}

func SortIntTable(table []int) []int {
	tablelen := 0
	for i := range table {
		tablelen = i + 1
	}

	var temp int

	for j := 0; j < tablelen; j++ {
		for i := 0; i < (tablelen - 1); i++ {
			if table[i] > table[i+1] {
				temp = table[i+1]
				table[i+1] = table[i]
				table[i] = temp
			}
		}
	}
	return table
}
