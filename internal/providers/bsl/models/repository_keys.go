package models

import "github.com/IgorKilipenko/go-tml-builder/internal/core/models"

const (
	KeyBasic         models.RepositoryKey = "basic"
	KeyMiscellaneous models.RepositoryKey = "miscellaneous"

	KeySupportOperators    models.RepositoryKey = "supportOperators"
	KeyLogicalOperators    models.RepositoryKey = "logicalOperators"
	KeyComparisonOperators models.RepositoryKey = "comparisonOperators"
	KeyTernaryOperators    models.RepositoryKey = "ternaryOperators"
	KeyStatementOperators  models.RepositoryKey = "statementOperators"

	KeyStringWithSingleValue models.RepositoryKey = "stringWithSingleSupportValue"
	KeyStringSupportValues   models.RepositoryKey = "stringSupportValues"
	KeyQuotedString          models.RepositoryKey = "quotedString"
	KeyExtensionRegions      models.RepositoryKey = "extensionRegions"

	KeyConstantsLiterals models.RepositoryKey = "constantsLiterals"

	KeyCommentBlock          models.RepositoryKey = "commentBlock"
	KeyCommentLine           models.RepositoryKey = "commentLine"
	KeyDeveloperCommentLine  models.RepositoryKey = "developerCommentLine"
	KeyFunctionDocumentation models.RepositoryKey = "functionDocumentation"

	KeyObjectDefinition models.RepositoryKey = "objectDefinition" // like: Новый Структура

	// Functions
	KeyCallSupportFunctions models.RepositoryKey = "supportFunctions" // Вызов глобальных предопределенных функций

	// Control flow
	KeyControlKeywords    models.RepositoryKey = "controlKeywords"
	KeyConditional        models.RepositoryKey = "conditional"
	KeyVariableAssignment models.RepositoryKey = "variableAssignment"
	KeyFunctionDefinition models.RepositoryKey = "functionDefinition"
	KeyFunctionEnd        models.RepositoryKey = "functionEnd"

	// Variables and types
	KeyVariableDefinition models.RepositoryKey = "variableDefinition"
	KeyStorageModifiers   models.RepositoryKey = "storageModifiers"
	KeyAnnotations        models.RepositoryKey = "annotations"

	// Block entities
	KeyBlockEntities         models.RepositoryKey = "blockEntities"
	KeyBlockVariables        models.RepositoryKey = "blockVariables"
	KeyBlockFunctions        models.RepositoryKey = "blockFunctions"
	KeyBlockObjectProperties models.RepositoryKey = "blockObjectProperties"
	KeyArrayLiteral          models.RepositoryKey = "arrayLiteral"
	KeyBlockAwait            models.RepositoryKey = "blockAwait"

	// Support types
	KeySupportEnums            models.RepositoryKey = "supportEnums"
	KeySupportClasses          models.RepositoryKey = "supportClasses"
	KeySupportValueTypes       models.RepositoryKey = "supportValueTypes"
	KeySupportLanguageConstant models.RepositoryKey = "supportLanguageConstant"
	KeySupportRegisterTable    models.RepositoryKey = "supportRegisterTable"

	// String patterns
	KeyQuotedStringBody       models.RepositoryKey = "quotedStringBody"
	KeyStringPatternParameter models.RepositoryKey = "stringPatternParameter"
	KeyStringInnerSupports    models.RepositoryKey = "_stringInnerSupports"
	KeyQuery                  models.RepositoryKey = "query"

	// Regions
	KeyMainRegion      models.RepositoryKey = "mainRegion"
	KeyMainRegionStart models.RepositoryKey = "mainRegionStart"
	KeyMainRegionEnd   models.RepositoryKey = "mainRegionEnd"

	// Documentation
	KeyDocParametersBlock     models.RepositoryKey = "docParametersBlock"
	KeyDocReturnsBlock        models.RepositoryKey = "docReturnsBlock"
	KeyDocParametersBlockArgs models.RepositoryKey = "_docParametersBlockArgs"
	KeyDocReturnsBlockArgs    models.RepositoryKey = "_docReturnsBlockArgs"
)
