package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {

	args := os.Args

	// Checks if there is any arguments
	if len(args) > 1 {

		// Merge them into one big string
		mergeArgs := ""

		for i := 1; i <= len(args)-1; i++ {

			mergeArgs += args[i] + " "
		}

		// Create a string array with that string
		stringArray := make([]string, len(mergeArgs))

		for i, v := range mergeArgs {

			stringArray[i] = string(v)
		}

		// Count all vowels
		countVowel := 0

		for _, v := range mergeArgs {

			if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' || v == 'A' || v == 'E' || v == 'I' || v == 'O' || v == 'U' {

				countVowel++
			}
		}

		// Loop that mirrors the position of the vowels (divided by 2 because 2 vowels are moved)
		for i := 0; i < countVowel/2; i++ {

			// Loop that acts in the first part of the string and stock the vowel which will move forward
			var firstVowel string
			limitPlus := 0

			for _, v := range mergeArgs {

				if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' || v == 'A' || v == 'E' || v == 'I' || v == 'O' || v == 'U' {

					firstVowel = string(v)
					limitPlus++

					if limitPlus == i+1 {

						limitPlus = 0
						break
					}
				}
			}

			// Loop that acts in the second part of the string and stock the vowel which will move backward
			var lastVowel string
			reachLimit := 0

			for _, v := range mergeArgs {

				if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' || v == 'A' || v == 'E' || v == 'I' || v == 'O' || v == 'U' {

					lastVowel = string(v)
					reachLimit++

					if reachLimit == countVowel-i {

						reachLimit = 0
						break
					}
				}
			}

			// Replace the first part of the array by the last part
			replaceFirst := 0

			for j, v := range stringArray {

				if v == "a" || v == "e" || v == "i" || v == "o" || v == "u" || v == "A" || v == "E" || v == "I" || v == "O" || v == "U" {

					if replaceFirst == i {

						stringArray[j] = lastVowel
						break
					}

					replaceFirst++
				}
			}

			// Replace the last part of the array by the first part
			replaceLast := 0

			for j, v := range stringArray {

				if v == "a" || v == "e" || v == "i" || v == "o" || v == "u" || v == "A" || v == "E" || v == "I" || v == "O" || v == "U" {

					if replaceLast == countVowel-i-1 {

						stringArray[j] = firstVowel
						break
					}

					replaceLast++
				}
			}
		}

		// Prints the array
		for i := range stringArray {

			for _, v := range stringArray[i] {

				z01.PrintRune(v)
			}
		}
		
	} else {

		z01.PrintRune('\n')
	}
}
