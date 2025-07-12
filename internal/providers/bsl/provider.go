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
		bslModels.KeyFunctionDocumentation.IncludeRef(),
		bslModels.KeyFunctionDefinition.IncludeRef(),
		bslModels.KeyFunctionEnd.IncludeRef(),
		bslModels.KeyVariableDefinition.IncludeRef(),
		bslModels.KeyConditional.IncludeRef(),
		bslModels.KeyVariableAssignment.IncludeRef(),
		bslModels.KeyControlKeywords.IncludeRef(),
		bslModels.KeyAnnotations.IncludeRef(),
		bslModels.KeyStorageModifiers.IncludeRef(),
		bslModels.KeyMainRegion.IncludeRef(),
		bslModels.KeyBlockEntities.IncludeRef(),
	}

	return rules
}

func (p *BslProvider) GetRepository() models.Repository {
	repo := make(models.Repository)

	// Базовые правила
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

	// Документация функций
	rule = bslRules.FunctionDocumentation()
	repo[rule.Key] = rule

	rule = bslRules.DocParametersBlock()
	repo[rule.Key] = rule

	rule = bslRules.DocReturnsBlock()
	repo[rule.Key] = rule

	rule = bslRules.DocParametersBlockArgs()
	repo[rule.Key] = rule

	rule = bslRules.DocReturnsBlockArgs()
	repo[rule.Key] = rule

	// Области
	rule = bslRules.MainRegion()
	repo[rule.Key] = rule

	rule = bslRules.MainRegionStart()
	repo[rule.Key] = rule

	rule = bslRules.MainRegionEnd()
	repo[rule.Key] = rule

	// Операторы
	rule = bslRules.SupportOperators()
	repo[rule.Key] = rule

	rule = bslRules.LogicalOperators()
	repo[rule.Key] = rule

	// Конструкторы: Новый ()
	rule = bslRules.ObjectDefinition()
	repo[rule.Key] = rule

	// Функции
	rule = bslRules.CallSupportFunctions()
	repo[rule.Key] = rule

	// Управление потоком
	rule = bslRules.ControlKeywordsRule()
	repo[rule.Key] = rule

	rule = bslRules.Conditional()
	repo[rule.Key] = rule

	rule = bslRules.VariableAssignment()
	repo[rule.Key] = rule

	rule = bslRules.FunctionDefinition()
	repo[rule.Key] = rule

	rule = bslRules.FunctionEnd()
	repo[rule.Key] = rule

	// Переменные и аннотации
	rule = bslRules.VariableDefinition()
	repo[rule.Key] = rule

	rule = bslRules.Annotations()
	repo[rule.Key] = rule

	rule = bslRules.StorageModifiers()
	repo[rule.Key] = rule

	// Блоки и строки
	rule = bslRules.BlockEntities()
	repo[rule.Key] = rule

	rule = bslRules.BlockVariables()
	repo[rule.Key] = rule

	rule = bslRules.BlockFunctions()
	repo[rule.Key] = rule

	rule = bslRules.BlockObjectProperties()
	repo[rule.Key] = rule

	rule = bslRules.ArrayLiteral()
	repo[rule.Key] = rule

	rule = bslRules.BlockAwait()
	repo[rule.Key] = rule

	rule = bslRules.QuotedString()
	repo[rule.Key] = rule

	rule = bslRules.QuotedStringBody()
	repo[rule.Key] = rule

	rule = bslRules.StringPatternParameter()
	repo[rule.Key] = rule

	// Строковые правила
	rule = bslRules.StringWithSingleValue()
	repo[rule.Key] = rule

	rule = bslRules.StringSupportValues()
	repo[rule.Key] = rule

	rule = bslRules.ExtensionRegions()
	repo[rule.Key] = rule

	// Поддерживаемые типы
	rule = bslRules.SupportEnums()
	repo[rule.Key] = rule

	rule = bslRules.SupportClasses()
	repo[rule.Key] = rule

	rule = bslRules.SupportValueTypes()
	repo[rule.Key] = rule

	rule = bslRules.SupportLanguageConstant()
	repo[rule.Key] = rule

	rule = bslRules.SupportRegisterTable()
	repo[rule.Key] = rule

	// Запросы
	rule = bslRules.Query()
	repo[rule.Key] = rule

	return repo
}
