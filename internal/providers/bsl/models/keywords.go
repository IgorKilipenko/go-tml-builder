package models

type BslKeywords string

const (
	Var      BslKeywords = "Перем"
	New      BslKeywords = "Новый"
	Export   BslKeywords = "Экспорт"
	Return   BslKeywords = "Возврат"
	Continue BslKeywords = "Продолжить"
	Brake    BslKeywords = "Прервать"
	And      BslKeywords = "И"
	Or       BslKeywords = "ИЛИ"

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
	True      BslConstants = "Истина"
	False     BslConstants = "Ложь"
)

func (k BslConstants) String() string {
	return string(k)
}

func AllConstants() []BslConstants {
	return []BslConstants {
		Undefined,
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

	Equal         BslPunctuation = "="
	Lower         BslPunctuation = "<"
	LowerOrEqual  BslPunctuation = "<="
	Bigger        BslPunctuation = ">"
	BiggerOrEqual BslPunctuation = "=>"
)

type BslBasicTypes string

const (
	String         BslBasicTypes = "Строка"
	Date           BslBasicTypes = "Дата"
	Strict         BslBasicTypes = "Структура"
	ReadOnlyStrict BslBasicTypes = "ФиксированнаяСтруктура"
	Array          BslBasicTypes = "Массив"
	ReadOnlyArray  BslBasicTypes = "ФиксированныйМассив"
	Map            BslBasicTypes = "Соответствие"
	ReadOnlyMap    BslBasicTypes = "ФиксированноеСоответствие"
)

type BslModelTypes string

const (
	DocumentObject BslModelTypes = "ДокументОбъект"
	DocumentRef    BslModelTypes = "ДокументСсылка"
	CatalogObject  BslModelTypes = "СправочникОбъект"
	CatalogRef     BslModelTypes = "СправочникСсылка"
)
