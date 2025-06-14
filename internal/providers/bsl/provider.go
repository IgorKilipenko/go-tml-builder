package bsl

import (
	"github.com/IgorKilipenko/go-tml-builder/internal/core/models"
	rs "github.com/IgorKilipenko/go-tml-builder/internal/providers/bsl/rules"
)

type BSLProvider struct{}

func NewProvider() *BSLProvider {
	return &BSLProvider{}
}

func (p *BSLProvider) GetMainRules() []*models.Rule {
	// var rules []*models.Rule

	// // Собираем все правила
	// rules = append(rules, ControlKeywords()...)
	// rules = append(rules, TypeKeywords()...)
	// rules = append(rules, CommentRules()...)
	// rules = append(rules, FunctionRules()...)

	rules := []*models.Rule{
		rs.KeyBasic.IncludeRef(),
		rs.KeyMiscellaneous.IncludeRef(),
	}

	return rules
}

func (p *BSLProvider) GetRepository() models.Repository {
	repo := make(models.Repository)

	// Стандартные правила
	rule := rs.NewBasicWithDefPatterns()
	repo[rule.Key] = rule

	rule = rs.NewMiscellaneousWithDefPatterns()
	repo[rule.Key] = rule

	return repo
}
