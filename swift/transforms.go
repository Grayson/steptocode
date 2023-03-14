package swift

import "unicode"

func CreateIdentifierName(input string) string {
	output := make([]rune, len(input))
	index := 0
	shouldCapitalizeNextChar := false
	for _, char := range input {

		if unicode.IsPunct(char) {
			char = '_'
		}

		if !(unicode.IsLetter(char) || unicode.IsDigit(char) || char == '_') {
			shouldCapitalizeNextChar = unicode.IsSpace(char)
			continue
		}

		if shouldCapitalizeNextChar {
			char = unicode.ToUpper(char)
			shouldCapitalizeNextChar = false
		}
		output[index] = char
		index++
	}

	for ; index > 1 && output[index-1] == '_'; index-- {
	}

	return string(output[:index])
}
