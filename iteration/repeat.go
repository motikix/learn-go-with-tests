package iteration

func Repeat(character string, repeat uint) (repeated string) {
	for i := uint(0); i < repeat; i++ {
		repeated += character
	}

	return
}
