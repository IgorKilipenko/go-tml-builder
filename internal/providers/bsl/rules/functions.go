package rules

import (
	"fmt"

	"github.com/IgorKilipenko/go-tml-builder/internal/core/models"
	"github.com/IgorKilipenko/go-tml-builder/internal/core/regexputil"
	bslm "github.com/IgorKilipenko/go-tml-builder/internal/providers/bsl/models"
)

func CallSupportFunctionsKey() models.RepositoryKey {
	return bslm.KeyCallSupportFunctions
}

// CallSupportFunctions правила для вызова глобальных предопределенных функций
func CallSupportFunctions() *models.Rule {
	patterns := []*models.Rule{
		{
			Name: "support.function.bsl",
			Match: fmt.Sprintf("(?x)(?i:(?<=[^\\wа-яё\\.]|^)(%s)\\s*(?=\\())",
				regexputil.ExpressionOrFunc(bslm.AllConstLiterals(), nil)),
		},
	}

	return newRule(CallSupportFunctionsKey(), patterns)
}
