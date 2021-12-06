package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	var sortedargs []string

	for i := range arguments {
		if i > 0 {
			sortedargs = append(sortedargs, arguments[i])
		}
	}

	lenargs := 0
	for i := range sortedargs {
		lenargs = i + 1
	}

	for j := 0; j < lenargs; j++ {
		for i := 0; i < lenargs-1; i++ {
			if sortedargs[i] > sortedargs[i+1] {
				temp := sortedargs[i+1]
				sortedargs[i+1] = sortedargs[i]
				sortedargs[i] = temp
			}
		}
	}

	for i := range sortedargs {
		runes := []rune(sortedargs[i])
		for j := range runes {
			z01.PrintRune(runes[j])
		}
		z01.PrintRune('\n')
	}
}
