package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var result strings.Builder

	for i := 0; i < len(s); i++ {
		char := rune(s[i])

		if unicode.IsDigit(char) {
			return "", ErrInvalidString
		}

		repeat := 1
		if i+1 < len(s) && unicode.IsDigit(rune(s[i+1])) {
			repeat, _ = strconv.Atoi(string(s[i+1]))
			i++
		}

		result.WriteString(strings.Repeat(string(char), repeat))
	}

	return result.String(), nil
}
