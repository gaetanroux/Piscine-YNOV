package main

//AtoiBase  takes a string number and its string base in parameters and returns its conversion as an int.
func AtoiBase(s string, base string) int {

	number := 0
	lenbase := len(base)
	lens := len(s)
	if lenbase < 2 {

		return 0
	}

	for _, v := range base {

		unique := 0
		for _, w := range base {

			if v == w {

				unique++
			}
		}

		if unique != 1 {

			return 0
		}

		if v == '+' || v == '-' {
			return 0
		}

	}

	for _, v := range s {
		stock := 0

		for _, w := range base {
			if v != w {
				stock++
			} else {
				break
			}
		}

		addnumber := 1
		for k := lens - 1; k > 0; k-- {
			addnumber = addnumber * lenbase
		}
		number = number + stock*addnumber
		lens = lens - 1

	}

	return number
}
