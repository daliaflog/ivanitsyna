package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var result strings.Builder
	var cacheVal *rune
	runes := []rune(str)

	for i := 0; i < len(runes); i++ {
		r := runes[i]

		if unicode.IsDigit(r) {
			if cacheVal == nil {
				return "", ErrInvalidString
			}
			if count := int(r - '0'); count > 0 {
				result.WriteString(strings.Repeat(string(*cacheVal), count))
			}
			cacheVal = nil
			continue
		}

		if cacheVal != nil {
			result.WriteRune(*cacheVal)
		}

		r, newI, err := resolveRune(runes, i)
		if err != nil {
			return "", err
		}
		i = newI
		cacheVal = &r
	}

	if cacheVal != nil {
		result.WriteRune(*cacheVal)
	}

	return result.String(), nil
}

// resolveRune получаем значение с учётом экранирования.
func resolveRune(runes []rune, i int) (rune, int, error) {
	r := runes[i]
	if r != '\\' {
		return r, i, nil
	}

	i++
	if i >= len(runes) {
		return 0, i, ErrInvalidString
	}

	next := runes[i]
	if next != '\\' && !unicode.IsDigit(next) {
		return 0, i, ErrInvalidString
	}
	return next, i, nil
}
