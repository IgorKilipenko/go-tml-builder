package generator

import (
	"github.com/IgorKilipenko/go-tml-builder/internal/models"
)

type Repository map[string]*Rule

type Grammar struct {
	ScopeName  string     `json:"scopeName"`
	FileTypes  []string   `json:"fileTypes"`
	Name       string     `json:"name"`
	Patterns   []*Rule    `json:"patterns"`
	Repository Repository `json:"repository"`
}
