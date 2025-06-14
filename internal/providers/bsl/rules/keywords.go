package rules

import "github.com/IgorKilipenko/go-tml-builder/internal/core/models"

// Ключевые слова управления потоком
func ControlKeywords() []*models.Rule {
    return []*models.Rule{
        {
            Name:  "keyword.control.bsl",
            Match: `\b(Если|Тогда|ИначеЕсли|Иначе|КонецЕсли|Для|Каждого|Из|Цикл|КонецЦикла|Пока|Продолжить|Прервать|Возврат|Попытка|Исключение|КонецПопытки|ВызватьИсключение)\b`,
        },
    }
}

// Типы данных
func TypeKeywords() []*models.Rule {
    return []*models.Rule{
        {
            Name:  "storage.type.bsl",
            Match: `\b(Перем|Var|Знач|Val|Экспорт|Export)\b`,
        },
    }
}