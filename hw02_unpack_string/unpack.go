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
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		char := runes[i]

		if unicode.IsDigit(char) {
			return "", ErrInvalidString
		}

		repeat := 1
		if i+1 < len(runes) && unicode.IsDigit(runes[i+1]) {
			repeat, _ = strconv.Atoi(string(runes[i+1]))
			i++
		}

		result.WriteString(strings.Repeat(string(char), repeat))
	}

	return result.String(), nil
}
