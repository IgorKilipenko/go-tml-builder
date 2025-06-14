package rules

import (
	"github.com/IgorKilipenko/go-tml-builder/internal/core/models"
)

func NewBasic(patterns []*models.Rule) *models.Rule {
	return newRule(KeyBasic, patterns)
}

func NewBasicWithDefPatterns() *models.Rule {
	patterns := []*models.Rule{
		{ // comments (single line)
			Name:  "comment.line.double-slash.bsl",
			Begin: "//",
			End:   "$",
		},
		KeyStringWithSingleValue.IncludeRef(),
		KeyStringSupportValues.IncludeRef(),
		KeyQuotedString.IncludeRef(),
		{ // keyword literals, like: Неопределено
			Name:  "constant.language.bsl",
			Match: "(?i:(?<=[^\\wа-яё\\.]|^)(Неопределено|Истина|Ложь)(?=[^\\wа-яё\\.]|$))",
		},
		{ // numerics literals
			Name:  "constant.numeric.bsl",
			Match: "(?<=[^\\wа-яё\\.]|^)(\\d+\\.?\\d*)(?=[^\\wа-яё\\.]|$)",
		},
		{ // date literals, like: '20250101'
			Name:  "constant.other.date.bsl",
			Match: "\\'((\\d{4}[^\\d\\']*\\d{2}[^\\d\\']*\\d{2})([^\\d\\']*\\d{2}[^\\d\\']*\\d{2}([^\\d\\']*\\d{2})?)?)\\'",
		},
		{ // like: ,
			Name:  "keyword.operator.bsl",
			Match: "(,)",
		},
		{ // like: (
			Name:  "punctuation.bracket.begin.bsl",
			Match: "(\\()",
		},
		{ // like: (
			Name:  "punctuation.bracket.end.bsl",
			Match: "(\\))",
		},
		KeyExtensionRegions.IncludeRef(),
	}

	return NewBasic(patterns)
}

func NewMiscellaneous(patterns []*models.Rule) *models.Rule {
	return newRule(KeyMiscellaneous, patterns)
}

func NewMiscellaneousWithDefPatterns() *models.Rule {
	patterns := []*models.Rule{
		{
			Include: "keywordOperators",
		},
		{
			Include: "objectDefinition",
		},
		{
			Include: "supportFunctions",
		},
		{
			Include: "supportEnums",
		},
		{
			Include: "supportClasses",
		},
		{
			Include: "supportValueTypes",
		},
		{
			Include: "supportLanguageConstant",
		},
	}

	return NewMiscellaneous(patterns)
}

func newRule(key models.RepositoryKey, patterns []*models.Rule) *models.Rule {
	rule := &models.Rule{
		Key: key,
	}

	rule.Patterns = append(rule.Patterns, patterns...)

	return rule
}
