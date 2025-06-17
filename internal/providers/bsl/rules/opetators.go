package rules

import (
	"fmt"

	"github.com/IgorKilipenko/go-tml-builder/internal/core/models"
	"github.com/IgorKilipenko/go-tml-builder/internal/core/regexputil"
	bslm "github.com/IgorKilipenko/go-tml-builder/internal/providers/bsl/models"
)

func SupportOperatorsKey() models.RepositoryKey {
	return bslm.KeySupportOperators
}

func LogicalOperatorsKey() models.RepositoryKey {
	return bslm.KeyLogicalOperators
}

func SupportOperators() *models.Rule {
	patterns := []*models.Rule{
		LogicalOperatorsKey().IncludeRef(), // И|ИЛИ|НЕ
		{
			// >, <, =, >=, <=, <>
			Name: "keyword.operator.comparison.bsl",
			Match: fmt.Sprintf(`%s`,
				regexputil.ExpressionOrFunc(bslm.AllComparisonOperators(), regexputil.EscapeRegex)),
		},
		{
			// +, -, *, /, %
			Name: "keyword.operator.arithmetic.bsl",
			Match: fmt.Sprintf(`%s`,
				regexputil.ExpressionOrFunc(bslm.AllArithmeticOperators(), regexputil.EscapeRegex)),
		},
		{
			// ?, ;
			Name: "keyword.operator.bsl",
			Match: fmt.Sprintf(`%s`,
				regexputil.ExpressionOrFunc([]bslm.BslPunctuation{bslm.BslCondOp, bslm.BslSemicolon}, regexputil.EscapeRegex)),
		},
	}

	return newRule(SupportOperatorsKey(), patterns)
}

func LogicalOperators() *models.Rule {
	allAlphabet := `\wа-яё`
	patterns := []*models.Rule{
		DeveloperCommentLineKey().IncludeRef(),
		{
			// И, ИЛИ, НЕ
			Name: "keyword.operator.logical.bsl",
			Match: fmt.Sprintf(`(?i:(?<=[^%s\.]|^)(%s)(?=[^%s\.]|$))`,
				allAlphabet,
				regexputil.ExpressionOrFunc(bslm.AllLogicalOperators(), nil), // И|ИЛИ|НЕ
				allAlphabet,
			),
		},
	}

	return newRule(LogicalOperatorsKey(), patterns)
}
