package converters

func TranslateToRuneSlice(bytes []byte) []rune {
	var text []rune

	for i := range bytes {
		text = append(text, rune(bytes[i]))
	}

	return text
}

func TranslateToByteSlice(runes []rune) []byte {
	var text []byte

	for i := range runes {
		text = append(text, byte(runes[i]))
	}

	return text
}
