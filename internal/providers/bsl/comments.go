package bsl

import "github.com/IgorKilipenko/go-tml-builder/internal/core/models"

func CommentRules() []*models.Rule {
	return []*models.Rule{
		{
			Name:  "comment.line.double-slash.bsl",
			Match: `//.*`,
		},
		{
			Name:  "comment.block.documentation.bsl",
			Begin: `/\*`,
			End:   `\*/`,
			Captures: map[string]models.Capture{
				"0": {Name: "punctuation.definition.comment.bsl"},
			},
		},
	}
}
