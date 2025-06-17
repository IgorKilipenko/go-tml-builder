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
	rules := []*models.Rule{
		bslModels.KeyBasic.IncludeRef(),
		bslModels.KeyMiscellaneous.IncludeRef(),
	}

	return rules
}

func (p *BslProvider) GetRepository() models.Repository {
	repo := make(models.Repository)

	// Стандартные правила
	rule := bslRules.Basic()
	repo[rule.Key] = rule

	rule = bslRules.Miscellaneous()
	repo[rule.Key] = rule

	rule = bslRules.ConstLiterals()
	repo[rule.Key] = rule

	// Комментарии
	rule = bslRules.CommentLine()
	repo[rule.Key] = rule

	rule = bslRules.CommentBlock()
	repo[rule.Key] = rule

	rule = bslRules.DeveloperCommentLine()
	repo[rule.Key] = rule

	// Операторы
	rule = bslRules.SupportOperators()
	repo[rule.Key] = rule

	rule = bslRules.LogicalOperators()
	repo[rule.Key] = rule

	return repo
}
