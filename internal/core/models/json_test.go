package models_test

import (
	"encoding/json"
	"testing"

	"github.com/IgorKilipenko/go-tml-builder/internal/core/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRule_JSON_Marshaling(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Simple regex",
			input:    `\w+`,
			expected: `{"match":"\\w+"}`,
		},
		{
			name:     "Escaped quotes",
			input:    `\"[^"]+\"`,
			expected: `{"match":"\\\"[^\"]+\\\""}`,
		},
		{
			name:     "Complex pattern",
			input:    `\\d{3}\s\p{L}+`,
			expected: `{"match":"\\\\d{3}\\s\\p{L}+"}`,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Тест маршалинга
			rule := models.Rule{Match: tc.input}
			jsonData, err := json.Marshal(rule)
			require.NoError(t, err)
			assert.JSONEq(t, tc.expected, string(jsonData))

			// Тест демаршалинга
			var decoded models.Rule
			err = json.Unmarshal(jsonData, &decoded)
			require.NoError(t, err)
			assert.Equal(t, tc.input, decoded.Match)
		})
	}
}

// Тест для полной структуры Rule
func TestFullRule_JSON(t *testing.T) {
	rule := models.Rule{
		Name:  "test",
		Match: `\d+`,
		Begin: `\(`,
	}

	jsonData, err := json.Marshal(rule)
	require.NoError(t, err)

	var decoded models.Rule
	require.NoError(t, json.Unmarshal(jsonData, &decoded))
	assert.Equal(t, rule, decoded)
}
