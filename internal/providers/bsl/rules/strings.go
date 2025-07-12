package rules

import (
	"github.com/IgorKilipenko/go-tml-builder/internal/core/models"
	bslm "github.com/IgorKilipenko/go-tml-builder/internal/providers/bsl/models"
)

func StringWithSingleValueKey() models.RepositoryKey {
	return bslm.KeyStringWithSingleValue
}

func StringSupportValuesKey() models.RepositoryKey {
	return bslm.KeyStringSupportValues
}

func ExtensionRegionsKey() models.RepositoryKey {
	return bslm.KeyExtensionRegions
}

// StringWithSingleValue правила для строк с одним поддерживаемым значением
func StringWithSingleValue() *models.Rule {
	patterns := []*models.Rule{
		{
			Begin: `(?i:)(?<=\b(?:ПредопределенноеЗначение|Тип)\s*\(\s*)((?:"|$))(?!\s*")`,
			BeginCaptures: map[string]models.Capture{
				"1": {Name: "string.quoted.double.bsl"},
			},
			End: `(?<!"\s*)(")(?!")(?=\))`,
			EndCaptures: map[string]models.Capture{
				"1": {Name: "string.quoted.double.bsl"},
			},
			Patterns: []*models.Rule{
				{
					Match: `(?<=^\s*)"(?!\s*")`,
					Name:  "string.quoted.double.bsl",
				},
				{Include: "_stringInnerSupports"},
			},
		},
		{
			// function callback arguments inline
			Begin: `(?i:)(?<=\b(?:ОписаниеОповещения)\s*\(\s*)(")(?!"\s*")`,
			BeginCaptures: map[string]models.Capture{
				"1": {Name: "string.quoted.double.bsl"},
			},
			End: `(?<!"\s*)(")(?!")`,
			EndCaptures: map[string]models.Capture{
				"1": {Name: "string.quoted.double.bsl"},
			},
			Patterns: []*models.Rule{
				{
					Match: `(?<=\(\s*")\b([_$[:alpha:]][_$[:alnum:]]*)\b(?=")`,
					Name:  "entity.name.function.bsl",
				},
			},
		},
		{
			// function callback arguments in next line
			Begin: `(?i:)(?<=\b(?:ОписаниеОповещения)\b\s*\(\s*)$`,
			End:   `(?<=\S)(")(?!")`,
			EndCaptures: map[string]models.Capture{
				"1": {Name: "string.quoted.double.bsl"},
			},
			Patterns: []*models.Rule{
				{
					Match: `(?<=^\s*)"(?!\s*")`,
					Name:  "string.quoted.double.bsl",
				},
				{
					Match: `(?<=^\s*")\b([_$[:alpha:]][_$[:alnum:]]*)\b(?=")`,
					Name:  "entity.name.function.bsl",
				},
			},
		},
		{
			Begin: `(?<!"\s*)(")(?=[_$[:alpha:]][_$[:alnum:]]*")(?!"\s*")`,
			BeginCaptures: map[string]models.Capture{
				"1": {Name: "string.quoted.double.bsl"},
			},
			End: `(?<!"\s*)(")(?!"\s*")`,
			EndCaptures: map[string]models.Capture{
				"1": {Name: "string.quoted.double.bsl"},
			},
			Patterns: []*models.Rule{
				{Include: "_stringInnerSupports"},
				{Include: "quotedStringBody"},
			},
		},
	}

	return newRule(StringWithSingleValueKey(), patterns)
}

// StringSupportValues правила для строк с поддерживаемыми значениями
func StringSupportValues() *models.Rule {
	patterns := []*models.Rule{
		{
			Begin: `(?i:)(")(Документ|Справочник|Перечисление)\b(?=(?:\s*\.\s*[_$[:alpha:]][_$[:alnum:]]*\s*)+)`,
			BeginCaptures: map[string]models.Capture{
				"1": {Name: "string.quoted.double.bsl"},
				"2": {Name: "support.class.bsl"},
			},
			End: `(?<!"\s*)(")(?!"\s*")`,
			EndCaptures: map[string]models.Capture{
				"1": {Name: "string.quoted.double.bsl"},
			},
			Patterns: []*models.Rule{
				{Include: "blockObjectProperties"},
			},
		},
	}

	return newRule(StringSupportValuesKey(), patterns)
}

// ExtensionRegions правила для регионов расширения
func ExtensionRegions() *models.Rule {
	patterns := []*models.Rule{
		{
			Match: `(?i)(#[Вв]ставка\b)\s*(//.*)?$`,
			Name:  "keyword.control.region.past.bsl",
			Captures: map[string]models.Capture{
				"1": {Name: "keyword.control.region.past.start.bsl"},
				"2": {Name: "comment.line.double-slash.bsl"},
			},
		},
		{
			Match: `(?i)(#[Кк]онец[Вв]ставки\b)\s*(//.*)?$`,
			Name:  "keyword.control.region.past.bsl",
			Captures: map[string]models.Capture{
				"1": {Name: "keyword.control.region.past.end.bsl"},
				"2": {Name: "comment.line.double-slash.bsl"},
			},
		},
		{
			Match: `(?i)(#[Уу]даление\b)\s*(//.*)?$`,
			Name:  "keyword.control.region.del.bsl",
			Captures: map[string]models.Capture{
				"1": {Name: "keyword.control.region.del.start.bsl"},
				"2": {Name: "comment.line.double-slash.bsl"},
			},
		},
		{
			Match: `(?i)(#[Кк]онец[Уу]даления\b)\s*(//.*)?$`,
			Name:  "keyword.control.region.past.del.bsl",
			Captures: map[string]models.Capture{
				"1": {Name: "keyword.control.region.del.end.bsl"},
				"2": {Name: "comment.line.double-slash.bsl"},
			},
		},
	}

	return newRule(ExtensionRegionsKey(), patterns)
}
