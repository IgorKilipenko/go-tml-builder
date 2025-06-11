package main

import (
	"log"

	"github.com/IgorKilipenko/go-tml-builder/internal/core/generator"
	"github.com/IgorKilipenko/go-tml-builder/internal/core/models"
)

func main() {
	grammar := generator.GenerateBSLGrammar()

	// Добавляем правила
	grammar.Patterns = append(grammar.Patterns, &models.Rule{
		Match: `\b(Процедура|Функция)\b`,
		Name:  "keyword.control.bsl",
	})

	// ... другие правила

	err := generator.SaveGrammarToFile(grammar, "bsl.tmLanguage.json")
	if err != nil {
		log.Fatalf("Error saving grammar: %v", err)
	}
}
