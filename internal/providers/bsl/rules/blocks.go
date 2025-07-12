package rules

import (
	"fmt"

	"github.com/IgorKilipenko/go-tml-builder/internal/core/models"
	bslm "github.com/IgorKilipenko/go-tml-builder/internal/providers/bsl/models"
)

func BlockEntitiesKey() models.RepositoryKey {
	return bslm.KeyBlockEntities
}

func BlockVariablesKey() models.RepositoryKey {
	return bslm.KeyBlockVariables
}

func BlockFunctionsKey() models.RepositoryKey {
	return bslm.KeyBlockFunctions
}

func BlockObjectPropertiesKey() models.RepositoryKey {
	return bslm.KeyBlockObjectProperties
}

func ArrayLiteralKey() models.RepositoryKey {
	return bslm.KeyArrayLiteral
}

func BlockAwaitKey() models.RepositoryKey {
	return bslm.KeyBlockAwait
}

func QuotedStringKey() models.RepositoryKey {
	return bslm.KeyQuotedString
}

func QuotedStringBodyKey() models.RepositoryKey {
	return bslm.KeyQuotedStringBody
}

func StringPatternParameterKey() models.RepositoryKey {
	return bslm.KeyStringPatternParameter
}

// BlockEntities правила для блоков сущностей
func BlockEntities() *models.Rule {
	patterns := []*models.Rule{
		BlockAwaitKey().IncludeRef(),
		ArrayLiteralKey().IncludeRef(),
		BlockFunctionsKey().IncludeRef(),
		BlockObjectPropertiesKey().IncludeRef(),
		BlockVariablesKey().IncludeRef(),
	}

	return newRule(BlockEntitiesKey(), patterns)
}

// BlockVariables правила для переменных в блоках
func BlockVariables() *models.Rule {
	patterns := []*models.Rule{
		{
			Match: fmt.Sprintf(`\b(%s)\b(?=[\s*\[]?)(?![\\(])`, bslm.Identifier),
			Captures: map[string]models.Capture{
				"1": {Name: "variable.other.readwrite.bsl, entity.name.variable.bsl"},
			},
		},
	}

	return newRule(BlockVariablesKey(), patterns)
}

// BlockFunctions правила для вызова функций в блоках
func BlockFunctions() *models.Rule {
	patterns := []*models.Rule{
		{
			Begin: fmt.Sprintf(`(?<!(?i:%s)\s+)(?:\s*(\.)\s*)?\b(%s\b)(?:\s*)(\()\s*`, bslm.BslNew, bslm.Identifier),
			BeginCaptures: map[string]models.Capture{
				"1": {Name: "punctuation.accessor.bsl"},
				"2": {Name: "entity.name.function.bsl"},
				"3": {Name: "punctuation.bracket.begin.bsl"},
			},
			End: `\s*(\))`,
			EndCaptures: map[string]models.Capture{
				"1": {Name: "punctuation.bracket.end.bsl"},
			},
			Patterns: []*models.Rule{
				bslm.KeyBasic.IncludeRef(),
				bslm.KeyMiscellaneous.IncludeRef(),
				{Include: "blockEntities"},
			},
		},
	}

	return newRule(BlockFunctionsKey(), patterns)
}

// BlockObjectProperties правила для свойств объектов
func BlockObjectProperties() *models.Rule {
	patterns := []*models.Rule{
		{
			Match: `(?<=(?i:Метаданные)\s*)(\.)(?:\s*)(РежимСовместимости|СвойстваОбъектов)(?:\s*)(?!\s*\()`,
			Captures: map[string]models.Capture{
				"1": {Name: "punctuation.accessor.bsl"},
				"2": {Name: "support.constant.bsl"},
			},
		},
		{
			Match: fmt.Sprintf(`(?<=(?:\b%s\s*)|(?:[\s\]]*))(\.)(?:\s*)(%s\b)(?:\s*)(?=[\[]?)(?!\s*\()`, bslm.Identifier, bslm.Identifier),
			Captures: map[string]models.Capture{
				"1": {Name: "punctuation.accessor.bsl"},
				"2": {Name: "variable.other.property.bsl"},
			},
		},
	}

	return newRule(BlockObjectPropertiesKey(), patterns)
}

// ArrayLiteral правила для литералов массивов
func ArrayLiteral() *models.Rule {
	patterns := []*models.Rule{
		bslm.KeyBasic.IncludeRef(),
		{Include: "blockEntities"},
	}

	rule := newRule(ArrayLiteralKey(), patterns)
	rule.Name = "meta.array.literal.bsl"
	rule.Begin = `\s*(\[)`
	rule.End = `\]`
	rule.BeginCaptures = map[string]models.Capture{
		"1": {Name: "meta.brace.square.bsl"},
	}
	rule.EndCaptures = map[string]models.Capture{
		"0": {Name: "meta.brace.square.bsl"},
	}

	return rule
}

// BlockAwait правила для await
func BlockAwait() *models.Rule {
	patterns := []*models.Rule{
		{
			Match: fmt.Sprintf(`(?i)\b(%s)\b`, bslm.BslAwait),
			Name:  "keyword.control.flow.bsl",
		},
	}

	return newRule(BlockAwaitKey(), patterns)
}

// QuotedString правила для строк в кавычках
func QuotedString() *models.Rule {
	patterns := []*models.Rule{
		QuotedStringBodyKey().IncludeRef(),
	}

	rule := newRule(QuotedStringKey(), patterns)
	rule.Begin = `(")`
	rule.End = `(")(?![")`
	rule.BeginCaptures = map[string]models.Capture{
		"1": {Name: "string.quoted.double.bsl"},
	}
	rule.EndCaptures = map[string]models.Capture{
		"1": {Name: "string.quoted.double.bsl"},
	}

	return rule
}

// QuotedStringBody правила для содержимого строк
func QuotedStringBody() *models.Rule {
	patterns := []*models.Rule{
		{
			Name:  "constant.character.escape.bsl",
			Match: bslm.EscapedQuote,
		},
		{
			Name:  "comment.line.double-slash.bsl",
			Match: `(^\s*//.*$)`,
		},
		StringPatternParameterKey().IncludeRef(),
		{
			Match: `[^"]`,
			Name:  "string.quoted.double.bsl",
		},
	}

	return newRule(QuotedStringBodyKey(), patterns)
}

// StringPatternParameter правила для параметров строковых шаблонов
func StringPatternParameter() *models.Rule {
	patterns := []*models.Rule{
		{
			Match: `%[0-9]+`,
			Name:  "variable.parameter.bsl",
		},
	}

	return newRule(StringPatternParameterKey(), patterns)
}
