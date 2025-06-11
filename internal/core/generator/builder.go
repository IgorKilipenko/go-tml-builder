package generator

import "github.com/IgorKilipenko/go-tml-builder/internal/core/models"

type GrammarBuilder struct {
    grammar *models.Grammar
}

func NewGrammarBuilder() *GrammarBuilder {
    return &GrammarBuilder{
        grammar: &models.Grammar{
            ScopeName: "source.bsl",
        },
    }
}

func (b *GrammarBuilder) AddRules(rules []*models.Rule) *GrammarBuilder {
    b.grammar.Patterns = append(b.grammar.Patterns, rules...)
    return b
}

func (b *GrammarBuilder) SetRepository(repo models.Repository) *GrammarBuilder {
    b.grammar.Repository = repo
    return b
}

func (b *GrammarBuilder) Build() *models.Grammar {
	b.grammar.Name = "1C:Community (BSL)"
    return b.grammar
}