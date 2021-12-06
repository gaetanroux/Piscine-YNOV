package printcomb2

import "github.com/01-edu/z01"

func PrintComb2() {
	var a, b, c, d rune = '0', '0', '0', '0'
	for i := a; i <= 57; i++ {
		for j := b; j <= 57; j++ {
			for k := c; k <= 57; k++ {
				for l := d; l <= 57; l++ {
					if k > i {
						PrintNumbers(i, j, k, l)
					} else if k == i && l > j {
						PrintNumbers(i, j, k, l)
					}
				}
			}
		}
	}
}

func PrintNumbers(i, j, k, l rune) {
	z01.PrintRune(i)
	z01.PrintRune(j)
	z01.PrintRune(32)
	z01.PrintRune(k)
	z01.PrintRune(l)
	if i == 57 && j == 56 && k == 57 && l == 57 {
		z01.PrintRune(10)
	} else {
		z01.PrintRune(44)
		z01.PrintRune(32)
	}
}
