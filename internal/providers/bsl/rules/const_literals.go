package rules

import (
	"fmt"

	"github.com/IgorKilipenko/go-tml-builder/internal/core/models"
	"github.com/IgorKilipenko/go-tml-builder/internal/core/regexputil"
	bslm "github.com/IgorKilipenko/go-tml-builder/internal/providers/bsl/models"
)

func ConstLiteralsKey() models.RepositoryKey {
	return bslm.KeyConstantsLiterals
}

func ConstLiterals() *models.Rule {
	patterns := []*models.Rule{
		{
			Name: "constant.language.bsl",
			Match: fmt.Sprintf(`(?i:%s(%s)%s)`,
				bslm.WordBoundaryLookBehind,
				regexputil.ExpressionOrFunc(bslm.AllConstLiterals(), nil),
				bslm.WordBoundaryLookAhead),
		},
	}

	return newRule(ConstLiteralsKey(), patterns)
}
