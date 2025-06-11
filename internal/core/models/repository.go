package models

import (
	"log"
)

// RepositoryKey - типизированный ключ для репозитория
type RepositoryKey string

// Repository - хранилище именованных правил
type Repository map[RepositoryKey]*Rule

// Предопределенные ключи (можно расширять)
const (
	KeyBuiltinFuncs   RepositoryKey = "builtin_functions"
	KeyStrings        RepositoryKey = "strings"
	KeyNumbers        RepositoryKey = "numbers"
	KeyPreprocessor   RepositoryKey = "preprocessor"
	KeyFunctions      RepositoryKey = "functions"
	KeyFunctionParams RepositoryKey = "function_params"
	KeyConditions     RepositoryKey = "conditions"
	KeyComments       RepositoryKey = "comments"
	KeyOperators      RepositoryKey = "operators"
)

// ValidKeys — все допустимые ключи для валидации
func ValidKeys() []RepositoryKey {
	return []RepositoryKey{
		KeyBuiltinFuncs,
		KeyStrings,
		KeyNumbers,
		KeyFunctions,
		KeyComments,
		KeyPreprocessor,
	}
}

// String реализует Stringer
func (k RepositoryKey) String() string {
	return string(k)
}

// IncludeRef создает ссылку на правило в репозитории
func (k RepositoryKey) IncludeRef() *Rule {
	return &Rule{Include: k}
}

func Validate(repo Repository) error {
	for key, _ := range repo {
		// Проверка ключа
		if _, isStandard := isStandardKey(key); !isStandard {
			log.Printf("Warning: non-standard repository key '%s'", key)
		}

		// // Проверка ссылок в Include
		// if err := validateRuleIncludes(rule, repo); err != nil {
		// 	return fmt.Errorf("invalid rule '%s': %v", key, err)
		// }
	}
	return nil
}

func isStandardKey(key RepositoryKey) (bool, bool) {
	for _, k := range ValidKeys() {
		if k == key {
			return true, true
		}
	}
	return false, false
}
