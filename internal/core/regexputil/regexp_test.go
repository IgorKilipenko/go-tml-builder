package regexputil_test

import (
	"fmt"
	"testing"

	"github.com/IgorKilipenko/go-tml-builder/internal/core/regexputil"
)

func TestExpressionOrFunc(t *testing.T) {
	tests := []struct {
		name         string
		input        []fmt.Stringer
		preprocessor func(fmt.Stringer) string
		expected     string
	}{
		{
			name:     "nil slice",
			input:    nil,
			expected: "",
		},
		{
			name:     "empty slice",
			input:    []fmt.Stringer{},
			expected: "",
		},
		{
			name:     "single item (default stringer)",
			input:    []fmt.Stringer{testStringer{"a"}},
			expected: "a",
		},
		{
			name:     "multiple items (default stringer)",
			input:    []fmt.Stringer{testStringer{"a"}, testStringer{"b"}},
			expected: "a|b",
		},
		{
			name:         "custom preprocessor",
			input:        []fmt.Stringer{testStringer{"a"}, testStringer{"b"}},
			preprocessor: func(s fmt.Stringer) string { return s.String() + "!" },
			expected:     "a!|b!",
		},
		{
			name:     "with empty values",
			input:    []fmt.Stringer{testStringer{""}, testStringer{"b"}, testStringer{""}},
			expected: "|b|",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := regexputil.ExpressionOrFunc(tt.input, tt.preprocessor)
			if got != tt.expected {
				t.Errorf("got %q, want %q", got, tt.expected)
			}
		})
	}
}

// Вспомогательный тип для тестов
type testStringer struct{ s string }

func (t testStringer) String() string { return t.s }
