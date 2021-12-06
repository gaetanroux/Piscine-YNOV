package main

import (
	"github.com/01-edu/z01"
)

type numbers struct {
	x int
	y int
}

func setPoint(ptr *numbers) {
	ptr.x = 42
	ptr.y = 21
}

// PrintInt oui
func PrintInt(a int) {
	r := '0'
	if a/10 > 0 {
		PrintInt(a / 10)
	}
	for i := 0; i < a%10; i++ {
		r++
	}
	z01.PrintRune(r)
}

// PrintStr oui
func PrintStr(str string) {
	arrayStr := []rune(str)

	for i := range arrayStr {
		z01.PrintRune(arrayStr[i])
	}
}

func main() {
	points := numbers{}

	setPoint(&points)
	PrintStr("x = ")
	PrintInt(points.x)
	PrintStr(", y = ")
	PrintInt(points.y)

	z01.PrintRune('\n')
}
