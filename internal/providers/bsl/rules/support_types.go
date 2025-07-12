package rules

import (
	"fmt"

	"github.com/IgorKilipenko/go-tml-builder/internal/core/models"
	"github.com/IgorKilipenko/go-tml-builder/internal/core/regexputil"
	bslm "github.com/IgorKilipenko/go-tml-builder/internal/providers/bsl/models"
)

func SupportEnumsKey() models.RepositoryKey {
	return bslm.KeySupportEnums
}

func SupportClassesKey() models.RepositoryKey {
	return bslm.KeySupportClasses
}

func SupportValueTypesKey() models.RepositoryKey {
	return bslm.KeySupportValueTypes
}

func SupportLanguageConstantKey() models.RepositoryKey {
	return bslm.KeySupportLanguageConstant
}

func SupportRegisterTableKey() models.RepositoryKey {
	return bslm.KeySupportRegisterTable
}

// SupportEnums правила для перечислений
func SupportEnums() *models.Rule {
	patterns := []*models.Rule{
		{
			Match: `(?i:)(?<![\.\)\]]\s*)\b(ОбходРезультатаЗапроса)\b(?=\s*\.)`,
			Name:  "support.class.bsl",
		},
		{
			Match: `(?i:)(?<![\.\)\]]\s*)\b(РежимЗаписиДокумента|РежимПроведенияДокумента)\b(?=\s*\.)`,
			Name:  "support.class.bsl",
		},
		{
			Match: `(?i:)(?<![\.\)\]]\s*)\b(РежимДиалогаВопрос|КодВозвратаДиалога)\b(?=\s*\.)`,
			Name:  "support.class.bsl",
		},
		{
			Match: `(?i:)(?<![\.\)\]]\s*)\b(ВидДвиженияНакопления|ВидПериодаРегистраРасчета)\b(?=\s*\.)`,
			Name:  "support.class.bsl",
		},
		{
			Match: `(?i:)(?<![\.\)\]]\s*)\b(РежимБлокировкиДанных|ВидГраницы)\b(?=\s*\.)`,
			Name:  "support.class.bsl",
		},
	}

	return newRule(SupportEnumsKey(), patterns)
}

// SupportClasses правила для поддерживаемых классов
func SupportClasses() *models.Rule {
	patterns := []*models.Rule{
		SupportRegisterTableKey().IncludeRef(),
		{
			Match: `(?i:)(?<![\.\)\]]\s*)\b(WSСсылки|БиблиотекаКартинок|БиблиотекаМакетовОформленияКомпоновкиДанных|БиблиотекаСтилей|БизнесПроцессы|ВнешниеИсточникиДанных|ВнешниеОбработки|ВнешниеОтчеты|Документы|ДоставляемыеУведомления|ЖурналыДокументов|Задачи|ИнформацияОбИнтернетСоединении|ИспользованиеРабочейДаты|ИсторияРаботыПользователя|Константы|КритерииОтбора|Метаданные|Обработки|ОтправкаДоставляемыхУведомлений|Отчеты|ПараметрыСеанса|Перечисления|ПланыВидовРасчета|ПланыВидовХарактеристик|ПланыОбмена|ПланыСчетов|ПолнотекстовыйПоиск|ПользователиИнформационнойБазы|Последовательности|РасширенияКонфигурации|РегистрыБухгалтерии|РегистрыНакопления|РегистрыРасчета|РегистрыСведений|РегламентныеЗадания|СериализаторXDTO|Справочники|СредстваГеопозиционирования|СредстваКриптографии|СредстваМультимедиа|СредстваОтображенияРекламы|СредстваПочты|СредстваТелефонии|ФабрикаXDTO|ФайловыеПотоки|ФоновыеЗадания|ХранилищаНастроек|ВстроенныеПокупки|ОтображениеРекламы|ПанельЗадачОС|ПроверкаВстроенныхПокупок)\b(?=\s*\.)`,
			Name:  "support.class.bsl",
		},
		{
			Match: `(?i:)(?<![\.\)\]]\s*)\b(ОбработкаОшибок)\b(?=\s*\.)`,
			Name:  "support.class.bsl",
		},
	}

	return newRule(SupportClassesKey(), patterns)
}

// SupportValueTypes правила для поддерживаемых типов значений
func SupportValueTypes() *models.Rule {
	patterns := []*models.Rule{
		{
			Match: fmt.Sprintf(`\b(?i:%s)\b`,
				regexputil.ExpressionOrFunc(bslm.AllUniversalTypes(), nil)),
			Name: "support.type.bsl",
		},
	}

	return newRule(SupportValueTypesKey(), patterns)
}

// SupportLanguageConstant правила для языковых констант
func SupportLanguageConstant() *models.Rule {
	patterns := []*models.Rule{
		{
			Match: `(?i:)(?<![\.\)\]]\s*)\b(ЭтотОбъект|ЭтаФорма)\b`,
			Name:  "constant.language.bsl",
		},
	}

	return newRule(SupportLanguageConstantKey(), patterns)
}

// SupportRegisterTable правила для таблиц регистров
func SupportRegisterTable() *models.Rule {
	patterns := []*models.Rule{
		{
			Match: `(?i:)(?<![\.\)\]]\s*)\b(Регистры(?:Накопления|Сведений|Бухгалтерии))\b(?=\s*\.)`,
			Name:  "support.class.bsl",
		},
	}

	return newRule(SupportRegisterTableKey(), patterns)
}
