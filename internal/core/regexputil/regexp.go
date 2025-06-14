package regexputil

import (
	"regexp"
	"strings"
)

// EscapeForJSON экранирует обратные слеши для JSON
func EscapeForJSON(pattern string) string {
	if !strings.ContainsAny(pattern, `\"`) {
		return pattern
	}

	var escaped strings.Builder
	escaped.Grow(len(pattern) + strings.Count(pattern, `\`) + strings.Count(pattern, `"`))

	for _, r := range pattern {
		switch r {
		case '\\', '"':
			escaped.WriteByte('\\')
			escaped.WriteRune(r)
		default:
			escaped.WriteRune(r)
		}
	}

	return escaped.String()
}

// UnescapeFromJSON убирает экранирование
func UnescapeFromJSON(escaped string) string {
	var buf strings.Builder
	buf.Grow(len(escaped)) // Предварительное выделение памяти

	runes := []rune(escaped) // Конвертируем в руны для корректной обработки Unicode
	n := len(runes)

	for i := 0; i < n; i++ {
		switch runes[i] {
		case '\\':
			if i+1 < n {
				switch runes[i+1] {
				case '\\', '"':
					buf.WriteRune(runes[i+1])
					i++ // Пропускаем следующий символ
				default:
					// Оставляем некорректную escape-последовательность как есть
					buf.WriteRune('\\')
					buf.WriteRune(runes[i+1])
					i++
				}
			} else {
				buf.WriteRune('\\') // Одиночный бэкслеш в конце строки
			}
		default:
			buf.WriteRune(runes[i])
		}
	}

	return buf.String()
}

// Validate проверяет валидность регулярного выражения
func Validate(pattern string) error {
	_, err := regexp.Compile(pattern)
	return err
}
