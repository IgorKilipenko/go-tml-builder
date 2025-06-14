package bsl

import "github.com/IgorKilipenko/go-tml-builder/internal/core/models"

func FunctionRules() []*models.Rule {
	return []*models.Rule{
		{
			Name:  "meta.function.bsl",
			Begin: `\b(Процедура|Функция)\s+([а-яА-ЯёЁ\w]+)\b`,
			End:   `\b(КонецПроцедуры|КонецФункции)\b`,
			Captures: map[string]models.Capture{
				"1": {Name: "keyword.declaration.bsl"},
				"2": {Name: "entity.name.function.bsl"},
			},
			Patterns: []*models.Rule{
				{
					Include: "$self", // Рекурсивно включаем другие правила
				},
			},
		},
	}
}
