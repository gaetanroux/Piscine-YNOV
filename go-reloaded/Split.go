package main

// Split OK
func Split(str, charset string) []string {

	answer := []string{}
	word := ""
	for i := 0; i < len(str); i++ {
		if isCharset(str, charset, i) && i < len(str)-1 {
			if word != "" {
				answer = app(answer, word)
				word = ""
				i = i + len(charset) - 1
			}
		} else {
			word = word + string(str[i])
		}
	}
	if word != "" {
		answer = app(answer, word)
	}
	return answer
}

func app(arr []string, str string) []string {
	arr2 := make([]string, len(arr)+1)
	for i := 0; i <= len(arr)-1; i++ {
		arr2[i] = arr[i]
	}
	arr2[len(arr2)-1] = str
	return arr2
}
func isCharset(str, charset string, i int) bool {
	j := 0
	for j < len(charset) && i < len(str) {
		if str[i] != charset[j] {
			return false
		}
		i++
		j++
	}
	return true
}
