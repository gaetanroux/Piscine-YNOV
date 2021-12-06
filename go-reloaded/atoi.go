package main

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
