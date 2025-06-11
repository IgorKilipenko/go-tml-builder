package vscode

import "github.com/IgorKilipenko/go-tml-builder/internal/core/models"

type Snippet struct {
	Prefix string
	Body   string
}

func GenerateSnippets(grammar *models.Grammar) map[string]Snippet {
	return map[string]Snippet{
		"Если": {
			Prefix: "Если",
			Body:   "Если ${1:условие} Тогда\n\t$0\nКонецЕсли;",
		},
	}
}
