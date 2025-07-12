package rules

import (
	"fmt"

	"github.com/IgorKilipenko/go-tml-builder/internal/core/models"
	"github.com/IgorKilipenko/go-tml-builder/internal/core/regexputil"
	bslm "github.com/IgorKilipenko/go-tml-builder/internal/providers/bsl/models"
)

func ControlKeywordsKey() models.RepositoryKey {
	return bslm.KeyControlKeywords
}

func ConditionalKey() models.RepositoryKey {
	return bslm.KeyConditional
}

func VariableAssignmentKey() models.RepositoryKey {
	return bslm.KeyVariableAssignment
}

func FunctionDefinitionKey() models.RepositoryKey {
	return bslm.KeyFunctionDefinition
}

func FunctionEndKey() models.RepositoryKey {
	return bslm.KeyFunctionEnd
}

// ControlKeywordsRule правила для ключевых слов управления потоком
func ControlKeywordsRule() *models.Rule {
	patterns := []*models.Rule{
		{
			Name:  "keyword.control.import.bsl",
			Match: fmt.Sprintf(`(?i)#(%s)(?=%s|$)`, bslm.BslUse, bslm.WordBoundary),
		},
		{
			Name:  "keyword.control.native.bsl",
			Match: `(?i)#native`,
		},
		{
			Name: "keyword.control.bsl",
			Match: fmt.Sprintf(`(?i:%s(%s)%s)`,
				bslm.WordBoundaryLookBehind,
				regexputil.ExpressionOrFunc([]bslm.BslKeywords{bslm.BslReturn, bslm.BslContinue, bslm.BslBrake}, nil),
				bslm.WordBoundaryLookAhead),
		},
		{
			Name: "keyword.control.conditional.bsl",
			Match: fmt.Sprintf(`(?i:%s(%s)%s)`,
				bslm.WordBoundaryLookBehind,
				regexputil.ExpressionOrFunc([]bslm.BslKeywords{bslm.BslIf, bslm.BslElse, bslm.BslElseIf, bslm.BslThen, bslm.BslEndIf}, nil),
				bslm.WordBoundaryLookAhead),
		},
		{
			Name: "keyword.control.exception.bsl",
			Match: fmt.Sprintf(`(?i:%s(%s)%s)`,
				bslm.WordBoundaryLookBehind,
				regexputil.ExpressionOrFunc([]bslm.BslKeywords{bslm.BslTry, bslm.BslCatch, bslm.BslEndTry, bslm.BslException}, nil),
				bslm.WordBoundaryLookAhead),
		},
		{
			Name: "keyword.control.repeat.bsl",
			Match: fmt.Sprintf(`(?i:%s(%s)%s)`,
				bslm.WordBoundaryLookBehind,
				regexputil.ExpressionOrFunc([]bslm.BslKeywords{bslm.BslWhile, bslm.BslFor, bslm.BslEach, bslm.BslIn, bslm.BslLoop, bslm.BslEndLoop}, nil),
				bslm.WordBoundaryLookAhead),
		},
	}

	return newRule(ControlKeywordsKey(), patterns)
}

// Conditional правила для условных конструкций
func Conditional() *models.Rule {
	patterns := []*models.Rule{
		bslm.KeyBasic.IncludeRef(),
		bslm.KeyMiscellaneous.IncludeRef(),
		{Include: "blockEntities"},
	}

	rule := newRule(ConditionalKey(), patterns)
	rule.Name = "meta.conditional.bsl"
	rule.Begin = fmt.Sprintf(`(?i:(?<=;|^)\s*(%s))`, bslm.BslIf)
	rule.End = fmt.Sprintf(`(?i:(%s))`, bslm.BslThen)
	rule.BeginCaptures = map[string]models.Capture{
		"1": {Name: "keyword.control.conditional.bsl"},
	}
	rule.EndCaptures = map[string]models.Capture{
		"1": {Name: "keyword.control.conditional.bsl"},
	}

	return rule
}

// VariableAssignment правила для присваивания переменных
func VariableAssignment() *models.Rule {
	patterns := []*models.Rule{
		bslm.KeyBasic.IncludeRef(),
		bslm.KeyMiscellaneous.IncludeRef(),
		{Include: "blockEntities"},
	}

	rule := newRule(VariableAssignmentKey(), patterns)
	rule.Name = "meta.var-single-variable.bsl"
	rule.Begin = `(?i:(?<=;|^)\s*([\wа-яё]+))\s*(=)`
	rule.End = bslm.StatementEnd
	rule.BeginCaptures = map[string]models.Capture{
		"1": {Name: "variable.assignment.bsl"},
		"2": {Name: "keyword.operator.assignment.bsl"},
	}

	return rule
}

// FunctionDefinition правила для определения функций и процедур
func FunctionDefinition() *models.Rule {
	patterns := []*models.Rule{
		bslm.KeyAnnotations.IncludeRef(),
		bslm.KeyBasic.IncludeRef(),
		{
			Name:  "storage.modifier.bsl",
			Match: fmt.Sprintf(`(?i:%s(%s)%s)`, bslm.WordBoundaryLookBehind, bslm.BslVal, bslm.WordBoundaryLookAhead),
		},
		{
			Name:  "invalid.illegal.bsl",
			Match: fmt.Sprintf(`%s((?<==)(?i)%s)%s`, bslm.WordBoundaryLookBehind, bslm.IdentifierWithCyrillic, bslm.WordBoundaryLookAhead),
		},
		{
			Name:  "invalid.illegal.bsl",
			Match: fmt.Sprintf(`%s((?<==\s)\s*(?i)%s)%s`, bslm.WordBoundaryLookBehind, bslm.IdentifierWithCyrillic, bslm.WordBoundaryLookAhead),
		},
		bslm.KeyMiscellaneous.IncludeRef(),
		{Include: "blockEntities"},
		{
			Name:  "variable.parameter.bsl",
			Match: bslm.IdentifierWithCyrillicCaseInsensitive,
		},
	}

	rule := newRule(FunctionDefinitionKey(), patterns)
	rule.Begin = fmt.Sprintf(`(?i:(%s)\s+)?(%s|%s)\s+(%s)\s*(\()`,
		bslm.BslAsync, bslm.BslProcedure, bslm.BslFunction, bslm.Identifier)
	rule.End = fmt.Sprintf(`(?i:(\))\s*((%s)(?=%s|$))?)`, bslm.BslExport, bslm.WordBoundary)
	rule.BeginCaptures = map[string]models.Capture{
		"1": {Name: "storage.modifier.async.bsl"},
		"2": {Name: "storage.type.bsl"},
		"3": {Name: "entity.name.function.bsl"},
		"4": {Name: "punctuation.bracket.begin.bsl"},
	}
	rule.EndCaptures = map[string]models.Capture{
		"1": {Name: "punctuation.bracket.end.bsl"},
		"2": {Name: "storage.modifier.bsl"},
	}

	return rule
}

// FunctionEnd правила для окончания функций и процедур
func FunctionEnd() *models.Rule {
	patterns := []*models.Rule{
		bslm.KeyCommentLine.IncludeRef(),
	}

	rule := newRule(FunctionEndKey(), patterns)
	rule.Begin = fmt.Sprintf(`(?i)(%s|%s)`, bslm.BslEndProc, bslm.BslEndFunc)
	rule.End = `$`
	rule.BeginCaptures = map[string]models.Capture{
		"1": {Name: "storage.type.bsl"},
	}

	return rule
}
