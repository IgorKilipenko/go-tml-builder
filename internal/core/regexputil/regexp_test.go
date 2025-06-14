package regexputil_test

import (
	"testing"

	"github.com/IgorKilipenko/go-tml-builder/internal/core/regexputil"
	"github.com/stretchr/testify/assert"
)

func TestEscapeUnescape(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Simple", `\w+`, `\\w+`},
		{"DoubleEscape", `\\`, `\\\\`},
		{"Mixed", `\d\\s\S`, `\\d\\\\s\\S`},
		{"NoEscape", "abc", "abc"},
		{"Complex", `([A-Za-zА-ЯЁа-яё]+)\s*=\s*"([^"]*)"`, `([A-Za-zА-ЯЁа-яё]+)\\s*=\\s*\"([^\"]*)\"`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Тест экранирования
			escaped := regexputil.EscapeForJSON(tt.input)
			assert.Equal(t, tt.expected, escaped)

			// Тест обратного преобразования
			unescaped := regexputil.UnescapeFromJSON(escaped)
			assert.Equal(t, tt.input, unescaped)
		})
	}
}

func TestEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty", "", ""},
		{"SingleBackslash", `\`, `\\`},
		{"MultipleBackslashes", `\\\`, `\\\\\\`},
		{"MixedWithQuotes", `\"Hello\"`, `\\\"Hello\\\"`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			escaped := regexputil.EscapeForJSON(tt.input)
			assert.Equal(t, tt.expected, escaped)
			assert.Equal(t, tt.input, regexputil.UnescapeFromJSON(escaped))
		})
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		pattern string
		isValid bool
	}{
		{`\w+`, true},
		{`[a-z`, false}, // Невалидное regexp
	}

	for _, tt := range tests {
		err := regexputil.Validate(tt.pattern)
		if tt.isValid {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}

func TestUnicodeUnescape(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`\\\\→`, `\\→`},   // Экранированный слеш + Unicode
		{`\\日本`, `\日本`},    // Иероглифы
		{`\\ы\\я`, `\ы\я`}, // Кириллица
		{`\\🦄`, `\🦄`},      // Emoji
		{`\\★`, `\★`},      // Символ звезды
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.expected, regexputil.UnescapeFromJSON(tt.input))
		})
	}
}

func TestEscapeForJSON_Unicode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`→`, `→`},
		{`\\→`, `\\\\→`},
		{`\"日本\"`, `\\\"日本\\\"`},
		{`🦄`, `🦄`},
		{`\\★`, `\\\\★`},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.expected, regexputil.EscapeForJSON(tt.input))
		})
	}
}
