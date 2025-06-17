package models

import "github.com/IgorKilipenko/go-tml-builder/internal/core/models"

const (
	KeyBasic         models.RepositoryKey = "basic"
	KeyMiscellaneous models.RepositoryKey = "miscellaneous"

	KeyStringWithSingleValue models.RepositoryKey = "stringWithSingleSupportValue"
	KeyStringSupportValues   models.RepositoryKey = "stringSupportValues"
	KeyQuotedString          models.RepositoryKey = "quotedString"
	KeyExtensionRegions      models.RepositoryKey = "extensionRegions"

	KeyConstantsLiterals models.RepositoryKey = "constantsLiterals"

	KeyCommentBlock          models.RepositoryKey = "commentBlock"
	KeyCommentLine           models.RepositoryKey = "commentLine"
	KeyDeveloperCommentLine  models.RepositoryKey = "developerCommentLine"
	KeyFunctionDocumentation models.RepositoryKey = "functionDocumentation"
)
