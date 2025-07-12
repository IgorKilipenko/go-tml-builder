package rules

import (
	"fmt"

	"github.com/IgorKilipenko/go-tml-builder/internal/core/models"
	bslm "github.com/IgorKilipenko/go-tml-builder/internal/providers/bsl/models"
)

func Basic() *models.Rule {
	patterns := []*models.Rule{
		CommentBlockKey().IncludeRef(),
		CommentLineKey().IncludeRef(),
		bslm.KeyStringWithSingleValue.IncludeRef(),
		bslm.KeyStringSupportValues.IncludeRef(),
		bslm.KeyQuotedString.IncludeRef(),
		bslm.KeyConstantsLiterals.IncludeRef(),
		{ // numerics literals
			Name:  "constant.numeric.bsl",
			Match: `(?<=[^\wа-яё\.]|^)(\d+\.?\d*)(?=[^\wа-яё\.]|$)`,
		},
		{ // date literals, like: '20250101'
			Name:  "constant.other.date.bsl",
			Match: `'((\d{4}[^\d']*\d{2}[^\d']*\d{2})([^\d']*\d{2}[^\d']*\d{2}([^\d']*\d{2})?)?)'`,
		},
		{ // comma like: ,
			Name:  "keyword.operator.bsl",
			Match: fmt.Sprintf(`(%s)`, bslm.BslSemicolon),
		},
		{ // like: (
			Name:  "punctuation.bracket.begin.bsl",
			Match: fmt.Sprintf(`(%s)`, bslm.BslParenOpen),
		},
		{ // like: )
			Name:  "punctuation.bracket.end.bsl",
			Match: fmt.Sprintf(`(%s)`, bslm.BslParenClose),
		},
		bslm.KeyExtensionRegions.IncludeRef(),
	}

	return newRule(bslm.KeyBasic, patterns)
}

func Miscellaneous() *models.Rule {
	patterns := []*models.Rule{
		SupportOperatorsKey().IncludeRef(),
		ObjectDefinitionKey().IncludeRef(),
		CallSupportFunctionsKey().IncludeRef(),
		SupportEnumsKey().IncludeRef(),
		SupportClassesKey().IncludeRef(),
		SupportValueTypesKey().IncludeRef(),
		SupportLanguageConstantKey().IncludeRef(),
	}

	return newRule(bslm.KeyMiscellaneous, patterns)
}

func newRule(key models.RepositoryKey, patterns []*models.Rule) *models.Rule {
	rule := &models.Rule{
		Key: key,
	}

	rule.Patterns = append(rule.Patterns, patterns...)

	return rule
}
