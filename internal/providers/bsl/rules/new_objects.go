package rules

import (
	"fmt"

	"github.com/IgorKilipenko/go-tml-builder/internal/core/models"
	bslm "github.com/IgorKilipenko/go-tml-builder/internal/providers/bsl/models"
)

func ObjectDefinitionKey() models.RepositoryKey {
	return bslm.KeyObjectDefinition
}

// CommentLine правила для однострочных комментариев вида: // ...
func ObjectDefinition() *models.Rule {
	newKeyword := bslm.BslNew
	basePattern := fmt.Sprintf(`(?i:(%s))\s+([_$[:alpha:]][_$[:alnum:]]*\b)`, newKeyword)

	patterns := []*models.Rule{
		{
			Match: basePattern,
			Captures: map[string]models.Capture{
				"1": {Name: "keyword.operator.new.bsl"}, // Новый
				"2": {Name: "support.class.bsl"},        // Тип
			},
		},
		{
			Begin: fmt.Sprintf(`%s(?:\s*)(\()\s*`, basePattern),
			BeginCaptures: map[string]models.Capture{
				"1": {Name: "keyword.operator.new.bsl"},
				"2": {Name: "support.class.bsl"},
				"3": {Name: "punctuation.bracket.begin.bsl"},
			},
			End: `\s*(\))`,
			EndCaptures: map[string]models.Capture{
				"1": {Name: "punctuation.bracket.end.bsl"},
			},
			Patterns: []*models.Rule{ // TODO: требуется рефакторинг вложенных правил
				bslm.KeyMiscellaneous.IncludeRef(),
				{Include: "blockEntities"}, // TODO: заменить на IncludeRef после реализации
				bslm.KeyBasic.IncludeRef(),
			},
		},
	}

	return newRule(ObjectDefinitionKey(), patterns)
}
