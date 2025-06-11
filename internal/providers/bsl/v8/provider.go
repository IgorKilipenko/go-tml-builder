package v8

import "github.com/IgorKilipenko/go-tml-builder/internal/core/models"

type BSLv8Provider struct{
	models.GrammarProvider
}

func NewProvider() *BSLv8Provider {
	return &BSLv8Provider{}
}

func (p *BSLv8Provider) GetMainRules() []*models.Rule {
	var rules []*models.Rule

	// Собираем все правила
	rules = append(rules, ControlKeywords()...)
	rules = append(rules, TypeKeywords()...)
	rules = append(rules, CommentRules()...)
	rules = append(rules, FunctionRules()...)

	return rules
}

func (p *BSLv8Provider) GetRepository() models.Repository {
	return models.Repository{
		// 1. Встроенные функции
		models.KeyBuiltinFuncs: {
			Name:  "support.function.builtin.bsl",
			Match: `\b(Строка|Формат|Лев|Прав|Сред|Найти|ВРег|НРег|ТРег|СокрЛ|СокрП|СокрЛП|Символ|КодСимвола|ПустаяСтрока|ЗаполнитьПравее|ЗаполнитьЛевее)\b`,
		},

		// 2. Параметры функций
		models.KeyFunctionParams: {
			Name:  "meta.function.parameters.bsl",
			Begin: `\b(Знач|Val)\s*\(`,
			End:   `\)`,
			BeginCaptures: map[string]models.Capture{
				"0": {Name: "punctuation.definition.parameters.begin.bsl"},
			},
			EndCaptures: map[string]models.Capture{
				"0": {Name: "punctuation.definition.parameters.end.bsl"},
			},
			Patterns: []*models.Rule{
				{
					Include: "#numbers", // Ссылка на другое правило
				},
				{
					Include: "#strings",
				},
			},
		},

		// 3. Строки (включая многострочные)
		"strings": {
			Name:  "string.quoted.double.bsl",
			Begin: `"`,
			End:   `"`,
			Patterns: []*models.Rule{
				{
					Name:  "constant.character.escape.bsl",
					Match: `\\("|n|r|t)`,
				},
			},
		},

		// 4. Числа
		"numbers": {
			Name:  "constant.numeric.bsl",
			Match: `\b\d+(\.\d+)?\b`,
		},

		// 5. Директивы препроцессора (#Если)
		"preprocessor": {
			Name:  "meta.preprocessor.bsl",
			Begin: `^#(Если|КонецЕсли|Иначе|ИначеЕсли)`,
			End:   `$`,
			Captures: map[string]models.Capture{
				"1": {Name: "keyword.control.preprocessor.bsl"},
			},
		},
	}
}
