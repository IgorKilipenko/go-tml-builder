package rules

import (
	"fmt"

	"github.com/IgorKilipenko/go-tml-builder/internal/core/models"
	bslm "github.com/IgorKilipenko/go-tml-builder/internal/providers/bsl/models"
)

func QueryKey() models.RepositoryKey {
	return bslm.KeyQuery
}

// Query правила для поддержки запросов BSL
func Query() *models.Rule {
	patterns := []*models.Rule{
		{
			Name:  "comment.line.double-slash.sdbl",
			Match: `(//(("")|[^"])*)`,
		},
		{
			Include: "extensionRegions",
		},
		{
			Name:  "string.quoted.double.sdbl",
			Match: `""[^"]*""`,
		},
		{
			Include: "source.sdbl",
		},
		{
			Name:  "comment.line.double-slash.bsl",
			Begin: `^\s*//`,
			End:   `$`,
		},
	}

	rule := newRule(QueryKey(), patterns)
	rule.Begin = fmt.Sprintf(`(?i)(?<=[^\wа-яё\.]|^)(%s(\s+%s)?(\s+%s)?(\s+%s)?)(?=[^\wа-яё\.]|$)`,
		bslm.BslSelect, bslm.BslAllowed, bslm.BslDistinct, bslm.BslFirst)
	rule.End = `(?=\"[^\"])`
	rule.BeginCaptures = map[string]models.Capture{
		"1": {Name: "keyword.control.sdbl"},
	}

	return rule
}
