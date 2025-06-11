package generator

import (
	"encoding/json"
	"os"

	"github.com/IgorKilipenko/go-tml-builder/internal/core/models"
)

type Repository map[string]*models.Rule

type Grammar struct {
	ScopeName  string         `json:"scopeName"`
	FileTypes  []string       `json:"fileTypes"`
	Name       string         `json:"name"`
	Patterns   []*models.Rule `json:"patterns"`
	Repository Repository     `json:"repository"`
}

func GenerateBSLGrammar() *Grammar {
	return &Grammar{
		ScopeName: "source.bsl",
		FileTypes: []string{"bsl", "os"},
		Name:      "1C:Community (BSL)",
		Patterns:  []*models.Rule{
			// Здесь будут основные правила
		},
		Repository: Repository{
			// Здесь будут дополнительные правила
		},
	}
}

func SaveGrammarToFile(grammar *Grammar, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(grammar)
}
