package piscine

func Index(s string, toFind string) int {
	substr := []rune(toFind)
	str := []rune(s)

	lenstr := 0
	for i := range str {
		lenstr = i + 1
	}

	lensubstr := 0
	for j := range substr {
		lensubstr = j + 1
	}

	switch {
	case lensubstr == 0:
		return 0
	case lensubstr == 1:
		return IndexRune(str, substr[0])
	case lensubstr == lenstr:
		if toFind == s {
			return 0
		}
		return -1
	case lensubstr > lenstr:
		return -1
	case lensubstr < lenstr:
		c0 := substr[0]
		c1 := substr[1]
		i := 0
		t := lenstr - lensubstr + 1
		for i < t {
			if str[i] != c0 {
				o := IndexRune(str, c0)
				if o < 0 {
					return -1
				}
				i += o
			}
			if str[i+1] == c1 && s[i:i+lensubstr] == toFind {
				return i
			}
			i++
		}
	}
	return -1
}

// IndexRune function returns the index of
// the first occurrence of a rune in a given string.
func IndexRune(s []rune, r rune) int {
	indexrune := -1
	for i := range s {
		if s[i] == r {
			indexrune = i
			return indexrune
		}
	}
	return indexrune
}
