package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args

	for i := range arguments {
		if i > 0 {
			runes := []rune(arguments[i])
			for j := range runes {
				z01.PrintRune(runes[j])
			}
			z01.PrintRune('\n')
		}
	}
}
