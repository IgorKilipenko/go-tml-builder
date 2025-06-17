package models

type BslKeywords string

const (
	Var      BslKeywords = "Перем"
	New      BslKeywords = "Новый"
	Export   BslKeywords = "Экспорт"
	Return   BslKeywords = "Возврат"
	Continue BslKeywords = "Продолжить"
	Brake    BslKeywords = "Прервать"

	Async BslKeywords = "Асинх"
	Wait  BslKeywords = "Ждать"

	If     BslKeywords = "Если"
	Else   BslKeywords = "Иначе"
	ElseIf BslKeywords = "ИначеЕсли"
	Тогда  BslKeywords = "Then"
	EndIf  BslKeywords = "КонецЕсли"

	For     BslKeywords = "Для"
	Each    BslKeywords = "Каждого"
	In      BslKeywords = "Из"
	Loop    BslKeywords = "Цикл"
	EndLoop BslKeywords = "КонецЦикла"
	While   BslKeywords = "Пока"

	Try       BslKeywords = "Попытка"
	Catch     BslKeywords = "Исключение"
	EndTry    BslKeywords = "КонецПопытки"
	Exception BslKeywords = "ВызватьИсключение"
)

type BslConstants string

const (
	Undefined BslConstants = "Неопределено"
	Null      BslConstants = "Null"
	True      BslConstants = "Истина"
	False     BslConstants = "Ложь"
)

func (k BslConstants) String() string {
	return string(k)
}

func AllConstLiterals() []BslConstants {
	return []BslConstants{
		Undefined,
		Null,
		True,
		False,
	}
}

type BslPunctuation string

const (
	EndSignature BslPunctuation = ";"

	StartStruct BslPunctuation = "{"
	EndStruct   BslPunctuation = "}"

	StartArray BslPunctuation = "["
	EndArray   BslPunctuation = "]"
)

type BslOperators string

const (
	Equal         BslOperators = "="
	Lower         BslOperators = "<"
	LowerOrEqual  BslOperators = "<="
	Bigger        BslOperators = ">"
	BiggerOrEqual BslOperators = "=>"

	Plus     BslOperators = "+"
	Minus    BslOperators = "-"
	Multiple BslOperators = "*"
	Divide   BslOperators = "/"
	Mode     BslOperators = "%"

	And BslOperators = "И"
	Or  BslOperators = "ИЛИ"
	Not BslOperators = "НЕ"
)

type BslUniversalTypes string

const (
	String         BslUniversalTypes = "Строка"
	Date           BslUniversalTypes = "Дата"
	Strict         BslUniversalTypes = "Структура"
	ReadOnlyStrict BslUniversalTypes = "ФиксированнаяСтруктура"
	Array          BslUniversalTypes = "Массив"
	ReadOnlyArray  BslUniversalTypes = "ФиксированныйМассив"
	Map            BslUniversalTypes = "Соответствие"
	ReadOnlyMap    BslUniversalTypes = "ФиксированноеСоответствие"
	Tree           BslUniversalTypes = "Дерево"
	Table          BslUniversalTypes = "ТаблицаЗначений"
)

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
