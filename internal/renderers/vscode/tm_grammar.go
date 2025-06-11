package vscode

import (
	"encoding/json"

	"github.com/IgorKilipenko/go-tml-builder/internal/core/models"
)

type VSCodeRenderer struct{}

func (r *VSCodeRenderer) Render(grammar *models.Grammar) ([]byte, error) {
	return json.MarshalIndent(grammar, "", "  ")
}
