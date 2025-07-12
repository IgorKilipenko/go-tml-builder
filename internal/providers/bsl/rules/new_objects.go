package rules

import (
	"fmt"

	"github.com/IgorKilipenko/go-tml-builder/internal/core/models"
	bslm "github.com/IgorKilipenko/go-tml-builder/internal/providers/bsl/models"
)

func ObjectDefinitionKey() models.RepositoryKey {
	return bslm.KeyObjectDefinition
}

func ObjectDefinition() *models.Rule {
	patterns := []*models.Rule{
		{
			Match: fmt.Sprintf(`(?i:(%s)\\s+)([_$[:alpha:]][_$[:alnum:]]*\\b)`, bslm.BslNew),
			Captures: map[string]models.Capture{
				"1": {Name: "keyword.operator.new.bsl"},
				"2": {Name: "support.class.bsl"},
			},
		},
		{
			Begin: fmt.Sprintf(`(?i:(%s)\\s+)([_$[:alpha:]][_$[:alnum:]]*\\b)(?:\\s*)(%s)\\s*`, bslm.BslNew, bslm.BslParenOpen),
			BeginCaptures: map[string]models.Capture{
				"1": {Name: "keyword.operator.new.bsl"},
				"2": {Name: "support.class.bsl"},
				"3": {Name: "punctuation.bracket.begin.bsl"},
			},
			End: fmt.Sprintf(`\\s*(%s)`, bslm.BslParenClose),
			EndCaptures: map[string]models.Capture{
				"1": {Name: "punctuation.bracket.end.bsl"},
			},
			Patterns: []*models.Rule{
				bslm.KeyMiscellaneous.IncludeRef(),
				{Include: "blockEntities"},
				bslm.KeyBasic.IncludeRef(),
			},
		},
	}

	return newRule(ObjectDefinitionKey(), patterns)
}
