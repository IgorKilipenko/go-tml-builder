# Динамическое формирование регулярных выражений с использованием ExpressionOrFunc

## Обзор

В проекте `go-tml-builder` мы используем подход динамического формирования регулярных выражений с помощью функции `ExpressionOrFunc`. Это позволяет создавать более поддерживаемые и читаемые правила грамматики TextMate.

## Преимущества подхода

### 1. Централизованное управление ключевыми словами
Все ключевые слова языка BSL определены в одном месте (`internal/providers/bsl/models/keywords.go`), что упрощает их поддержку и обновление.

### 2. Автоматическое формирование регулярных выражений
Вместо ручного написания регулярных выражений вида:
```go
Match: `\b(Если|Тогда|ИначеЕсли|Иначе|КонецЕсли|Для|Каждого|Из|Цикл|КонецЦикла|Пока|Продолжить|Прервать|Возврат|Попытка|Исключение|КонецПопытки|ВызватьИсключение)\b`
```

Мы используем:
```go
Match: fmt.Sprintf(`(?i:(?<=[^\wа-яё\.]|^)(%s)(?=[^\wа-яё\.]|$))`,
    regexputil.ExpressionOrFunc(bslm.AllControlKeywords(), nil))
```

### 3. Гибкость с препроцессорами
Функция `ExpressionOrFunc` поддерживает препроцессоры для модификации элементов:

```go
// Без препроцессора
regexputil.ExpressionOrFunc(bslm.AllLogicalOperators(), nil)
// Результат: И|ИЛИ|НЕ

// С препроцессором экранирования
regexputil.ExpressionOrFunc(bslm.AllArithmeticOperators(), regexputil.EscapeRegex)
// Результат: \+|-|\*|/|%
```

## Примеры использования

### Ключевые слова управления потоком
```go
func ControlKeywordsRule() *models.Rule {
    patterns := []*models.Rule{
        {
            Name: "keyword.control.bsl",
            Match: fmt.Sprintf(`(?i:(?<=[^\wа-яё\.]|^)(%s)(?=[^\wа-яё\.]|$))`,
                regexputil.ExpressionOrFunc([]bslm.BslKeywords{
                    bslm.BslReturn, bslm.BslContinue, bslm.BslBrake
                }, nil)),
        },
        {
            Name: "keyword.control.conditional.bsl",
            Match: fmt.Sprintf(`(?i:(?<=[^\wа-яё\.]|^)(%s)(?=[^\wа-яё\.]|$))`,
                regexputil.ExpressionOrFunc([]bslm.BslKeywords{
                    bslm.BslIf, bslm.BslElse, bslm.BslElseIf, bslm.BslТогда, bslm.BslEndIf
                }, nil)),
        },
    }
    return newRule(ControlKeywordsKey(), patterns)
}
```

### Операторы
```go
func SupportOperators() *models.Rule {
    patterns := []*models.Rule{
        {
            Name: "keyword.operator.comparison.bsl",
            Match: fmt.Sprintf(`%s`,
                regexputil.ExpressionOrFunc(bslm.AllComparisonOperators(), regexputil.EscapeRegex)),
        },
        {
            Name: "keyword.operator.arithmetic.bsl",
            Match: fmt.Sprintf(`%s`,
                regexputil.ExpressionOrFunc(bslm.AllArithmeticOperators(), regexputil.EscapeRegex)),
        },
    }
    return newRule(SupportOperatorsKey(), patterns)
}
```

### Константы
```go
func ConstLiterals() *models.Rule {
    patterns := []*models.Rule{
        {
            Name: "constant.language.bsl",
            Match: fmt.Sprintf(`(?i:(?<=[^\wа-яё\.]|^)(%s)(?=[^\wа-яё\.]|$))`,
                regexputil.ExpressionOrFunc(bslm.AllConstLiterals(), nil)),
        },
    }
    return newRule(ConstLiteralsKey(), patterns)
}
```

## Использование литералов для символов

Мы также используем литералы для символов и пунктуации:

```go
// Вместо: Match: `(,)`
Match: fmt.Sprintf(`(%s)`, bslm.BslSemicolon)

// Вместо: Match: `(\()`
Match: fmt.Sprintf(`(%s)`, bslm.BslParenOpen)

// Вместо: Match: `(\))`
Match: fmt.Sprintf(`(%s)`, bslm.BslParenClose)
```

## Преимущества для поддержки

1. **Легкость добавления новых ключевых слов**: достаточно добавить константу в `keywords.go`
2. **Автоматическое обновление регулярных выражений**: при изменении ключевых слов все связанные правила обновляются автоматически
3. **Меньше ошибок**: исключается возможность опечаток в регулярных выражениях
4. **Лучшая читаемость**: код становится более понятным и самодокументируемым
5. **Консистентность**: все правила используют одинаковый подход к формированию паттернов

## Заключение

Использование `ExpressionOrFunc` для динамического формирования регулярных выражений делает код более поддерживаемым, читаемым и менее подверженным ошибкам. Этот подход особенно эффективен для языков программирования с большим количеством ключевых слов и операторов, таких как BSL. 