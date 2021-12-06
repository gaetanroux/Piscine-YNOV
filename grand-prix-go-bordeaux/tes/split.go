package piscine

// Split oui
func Split(s, sep string) []string {

	temp := 0
	var tab []string
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			tab = append(tab, s[temp:])
		}
		if s[i] == sep[0] && s[i+len(sep)-1] == sep[len(sep)-1] && s[i+len(sep)-2] == sep[len(sep)-2] {
			tab = append(tab, s[temp:i])
			temp = i + len(sep)

		}
	}
	return tab

}
