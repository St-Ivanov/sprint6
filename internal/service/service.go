package service

import (
	"github.com/St-Ivanov/sprint6/internal/errors"
	"github.com/St-Ivanov/sprint6/pkg/morse"
)

func DataConversion(data string) (string, error) {
	if len(data) == 0 {
		return "", errors.ErrEmptyValue
	}
	if isMorseCode(data) {
		return morse.ToText(data), nil
	}
	return morse.ToMorse(data), nil
}

// Проверка строка содержит азбуку морзе.
func isMorseCode(data string) bool {
	for _, val := range data {
		if val != rune('.') && val != rune('-') && val != rune(' ') {
			return false
		}
	}
	return true
}
