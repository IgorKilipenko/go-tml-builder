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
			Match: fmt.Sprintf(`(?<=%s|^)%s(?=%s|$)`, bslm.WordBoundary, bslm.NumericLiteral, bslm.WordBoundary),
		},
		{ // date literals, like: '20250101'
			Name:  "constant.other.date.bsl",
			Match: bslm.DateLiteral,
		},
		{ // comma like: ,
			Name:  "keyword.operator.bsl",
			Match: `(,)`,
		},
		{ // like: (
			Name:  "punctuation.bracket.begin.bsl",
			Match: `(\()`,
		},
		{ // like: )
			Name:  "punctuation.bracket.end.bsl",
			Match: `(\))`,
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
