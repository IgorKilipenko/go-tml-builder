package rules

import (
	"fmt"

	"github.com/IgorKilipenko/go-tml-builder/internal/core/models"
	"github.com/IgorKilipenko/go-tml-builder/internal/core/regexputil"
	bslm "github.com/IgorKilipenko/go-tml-builder/internal/providers/bsl/models"
)

// StringType вспомогательный тип для работы с regexputil
type StringType string

func (s StringType) String() string {
	return string(s)
}

func CallSupportFunctionsKey() models.RepositoryKey {
	return bslm.KeyCallSupportFunctions
}

// CallSupportFunctions правила для вызова глобальных предопределенных функций
func CallSupportFunctions() *models.Rule {
	patterns := []*models.Rule{
		{
			Name: "support.function.bsl",
			Match: fmt.Sprintf("(?x)(?i:(?<=[^\\wа-яё\\.]|^)(%s)\\s*(?=\\())",
				regexputil.ExpressionOrFunc(AllSupportFunctionsNames(), nil)),
		},
	}

	return newRule(CallSupportFunctionsKey(), patterns)
}

// AllSupportFunctionsNames возвращает все имена поддерживаемых функций
func AllSupportFunctionsNames() []StringType {
	return []StringType{
		// Строковые функции
		"СтрДлина", "СокрЛ", "СокрП", "СокрЛП", "Лев", "Прав", "Сред", "СтрНайти",
		"ВРег", "НРег", "ТРег", "Символ", "КодСимвола", "ПустаяСтрока", "СтрЗаменить",
		"СтрЧислоСтрок", "СтрПолучитьСтроку", "СтрЧислоВхождений", "СтрСравнить",
		"СтрНачинаетсяС", "СтрЗаканчиваетсяНа", "СтрРазделить", "СтрСоединить",

		// Математические функции
		"Цел", "Окр", "ACos", "ASin", "ATan", "Cos", "Exp", "Log", "Log10", "Pow", "Sin", "Sqrt", "Tan",

		// Работа с датами
		"Год", "Месяц", "День", "Час", "Минута", "Секунда", "ТекущаяДата",
		"НачалоГода", "НачалоДня", "НачалоКвартала", "НачалоМесяца", "КонецГода", "КонецДня",

		// Работа с типами
		"Тип", "ТипЗнч", "Булево", "Число", "Строка", "Дата",

		// Диалоги
		"ПоказатьВопрос", "Вопрос", "ПоказатьПредупреждение", "Предупреждение", "Сообщить",
		"ПоказатьЗначение", "ОткрытьЗначение", "Оповестить",

		// Форматирование
		"Формат", "ЧислоПрописью", "НСтр", "ПредставлениеПериода", "СтрШаблон",

		// Работа с файлами
		"КопироватьФайл", "ПереместитьФайл", "УдалитьФайлы", "НайтиФайлы", "СоздатьКаталог",
		"ПолучитьИмяВременногоФайла", "РазделитьФайл", "ОбъединитьФайлы",

		// Работа с ИБ
		"НачатьТранзакцию", "ЗафиксироватьТранзакцию", "ОтменитьТранзакцию",
		"ПолучитьОперативнуюОтметкуВремени", "ТранзакцияАктивна",

		// Работа с XML/JSON
		"XMLСтрока", "XMLЗначение", "ЗаписатьJSON", "ПрочитатьJSON",

		// Прочие
		"Мин", "Макс", "ОписаниеОшибки", "Вычислить", "ИнформацияОбОшибке",
		"Base64Значение", "Base64Строка", "ЗаполнитьЗначенияСвойств", "ЗначениеЗаполнено",
	}
}
