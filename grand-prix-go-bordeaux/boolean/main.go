package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {

	for _, r := range s {

		z01.PrintRune(r)
	}

	z01.PrintRune('\n')
}

func isEven(nbr int) bool {

	even := nbr % 2
	if even != 1 {

		return true
	}

	return false

}

func main() {

	args := os.Args
	lengthOfArg := len(args) - 1

	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"

	if isEven(lengthOfArg) {

		printStr(EvenMsg)
	} else {

		printStr(OddMsg)
	}
}
