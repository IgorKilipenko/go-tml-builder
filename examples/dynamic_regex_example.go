package main

import (
	"fmt"

	"github.com/IgorKilipenko/go-tml-builder/internal/core/regexputil"
	bslm "github.com/IgorKilipenko/go-tml-builder/internal/providers/bsl/models"
)

func main() {
	// Пример 1: Формирование регулярного выражения для ключевых слов управления потоком
	controlKeywords := bslm.AllControlKeywords()
	controlRegex := regexputil.ExpressionOrFunc(controlKeywords, nil)
	fmt.Printf("Control keywords regex: %s\n", controlRegex)
	// Вывод: Control keywords regex: Возврат|Продолжить|Прервать|Если|Иначе|ИначеЕсли|Then|КонецЕсли|Для|Каждого|Из|Цикл|КонецЦикла|Пока|Попытка|Исключение|КонецПопытки|ВызватьИсключение

	// Пример 2: Формирование регулярного выражения для логических операторов
	logicalOperators := bslm.AllLogicalOperators()
	logicalRegex := regexputil.ExpressionOrFunc(logicalOperators, nil)
	fmt.Printf("Logical operators regex: %s\n", logicalRegex)
	// Вывод: Logical operators regex: И|ИЛИ|НЕ

	// Пример 3: Формирование регулярного выражения для операторов сравнения с экранированием
	comparisonOperators := bslm.AllComparisonOperators()
	comparisonRegex := regexputil.ExpressionOrFunc(comparisonOperators, regexputil.EscapeRegex)
	fmt.Printf("Comparison operators regex: %s\n", comparisonRegex)
	// Вывод: Comparison operators regex: =|<>|<|<=|>|=>

	// Пример 4: Формирование регулярного выражения для арифметических операторов с экранированием
	arithmeticOperators := bslm.AllArithmeticOperators()
	arithmeticRegex := regexputil.ExpressionOrFunc(arithmeticOperators, regexputil.EscapeRegex)
	fmt.Printf("Arithmetic operators regex: %s\n", arithmeticRegex)
	// Вывод: Arithmetic operators regex: \+|-|\*|/|%

	// Пример 5: Формирование регулярного выражения для констант
	constants := bslm.AllConstLiterals()
	constantsRegex := regexputil.ExpressionOrFunc(constants, nil)
	fmt.Printf("Constants regex: %s\n", constantsRegex)
	// Вывод: Constants regex: Неопределено|Null|Истина|Ложь

	// Пример 6: Формирование полного регулярного выражения для условных конструкций
	conditionalPattern := fmt.Sprintf(`(?i:(?<=[^\wа-яё\.]|^)(%s)(?=[^\wа-яё\.]|$))`, controlRegex)
	fmt.Printf("Full conditional pattern: %s\n", conditionalPattern)
}
