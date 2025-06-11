package main

import (
	"log"
	"github.com/IgorKilipenko/go-tml-builder"
)

func main() {
	grammar := generator.GenerateBSLGrammar()

	// Добавляем правила (можно вынести в отдельный пакет)
	grammar.Patterns = append(grammar.Patterns, &generator.Rule{
		Match: `\b(Процедура|Функция)\b`,
		Name:  "keyword.control.bsl",
	})

	// ... другие правила

	err := generator.SaveGrammarToFile(grammar, "bsl.tmLanguage.json")
	if err != nil {
		log.Fatalf("Error saving grammar: %v", err)
	}
}
