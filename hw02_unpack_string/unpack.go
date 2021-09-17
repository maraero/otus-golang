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
		isCurrentSymbolDigit := unicode.IsDigit(char)
		isPreviousSymbolEmpty := symbol == ""

		switch {
		case isCurrentSymbolDigit && isPreviousSymbolEmpty:
			return "", ErrInvalidString
		case isCurrentSymbolDigit && !isPreviousSymbolEmpty:
			mult, e := strconv.Atoi(string(char))

			if e != nil {
				return "", e
			}

			result.WriteString(strings.Repeat(symbol, mult))

			symbol = ""
		case !isCurrentSymbolDigit && isPreviousSymbolEmpty:
			symbol = string(char)
		case !isCurrentSymbolDigit && !isPreviousSymbolEmpty:
			result.WriteString(symbol)

			symbol = string(char)
		}
	}

	if symbol != "" {
		result.WriteString(symbol)
	}

	return result.String(), nil
}
