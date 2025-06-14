package main

import (
	"log"

	"github.com/IgorKilipenko/go-tml-builder/internal/core/generator"
	providers "github.com/IgorKilipenko/go-tml-builder/internal/providers/bsl"
)

func main() {
	provider := providers.NewProvider()
	builder := generator.NewGrammarBuilder()
	builder.AddRules(provider.GetMainRules()).SetRepository(provider.GetRepository())
	grammar := builder.Build()

	err := generator.SaveGrammarToFile(grammar, "bsl.tmLanguage.json")
	if err != nil {
		log.Fatalf("Error saving grammar: %v", err)
	}
}
