package models

// RegexPatterns содержит часто используемые паттерны регулярных выражений
const (
	// WordBoundary - граница слова (включая кириллицу)
	WordBoundary = `[^\wа-яё\.]`

	// WordBoundaryStart - граница слова в начале строки
	WordBoundaryStart = `[^\wа-яё\.]|^`

	// WordBoundaryEnd - граница слова в конце строки
	WordBoundaryEnd = `[^\wа-яё\.]|$`

	// Identifier - идентификатор (буква, цифра, подчеркивание, кириллица)
	Identifier = `[_$[:alpha:]][_$[:alnum:]]*`

	// IdentifierWithCyrillic - идентификатор с поддержкой кириллицы
	IdentifierWithCyrillic = `[a-zа-яё0-9_]+`

	// IdentifierWithCyrillicCaseInsensitive - идентификатор с кириллицей без учета регистра
	IdentifierWithCyrillicCaseInsensitive = `(?i:[a-zа-яё0-9_]+)`

	// WordBoundaryLookBehind - позитивный lookbehind для границы слова
	WordBoundaryLookBehind = `(?<=[^\wа-яё\.]|^)`

	// WordBoundaryLookAhead - позитивный lookahead для границы слова
	WordBoundaryLookAhead = `(?=[^\wа-яё\.]|$)`

	// CaseInsensitiveWordBoundary - граница слова без учета регистра
	CaseInsensitiveWordBoundary = `(?i:(?<=[^\wа-яё\.]|^)(.*?)(?=[^\wа-яё\.]|$))`

	// StatementEnd - конец оператора (точка с запятой или ключевые слова)
	StatementEnd = `(?i:(?=(;|Иначе|Конец)))`

	// DateLiteral - литерал даты
	DateLiteral = `'((\d{4}[^\d']*\d{2}[^\d']*\d{2})([^\d']*\d{2}[^\d']*\d{2}([^\d']*\d{2})?)?)'`

	// NumericLiteral - числовой литерал
	NumericLiteral = `(\d+\.?\d*)`

	// EscapedQuote - экранированная кавычка
	EscapedQuote = `""`
)

// RegexTemplates содержит шаблоны для часто используемых конструкций
const (
	// WordBoundaryTemplate - шаблон для границы слова
	WordBoundaryTemplate = `(?<=%s|^)(%s)(?=%s|$)`

	// CaseInsensitiveWordBoundaryTemplate - шаблон для границы слова без учета регистра
	CaseInsensitiveWordBoundaryTemplate = `(?i:(?<=%s|^)(%s)(?=%s|$))`
)
