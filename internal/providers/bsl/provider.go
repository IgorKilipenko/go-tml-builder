package bsl

import (
	"github.com/IgorKilipenko/go-tml-builder/internal/core/models"
	bslModels "github.com/IgorKilipenko/go-tml-builder/internal/providers/bsl/models"
	bslRules "github.com/IgorKilipenko/go-tml-builder/internal/providers/bsl/rules"
)

type BslProvider struct{}

func NewProvider() *BslProvider {
	return &BslProvider{}
}

func (p *BslProvider) GetMainRules() []*models.Rule {
	// var rules []*models.Rule

	// // Собираем все правила
	// rules = append(rules, ControlKeywords()...)
	// rules = append(rules, TypeKeywords()...)
	// rules = append(rules, CommentRules()...)
	// rules = append(rules, FunctionRules()...)

	rules := []*models.Rule{
		bslModels.KeyBasic.IncludeRef(),
		bslModels.KeyMiscellaneous.IncludeRef(),
	}

	return rules
}

func (p *BslProvider) GetRepository() models.Repository {
	repo := make(models.Repository)

	// Стандартные правила
	rule := bslRules.BasicWithDefPatterns()
	repo[rule.Key] = rule

	rule = bslRules.MiscellaneousWithDefPatterns()
	repo[rule.Key] = rule

	return repo
}
