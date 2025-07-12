package rules

import (
	"github.com/IgorKilipenko/go-tml-builder/internal/core/models"
	bslm "github.com/IgorKilipenko/go-tml-builder/internal/providers/bsl/models"
)

func FunctionDocumentationKey() models.RepositoryKey {
	return bslm.KeyFunctionDocumentation
}

func DocParametersBlockKey() models.RepositoryKey {
	return bslm.KeyDocParametersBlock
}

func DocReturnsBlockKey() models.RepositoryKey {
	return bslm.KeyDocReturnsBlock
}

func DocParametersBlockArgsKey() models.RepositoryKey {
	return bslm.KeyDocParametersBlockArgs
}

func DocReturnsBlockArgsKey() models.RepositoryKey {
	return bslm.KeyDocReturnsBlockArgs
}

// FunctionDocumentation правила для документации функций
func FunctionDocumentation() *models.Rule {
	patterns := []*models.Rule{
		DocParametersBlockKey().IncludeRef(),
		DocReturnsBlockKey().IncludeRef(),
	}

	rule := newRule(FunctionDocumentationKey(), patterns)
	rule.Name = "meta.documentation.block.bsl"
	rule.Begin = `^\s*(?=(//\s*)([Пп]араметры|[Вв]озвращаемое значение):)`
	rule.End = `(?=^(?!\s*//))`

	return rule
}

// DocParametersBlock правила для блока параметров
func DocParametersBlock() *models.Rule {
	patterns := []*models.Rule{
		DocParametersBlockArgsKey().IncludeRef(),
		bslm.KeyCommentLine.IncludeRef(),
	}

	rule := newRule(DocParametersBlockKey(), patterns)
	rule.Name = "meta.documentation.parameters.bsl"
	rule.Begin = `^\s*(//\s*)([Пп]араметры:)`
	rule.End = `(?=^\s*//\s*[Вв]озвращаемое значение:|^(?!\s*//))`
	rule.BeginCaptures = map[string]models.Capture{
		"1": {Name: "comment.line.double-slash.bsl"},
		"2": {Name: "keyword.control.documentation.bsl"},
	}

	return rule
}

// DocReturnsBlock правила для блока возвращаемых значений
func DocReturnsBlock() *models.Rule {
	patterns := []*models.Rule{
		DocReturnsBlockArgsKey().IncludeRef(),
		DocParametersBlockArgsKey().IncludeRef(),
		bslm.KeyCommentLine.IncludeRef(),
	}

	rule := newRule(DocReturnsBlockKey(), patterns)
	rule.Name = "meta.documentation.returns.bsl"
	rule.Begin = `^\s*(//\s*)([Вв]озвращаемое значение:)`
	rule.End = `(?=^(?!\s*//))`
	rule.BeginCaptures = map[string]models.Capture{
		"1": {Name: "comment.line.double-slash.bsl"},
		"2": {Name: "keyword.control.documentation.bsl"},
	}

	return rule
}

// DocParametersBlockArgs правила для аргументов параметров
func DocParametersBlockArgs() *models.Rule {
	patterns := []*models.Rule{
		{
			Name:  "meta.documentation.parameters.type.list.bsl",
			Match: `(,\s*)([А-Яа-яЁё\w.]+)(?:(\sиз\s+)([А-Яа-яЁё\w.]+))?`,
			Captures: map[string]models.Capture{
				"1": {Name: "punctuation.separator.description.bsl"},
				"2": {Name: "entity.name.type.bsl"},
				"3": {Name: "punctuation.separator.description.bsl"},
				"4": {Name: "entity.name.type.bsl"},
			},
		},
		{
			Name:  "comment.line.description.bsl",
			Begin: `\s*-\s*`,
			End:   `(?=\n|$)`,
			BeginCaptures: map[string]models.Capture{
				"0": {Name: "punctuation.separator.description.bsl"},
			},
		},
	}

	rule := newRule(DocParametersBlockArgsKey(), patterns)
	rule.Begin = `^\s*(//\s*)(\*+\s*)?([А-Яа-яЁё\w]+)\s*-\s*([А-Яа-яЁё\w.]+)(?:(\sиз\s+)([А-Яа-яЁё\w.]+))?`
	rule.End = `(?=\n|$)`
	rule.BeginCaptures = map[string]models.Capture{
		"1": {Name: "comment.line.double-slash.bsl"},
		"2": {Name: "punctuation.separator.description.bsl"},
		"3": {Name: "variable.parameter.bsl"},
		"4": {Name: "entity.name.type.bsl"},
		"5": {Name: "punctuation.separator.description.bsl"},
		"6": {Name: "entity.name.type.bsl"},
	}

	return rule
}

// DocReturnsBlockArgs правила для аргументов возвращаемых значений
func DocReturnsBlockArgs() *models.Rule {
	patterns := []*models.Rule{
		{
			Name:  "meta.documentation.returns.type.list.bsl",
			Match: `(,\s*)([А-Яа-яЁё\w.]+)(?:(\sиз\s+)([А-Яа-яЁё\w.]+))?`,
			Captures: map[string]models.Capture{
				"1": {Name: "punctuation.separator.description.bsl"},
				"2": {Name: "entity.name.type.bsl"},
				"3": {Name: "punctuation.separator.description.bsl"},
				"4": {Name: "entity.name.type.bsl"},
			},
		},
		{
			Name:  "comment.line.description.bsl",
			Begin: `\s*-\s*`,
			End:   `(?=\n|$)`,
			BeginCaptures: map[string]models.Capture{
				"0": {Name: "punctuation.separator.description.bsl"},
			},
		},
	}

	rule := newRule(DocReturnsBlockArgsKey(), patterns)
	rule.Begin = `^\s*(//\s*)(-\s*)?([А-Яа-яЁё\w.]+)(?:(\sиз\s+)([А-Яа-яЁё\w.]+))?`
	rule.End = `(?=\n|$)`
	rule.BeginCaptures = map[string]models.Capture{
		"1": {Name: "comment.line.double-slash.bsl"},
		"2": {Name: "punctuation.separator.description.bsl"},
		"3": {Name: "entity.name.type.bsl"},
		"4": {Name: "punctuation.separator.description.bsl"},
		"5": {Name: "entity.name.type.bsl"},
	}

	return rule
}
