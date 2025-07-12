package main

import (
	"fmt"

	bslm "github.com/IgorKilipenko/go-tml-builder/internal/providers/bsl/models"
)

func main() {
	fmt.Println("=== Примеры использования констант регулярных выражений ===\n")

	// Границы слов
	fmt.Printf("WordBoundary: %s\n", bslm.WordBoundary)
	fmt.Printf("WordBoundaryStart: %s\n", bslm.WordBoundaryStart)
	fmt.Printf("WordBoundaryEnd: %s\n", bslm.WordBoundaryEnd)
	fmt.Printf("WordBoundaryLookBehind: %s\n", bslm.WordBoundaryLookBehind)
	fmt.Printf("WordBoundaryLookAhead: %s\n", bslm.WordBoundaryLookAhead)
	fmt.Println()

	// Идентификаторы
	fmt.Printf("Identifier: %s\n", bslm.Identifier)
	fmt.Printf("IdentifierWithCyrillic: %s\n", bslm.IdentifierWithCyrillic)
	fmt.Printf("IdentifierWithCyrillicCaseInsensitive: %s\n", bslm.IdentifierWithCyrillicCaseInsensitive)
	fmt.Println()

	// Литералы
	fmt.Printf("NumericLiteral: %s\n", bslm.NumericLiteral)
	fmt.Printf("DateLiteral: %s\n", bslm.DateLiteral)
	fmt.Printf("EscapedQuote: %s\n", bslm.EscapedQuote)
	fmt.Println()

	// Шаблоны
	fmt.Printf("WordBoundaryTemplate: %s\n", bslm.WordBoundaryTemplate)
	fmt.Printf("CaseInsensitiveWordBoundaryTemplate: %s\n", bslm.CaseInsensitiveWordBoundaryTemplate)
	fmt.Println()

	// Примеры использования в контексте
	fmt.Println("=== Примеры использования в контексте ===")

	// Пример 1: Граница слова для ключевого слова
	keywordPattern := fmt.Sprintf(`(?i:%s(%s)%s)`,
		bslm.WordBoundaryLookBehind,
		"Если",
		bslm.WordBoundaryLookAhead)
	fmt.Printf("Keyword pattern: %s\n", keywordPattern)

	// Пример 2: Вызов функции
	functionCallPattern := fmt.Sprintf(`(?<!(?i:Новый)\s+)(?:\s*(\.)\s*)?\\b(%s\\b)(?:\\s*)(\()\\s*`,
		bslm.Identifier)
	fmt.Printf("Function call pattern: %s\n", functionCallPattern)

	// Пример 3: Числовой литерал с границами
	numericPattern := fmt.Sprintf(`(?<=%s|^)%s(?=%s|$)`,
		bslm.WordBoundary, bslm.NumericLiteral, bslm.WordBoundary)
	fmt.Printf("Numeric pattern: %s\n", numericPattern)

	// Пример 4: Строковый литерал (без констант для простоты)
	stringPattern := `(")(.*?)(")(?![")`
	fmt.Printf("String pattern: %s\n", stringPattern)

	// Пример 5: Комментарий (без констант для простоты)
	commentPattern := `(^\s*//.*$)`
	fmt.Printf("Comment pattern: %s\n", commentPattern)
}
