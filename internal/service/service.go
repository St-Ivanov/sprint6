package service

import (
	"strings"
	"unicode"

	"github.com/St-Ivanov/sprint6/pkg/morse"
)

func DataConversion(data string) string {
	if len(data) == 0 {
		return ""
	}
	if strings.ContainsFunc(data, unicode.IsLetter) {
		return morse.ToMorse(data)
	}
	return morse.ToText(data)
}
