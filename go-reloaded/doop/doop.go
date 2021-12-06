package main

import (
	"github.com/01-edu/z01"
	"os"
)

//validArg checks if the argument is an int
func validArg(s string) bool {

	validarg := true
	for i, v := range s {
		if i == 0 {
			if v != '-' && v != '+' && (v <= 47 || v >= 58) {
				return false
			}
		}
		if i != 0 {
			if v <= 47 && v >= 58 {
				return false
			}
		}
	}

	return validarg
}

//Atoi simulates the behaviour of the Atoi function in Go. Atoi transforms a number represented as a string in a number represented as an int.
func Atoi(s string) int {
	str := 0
	signCount := 0
	addMinus := false
	for i, v := range s {
		if int(rune(v)) > 47 && int(rune(v)) < 58 || v == '+' || v == '-' {
			if (v == '-' || v == '+') && i > 0 {
				return 0
			}
			if v == '-' && i == 0 {
				addMinus = true
			}
			if v == '+' || v == '-' {
				signCount++
				if signCount > 1 {
					return 0
				}
				continue
			}
			a := 0
			for i := '1'; i <= v; i++ {
				a++
			}
			str = str*10 + a
		} else {
			return 0
		}

	}
	if addMinus {
		str = str - (str * 2)
	}
	return str

}


func calcul(a int, b string, c int) {

	result := 0

	if a > -500000 && a < 500000 {
		if b == "+" {

			result = a + c
		} else if b == "-" {

			result = a - c
		} else if b == "*" {

			result = a * c
		} else if b == "/" {

			if c != 0 {

				result = a / c
			} else {

				div0 := "No division by "
				for _, v := range div0 {

					z01.PrintRune(v)
				}
			}
		} else if b == "%" {

			if c != 0 {

				result = a % c
			} else {

				mod0 := "No modulo by "
				for _, v := range mod0 {

					z01.PrintRune(v)
				}
			}
		} else {

			result = 0
		}
	}

	printer := numberToString(result)
	for _, v := range printer {

		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

//numberToString change an int into a string
func numberToString(number int) string {

	var s string
	sign := true

	if number == -9223372036854775808  {

		s = "0"
		return s
	}

	if number < 0 {

		number = number * -1
		sign = false
	}

	for {

		mod := number % 10
		s = string(mod+'0') + s
		number = number / 10
		if number == 0 {
			break
		}
	}

	if sign == false {

		s = "-" + s
	}

	return s
}

func main() {

	args := os.Args
	if len(args) == 4 {

		if validArg(args[1]) {

			if validArg(args[3]) {

				nb1 := Atoi(args[1])
				ope := args[2]
				nb2 := Atoi(args[3])

				calcul(nb1, ope, nb2)

			} else {

				z01.PrintRune('0')
				z01.PrintRune('\n')
			}
		} else {

			z01.PrintRune('0')
			z01.PrintRune('\n')
		}
	}

}
