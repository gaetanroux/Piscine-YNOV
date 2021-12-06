package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args

	lenargs := 0
	for i := range arguments {
		lenargs = i + 1
	}

	for i := lenargs - 1; i > 0; i-- {
		runes := []rune(arguments[i])
		for j := range runes {
			z01.PrintRune(runes[j])
		}
		z01.PrintRune('\n')
	}
}
