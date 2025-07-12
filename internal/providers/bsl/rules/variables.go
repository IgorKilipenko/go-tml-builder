package rules

import (
	"fmt"

	"github.com/IgorKilipenko/go-tml-builder/internal/core/models"
	bslm "github.com/IgorKilipenko/go-tml-builder/internal/providers/bsl/models"
)

func VariableDefinitionKey() models.RepositoryKey {
	return bslm.KeyVariableDefinition
}

func StorageModifiersKey() models.RepositoryKey {
	return bslm.KeyStorageModifiers
}

func AnnotationsKey() models.RepositoryKey {
	return bslm.KeyAnnotations
}

// VariableDefinition правила для определения переменных
func VariableDefinition() *models.Rule {
	patterns := []*models.Rule{
		{
			Name:  "keyword.operator.bsl",
			Match: `(;)`,
		},
		{
			Name:  "storage.modifier.bsl",
			Match: fmt.Sprintf(`(?i:%s(%s)%s)`, bslm.WordBoundaryLookBehind, bslm.BslExport, bslm.WordBoundaryLookAhead),
		},
	}

	rule := newRule(VariableDefinitionKey(), patterns)
	rule.Begin = fmt.Sprintf(`(?i:%s(%s)\s+(%s)\s*)`, bslm.WordBoundaryLookBehind, bslm.BslVar, bslm.IdentifierWithCyrillic)
	rule.End = `(;)`
	rule.BeginCaptures = map[string]models.Capture{
		"1": {Name: "storage.type.var.bsl"},
		"2": {Name: "variable.bsl"},
	}
	rule.EndCaptures = map[string]models.Capture{
		"1": {Name: "keyword.operator.bsl"},
	}

	return rule
}

// StorageModifiers правила для модификаторов хранения
func StorageModifiers() *models.Rule {
	patterns := []*models.Rule{}

	rule := newRule(StorageModifiersKey(), patterns)
	rule.Name = "storage.modifier.directive.bsl"
	rule.Match = fmt.Sprintf(`(?i:&(%s((%s(%s)?)?)|%s(%s)?))`,
		bslm.BslOnClient, bslm.BslOnServer, bslm.BslNoContext, bslm.BslOnServer, bslm.BslNoContext)

	return rule
}

// Annotations правила для аннотаций
func Annotations() *models.Rule {
	patterns := []*models.Rule{
		{
			// Аннотации с параметрами
			Begin: `(?i)(&([a-zа-яё0-9_]+))\s*(\()`,
			BeginCaptures: map[string]models.Capture{
				"1": {Name: "storage.type.annotation.bsl"},
				"3": {Name: "punctuation.bracket.begin.bsl"},
			},
			End: `(\))`,
			EndCaptures: map[string]models.Capture{
				"1": {Name: "punctuation.bracket.end.bsl"},
			},
			Patterns: []*models.Rule{
				bslm.KeyBasic.IncludeRef(),
				{
					Name:  "keyword.operator.assignment.bsl",
					Match: `(=)`,
				},
				{
					Name:  "invalid.illegal.bsl",
					Match: fmt.Sprintf(`%s((?<==)(?i)%s)%s`, bslm.WordBoundaryLookBehind, bslm.IdentifierWithCyrillic, bslm.WordBoundaryLookAhead),
				},
				{
					Name:  "invalid.illegal.bsl",
					Match: fmt.Sprintf(`%s((?<==\s)\s*(?i)%s)%s`, bslm.WordBoundaryLookBehind, bslm.IdentifierWithCyrillic, bslm.WordBoundaryLookAhead),
				},
				{
					Name:  "variable.annotation.bsl",
					Match: bslm.IdentifierWithCyrillicCaseInsensitive,
				},
			},
		},
		{
			// Аннотации без параметров
			Name:  "storage.type.annotation.bsl",
			Match: `(?i)(&([a-zа-яё0-9_]+))`,
		},
	}

	return newRule(AnnotationsKey(), patterns)
}
