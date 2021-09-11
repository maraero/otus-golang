package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var result strings.Builder

	symbol := ""

	for _, char := range input {
		isDigit := unicode.IsDigit(char)
		isPreviuousSymbolEmpty := symbol == ""

		switch {
		case isDigit && isPreviuousSymbolEmpty:
			return "", ErrInvalidString
		case isDigit && !isPreviuousSymbolEmpty:
			mult, _ := strconv.Atoi(string(char))

			result.WriteString(strings.Repeat(symbol, mult))

			symbol = ""
		case !isDigit && isPreviuousSymbolEmpty:
			symbol = string(char)
		case !isDigit && !isPreviuousSymbolEmpty:
			result.WriteString(symbol)

			symbol = string(char)
		}
	}

	if symbol != "" {
		result.WriteString(symbol)
	}

	return result.String(), nil
}
