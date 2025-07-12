# Changelog

## [1.2.0] - 2024-12-19

### Добавлено
- **Константы регулярных выражений** для улучшения читаемости и поддерживаемости
- Новый файл `internal/providers/bsl/models/regex_patterns.go` с константами
- Примеры использования констант в `examples/regex_patterns_demo.go`
- Документация по константам в `docs/regex_patterns_constants.md`

### Изменено
- **Обновлены все файлы правил** для использования констант регулярных выражений:
  - `internal/providers/bsl/rules/control_flow.go` - использование `WordBoundaryLookBehind`, `WordBoundaryLookAhead`
  - `internal/providers/bsl/rules/variables.go` - использование `IdentifierWithCyrillic`
  - `internal/providers/bsl/rules/rules.go` - использование `NumericLiteral`, `DateLiteral`
  - `internal/providers/bsl/rules/const_literals.go` - использование `WordBoundaryLookBehind`
  - `internal/providers/bsl/rules/new_objects.go` - использование `Identifier`
  - `internal/providers/bsl/rules/blocks.go` - использование `Identifier`, `EscapedQuote`
- **Обновлен README.md** с описанием констант и примерами использования

### Улучшения
- **Повышенная читаемость**: сложные паттерны стали понятными благодаря именованным константам
- **Лучшая поддерживаемость**: изменения в сложных паттернах нужно делать только в одном месте
- **Консистентность**: все правила используют одинаковые сложные паттерны через константы
- **Меньше ошибок**: исключена возможность опечаток в сложных регулярных выражениях
- **Баланс**: простые символы остаются читаемыми, сложные - вынесены в константы

### Категории констант
- **Границы слов**: `WordBoundary`, `WordBoundaryLookBehind`, `WordBoundaryLookAhead`
- **Идентификаторы**: `Identifier`, `IdentifierWithCyrillic`, `IdentifierWithCyrillicCaseInsensitive`
- **Литералы**: `NumericLiteral`, `DateLiteral`, `EscapedQuote`
- **Специальные паттерны**: `StatementEnd`, `CaseInsensitiveWordBoundary`
- **Шаблоны**: `WordBoundaryTemplate`, `CaseInsensitiveWordBoundaryTemplate`

### Примеры изменений
```go
// Было:
Match: `(?i:(?<=[^\wа-яё\.]|^)(Если)(?=[^\wа-яё\.]|$))`
Begin: `(?i:(?<=;|^)\s*(Если))`

// Стало:
Match: fmt.Sprintf(`(?i:%s(%s)%s)`, 
    bslm.WordBoundaryLookBehind, "Если", bslm.WordBoundaryLookAhead)
Begin: `(?i:(?<=;|^)\s*(Если))`  // Простые символы напрямую
```

## [1.1.0] - 2024-12-19

### Добавлено
- **Динамическое формирование регулярных выражений** с использованием `ExpressionOrFunc`
- Централизованное управление ключевыми словами языка BSL
- Поддержка препроцессоров для экранирования и модификации регулярных выражений
- Примеры использования в `examples/dynamic_regex_example.go`
- Документация по новому подходу в `docs/dynamic_regex_approach.md`

### Изменено
- **Обновлены все файлы правил** для использования литералов вместо жестко заданных строк:
  - `internal/providers/bsl/rules/control_flow.go` - ключевые слова управления потоком
  - `internal/providers/bsl/rules/variables.go` - переменные и операторы
  - `internal/providers/bsl/rules/rules.go` - базовые правила и пунктуация
  - `internal/providers/bsl/rules/const_literals.go` - константы языка
  - `internal/providers/bsl/rules/new_objects.go` - конструкторы объектов
  - `internal/providers/bsl/rules/blocks.go` - блоки кода и скобки
- **Обновлен README.md** с описанием нового подхода и примерами

### Улучшения
- **Лучшая поддерживаемость**: добавление новых ключевых слов требует изменения только одного файла
- **Меньше ошибок**: исключена возможность опечаток в регулярных выражениях
- **Повышенная читаемость**: код стал более понятным и самодокументируемым
- **Консистентность**: все правила используют единый подход к формированию паттернов

### Технические детали
- Использование `fmt.Sprintf` для динамического формирования регулярных выражений
- Применение `regexputil.ExpressionOrFunc` для объединения ключевых слов
- Использование `regexputil.EscapeRegex` для экранирования специальных символов
- Замена жестко заданных символов на литералы из `bslm` пакета

### Примеры изменений
```go
// Было:
Match: `\b(Если|Тогда|ИначеЕсли|Иначе|КонецЕсли)\b`

// Стало:
Match: fmt.Sprintf(`(?i:(?<=[^\wа-яё\.]|^)(%s)(?=[^\wа-яё\.]|$))`,
    regexputil.ExpressionOrFunc([]bslm.BslKeywords{
        bslm.BslIf, bslm.BslElse, bslm.BslElseIf, bslm.BslТогда, bslm.BslEndIf
    }, nil))
```

## [1.0.0] - 2024-12-18

### Добавлено
- Базовая функциональность генератора грамматики TextMate для BSL
- Поддержка всех основных элементов языка 1С
- Модульная архитектура с разделением правил по категориям
- Генерация валидного JSON файла грамматики 