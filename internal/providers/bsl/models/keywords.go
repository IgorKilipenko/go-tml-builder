package models

type BslKeywords string

const (
	BslVar      BslKeywords = "Перем"
	BslNew      BslKeywords = "Новый"
	BslExport   BslKeywords = "Экспорт"
	BslReturn   BslKeywords = "Возврат"
	BslContinue BslKeywords = "Продолжить"
	BslBrake    BslKeywords = "Прервать"

	BslAsync BslKeywords = "Асинх"
	BslAwait BslKeywords = "Ждать"

	BslIf     BslKeywords = "Если"
	BslElse   BslKeywords = "Иначе"
	BslElseIf BslKeywords = "ИначеЕсли"
	BslТогда  BslKeywords = "Then"
	BslEndIf  BslKeywords = "КонецЕсли"

	BslFor     BslKeywords = "Для"
	BslEach    BslKeywords = "Каждого"
	BslIn      BslKeywords = "Из"
	BslLoop    BslKeywords = "Цикл"
	BslEndLoop BslKeywords = "КонецЦикла"
	BslWhile   BslKeywords = "Пока"

	BslTry       BslKeywords = "Попытка"
	BslCatch     BslKeywords = "Исключение"
	BslEndTry    BslKeywords = "КонецПопытки"
	BslException BslKeywords = "ВызватьИсключение"

	// Function keywords
	BslProcedure BslKeywords = "Процедура"
	BslFunction  BslKeywords = "Функция"
	BslEndProc   BslKeywords = "КонецПроцедуры"
	BslEndFunc   BslKeywords = "КонецФункции"
	BslVal       BslKeywords = "Знач"

	// Import and directives
	BslUse       BslKeywords = "Использовать"
	BslNative    BslKeywords = "native"
	BslOnClient  BslKeywords = "НаКлиенте"
	BslOnServer  BslKeywords = "НаСервере"
	BslNoContext BslKeywords = "БезКонтекста"
)

func (k BslKeywords) String() string {
	return string(k)
}

type BslConstants string

const (
	BslUndefined BslConstants = "Неопределено"
	BslNull      BslConstants = "Null"
	BslTrue      BslConstants = "Истина"
	BslFalse     BslConstants = "Ложь"
)

func (k BslConstants) String() string {
	return string(k)
}

func AllConstLiterals() []BslConstants {
	return []BslConstants{
		BslUndefined,
		BslNull,
		BslTrue,
		BslFalse,
	}
}

func AllControlKeywords() []BslKeywords {
	return []BslKeywords{
		BslReturn,
		BslContinue,
		BslBrake,
		BslIf,
		BslElse,
		BslElseIf,
		BslТогда,
		BslEndIf,
		BslFor,
		BslEach,
		BslIn,
		BslLoop,
		BslEndLoop,
		BslWhile,
		BslTry,
		BslCatch,
		BslEndTry,
		BslException,
	}
}

func AllFunctionKeywords() []BslKeywords {
	return []BslKeywords{
		BslProcedure,
		BslFunction,
		BslEndProc,
		BslEndFunc,
		BslVal,
	}
}

type BslPunctuation string

const (
	BslSemicolon BslPunctuation = ";"
	BslCondOp    BslPunctuation = "?"

	BslParenOpen  BslPunctuation = "("
	BslParenClose BslPunctuation = ")"

	BslCurlyOpen  BslPunctuation = "{"
	BslCurlyClose BslPunctuation = "}"

	BslSquareOpen  BslPunctuation = "["
	BslSquareClose BslPunctuation = "]"
)

func (k BslPunctuation) String() string {
	return string(k)
}

type BslOperators string

const (
	// Операторы сравнения (Comparison Operators)
	BslEqual          BslOperators = "="
	BslNotEqual       BslOperators = "<>"
	BslLess           BslOperators = "<"
	BslLessOrEqual    BslOperators = "<="
	BslGreater        BslOperators = ">"
	BslGreaterOrEqual BslOperators = "=>"

	// Арифметические операторы (Arithmetic Operators)
	BslAdd      BslOperators = "+"
	BslSubtract BslOperators = "-"
	BslMultiply BslOperators = "*"
	BslDivide   BslOperators = "/"
	BslMod      BslOperators = "%"

	// Логические операторы (Logical Operators)
	BslAnd BslOperators = "И"
	BslOr  BslOperators = "ИЛИ"
	BslNot BslOperators = "НЕ"
)

func (k BslOperators) String() string {
	return string(k)
}

// AllComparisonOperators операторы сравнения: [=, <>, <, >, =>]
func AllComparisonOperators() []BslOperators {
	return []BslOperators{
		BslEqual,
		BslNotEqual,
		BslLess,
		BslLessOrEqual,
		BslGreater,
		BslGreaterOrEqual,
	}
}

// AllArithmeticOperators логические операторы: [+, -, *, /, %]
func AllArithmeticOperators() []BslOperators {
	return []BslOperators{
		BslAdd,
		BslSubtract,
		BslMultiply,
		BslDivide,
		BslMod,
	}
}

// AllLogicalOperators арифметические операторы: [И, ИЛИ, НЕ]
func AllLogicalOperators() []BslOperators {
	return []BslOperators{
		BslAnd,
		BslOr,
		BslNot,
	}
}

type BslUniversalTypes string

const (
	BslStr         BslUniversalTypes = "Строка"
	BslDate        BslUniversalTypes = "Дата"
	BslStruct      BslUniversalTypes = "Структура"
	BslFixedStruct BslUniversalTypes = "ФиксированнаяСтруктура"
	BslArr         BslUniversalTypes = "Массив"
	BslFixedArr    BslUniversalTypes = "ФиксированныйМассив"
	BslMap         BslUniversalTypes = "Соответствие"
	BslFixedMap    BslUniversalTypes = "ФиксированноеСоответствие"
	BslTree        BslUniversalTypes = "Дерево"
	BslValTable    BslUniversalTypes = "ТаблицаЗначений"
)

func (k BslUniversalTypes) String() string {
	return string(k)
}

func AllUniversalTypes() []BslUniversalTypes {
	return []BslUniversalTypes{
		BslStr,
		BslDate,
		BslStruct,
		BslFixedStruct,
		BslArr,
		BslFixedArr,
		BslMap,
		BslFixedMap,
		BslTree,
		BslValTable,
	}
}

type BslModelTypes string

var allModelTypes []BslModelTypes

func registerModelType(val BslModelTypes) BslModelTypes {
	allModelTypes = append(allModelTypes, val)
	return val
}

const (
	DocumentObject BslModelTypes = "ДокументОбъект"
	DocumentRef    BslModelTypes = "ДокументСсылка"
	CatalogObject  BslModelTypes = "СправочникОбъект"
	CatalogRef     BslModelTypes = "СправочникСсылка"
)

func AllModelTypes() []BslModelTypes {
	return []BslModelTypes{
		DocumentObject,
		DocumentRef,
		CatalogObject,
		CatalogRef,
	}
}
