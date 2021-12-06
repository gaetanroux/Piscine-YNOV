package abc

// FS choose the police of the art
func FS(file string) string {

	police := "standard.txt"
	if file == "shadow" {

		police = "shadow.txt"
	} else if file == "thinkertoy" {

		police = "thinkertoy.txt"
	}

	return police
}
