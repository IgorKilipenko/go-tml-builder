package regexputil

import (
	"fmt"
	"strings"
)

// ExpressionOrFunc формирует строку элементов с разделителем "|" прим: Элемент1|Элемент2|...
func ExpressionOrFunc[S fmt.Stringer](items []S, preprocessor func(S) string) string {
	if preprocessor == nil {
		preprocessor = func(item S) string { return item.String() }
	}

	var builder strings.Builder
	for i, curr := range items {
		if i > 0 {
			builder.WriteRune('|')
		}

		builder.WriteString(preprocessor(curr))
	}

	return builder.String()
}

// ExpressionOrFunc формирует строку элементов с разделителем "," прим: Элемент1,Элемент2,...
func ExpressionListFunc[S fmt.Stringer](items []S, preprocessor func(S) string) string {
	if preprocessor == nil {
		preprocessor = func(item S) string { return item.String() }
	}

	var builder strings.Builder
	for i, curr := range items {
		if i > 0 {
			builder.WriteRune(',')
		}

		builder.WriteString(preprocessor(curr))
	}

	return builder.String()
}

func EscapeRegex[S fmt.Stringer](s S) string {
	replacer := strings.NewReplacer(
		`\`, `\\`,
		`.`, `\.`,
		`*`, `\*`,
		`+`, `\+`,
		`?`, `\?`,
		`^`, `\^`,
		`$`, `\$`,
		`{`, `\{`,
		`}`, `\}`,
		`(`, `\(`,
		`)`, `\)`,
		`|`, `\|`,
		`[`, `\[`,
		`]`, `\]`,
	)
	return replacer.Replace(s.String())
}
