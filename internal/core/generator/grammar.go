package generator

import (
	"encoding/json"
	"os"

	"github.com/IgorKilipenko/go-tml-builder/internal/core/models"
)

func GenerateBSLGrammar() *models.Grammar {
	return &models.Grammar{
		ScopeName: "source.bsl",
		FileTypes: []string{"bsl", "os"},
		Name:      "1C:Community (BSL)",
		Patterns:  []*models.Rule{
			// Здесь будут основные правила
		},
		Repository: models.Repository{
			// Здесь будут дополнительные правила
		},
	}
}

func SaveGrammarToFile(grammar *models.Grammar, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(grammar)
}
