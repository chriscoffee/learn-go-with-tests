package iteration

func Repeat(letter string, repeat int) string {
	var repeated string
	for i := 0; i < repeat; i++ {
		repeated += letter
	}

	return repeated
}
