package piscine

import "github.com/01-edu/z01"

// PrintWordsTables function prints the words of a string array
// that will be returned by a function SplitWhiteSpaces (provided by tester).
// Each word is on a single line. Each word ends with a \n.
func PrintWordsTables(table []string) {
	for i := range table {
		runes := []rune(table[i])
		for j := range runes {
			z01.PrintRune(runes[j])
		}
		z01.PrintRune('\n')
	}
}
