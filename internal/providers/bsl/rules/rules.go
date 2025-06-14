package rules

import (
	"fmt"

	"github.com/IgorKilipenko/go-tml-builder/internal/core/models"
	bslm "github.com/IgorKilipenko/go-tml-builder/internal/providers/bsl/models"
	"github.com/IgorKilipenko/go-tml-builder/internal/utils/exprutils"
)

func Basic(patterns []*models.Rule) *models.Rule {
	return newRule(bslm.KeyBasic, patterns)
}

func BasicWithDefPatterns() *models.Rule {
	patterns := []*models.Rule{
		{ // comments (single line)
			Name:  "comment.line.double-slash.bsl",
			Begin: "//",
			End:   "$",
		},
		bslm.KeyStringWithSingleValue.IncludeRef(),
		bslm.KeyStringSupportValues.IncludeRef(),
		bslm.KeyQuotedString.IncludeRef(),
		{ // keyword literals, like: Неопределено
			Name: "constant.language.bsl",
			Match: fmt.Sprintf("(?i:(?<=[^\\wа-яё\\.]|^)(%s)(?=[^\\wа-яё\\.]|$))",
				exprutils.ExpressionOrFunc(bslm.AllConstants(), nil)),
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
		bslm.KeyExtensionRegions.IncludeRef(),
	}

	return Basic(patterns)
}

func Miscellaneous(patterns []*models.Rule) *models.Rule {
	return newRule(bslm.KeyMiscellaneous, patterns)
}

func MiscellaneousWithDefPatterns() *models.Rule {
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

	return Miscellaneous(patterns)
}

func newRule(key models.RepositoryKey, patterns []*models.Rule) *models.Rule {
	rule := &models.Rule{
		Key: key,
	}

	rule.Patterns = append(rule.Patterns, patterns...)

	return rule
}
