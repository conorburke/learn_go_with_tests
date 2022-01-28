package iteration

func Iterate(character string, count int) (output string) {
	for i := 0; i < count; i++ {
		output += character
	}

	return output
}
