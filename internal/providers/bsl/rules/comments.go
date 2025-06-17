package rules

import (
	"github.com/IgorKilipenko/go-tml-builder/internal/core/models"
	bslm "github.com/IgorKilipenko/go-tml-builder/internal/providers/bsl/models"
)

func CommentLineKey() models.RepositoryKey {
	return bslm.KeyCommentLine
}

func CommentBlockKey() models.RepositoryKey {
	return bslm.KeyCommentBlock
}

func DeveloperCommentLineKey() models.RepositoryKey {
	return bslm.KeyDeveloperCommentLine
}

func FunctionDocumentationKey() models.RepositoryKey {
	return bslm.KeyFunctionDocumentation
}

func CommentLine() *models.Rule {
	patterns := []*models.Rule{
		DeveloperCommentLineKey().IncludeRef(),
		{
			Name:  "comment.line.double-slash.bsl",
			Match: `//.*$`,
		},
	}

	return newRule(CommentLineKey(), patterns)
}

func CommentBlock() *models.Rule {
	patterns := []*models.Rule{
		FunctionDocumentationKey().IncludeRef(),
		{
			Match: `^\s*(//\s*)(Устарела)(\.\s*.*)\s*$`,
			Captures: map[string]models.Capture{
				"1": {Name: "comment.line.double-slash.bsl"},
				"2": {Name: "keyword.deprecated.bsl"},
				"3": {Name: "comment.line.double-slash.bsl"},
			},
		},
		CommentLineKey().IncludeRef(),
	}

	commentBlockRule := newRule(CommentBlockKey(), patterns)
	commentBlockRule.Name = "comment.line.double-slash.bsl"
	commentBlockRule.Begin = `^\s*(?=//\s*)`
	commentBlockRule.End = `(?=^(?!\s*//))`

	return commentBlockRule
}
