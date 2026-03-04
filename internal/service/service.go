package service

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func AutoDetect(s string) (string, error) {
	value := strings.TrimSpace(s)

	if value == "" {
		return "", errors.New("empty string")
	}

	isNotMorse := func(c rune) bool {
		return c != ' ' && c != '-' && c != '.'
	}

	index := strings.IndexFunc(value, isNotMorse)

	if index >= 0 {
		for _, v := range value {
			if v == ' ' {
				continue
			}
			upperR := unicode.ToUpper(v)
			if _, ok := morse.DefaultMorse[upperR]; !ok {
				return "", fmt.Errorf("unknown symbol: %q", v)
			}
		}
		return morse.ToMorse(value), nil
	}

	return morse.ToText(value), nil
}
