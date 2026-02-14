package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Convert автоматически определяет тип входных данных (текст или код Морзе) и выполняет конвертацию.
func Convert(input string) (string, error) {
	if strings.TrimSpace(input) == "" {
		return "", nil
	}

	if isMorseCode(input) {
		return morse.ToText(input), nil
	}

	return morse.ToMorse(input), nil
}

// isMorseCode проверяет, является ли строка кодом Морзе.
// Строка должна содержать хотя бы один символ точки или тире.
func isMorseCode(s string) bool {
	if s == "" {
		return false
	}

	hasMorseChar := false
	for _, r := range s {
		if r != '.' && r != '-' && r != ' ' && r != '\t' && r != '\n' && r != '\r' {
			return false
		}
		if r == '.' || r == '-' {
			hasMorseChar = true
		}
	}

	return hasMorseChar
}
