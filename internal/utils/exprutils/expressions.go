package exprutils

import (
	"fmt"
	"strings"
)

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
