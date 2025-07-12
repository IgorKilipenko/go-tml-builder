package rules

import (
	"github.com/IgorKilipenko/go-tml-builder/internal/core/models"
	bslm "github.com/IgorKilipenko/go-tml-builder/internal/providers/bsl/models"
)

func MainRegionKey() models.RepositoryKey {
	return bslm.KeyMainRegion
}

func MainRegionStartKey() models.RepositoryKey {
	return bslm.KeyMainRegionStart
}

func MainRegionEndKey() models.RepositoryKey {
	return bslm.KeyMainRegionEnd
}

// MainRegion правила для основных областей
func MainRegion() *models.Rule {
	patterns := []*models.Rule{
		MainRegionStartKey().IncludeRef(),
		MainRegionEndKey().IncludeRef(),
	}

	return newRule(MainRegionKey(), patterns)
}

// MainRegionStart правила для начала областей
func MainRegionStart() *models.Rule {
	patterns := []*models.Rule{
		bslm.KeyCommentLine.IncludeRef(),
	}

	rule := newRule(MainRegionStartKey(), patterns)
	rule.Begin = `(?i)(#(Область))(?:\s+([\wа-яёА-ЯЁ]+))?`
	rule.End = `$`
	rule.BeginCaptures = map[string]models.Capture{
		"1": {Name: "keyword.other.section.bsl"},
		"3": {Name: "entity.name.section.bsl"},
	}

	return rule
}

// MainRegionEnd правила для окончания областей
func MainRegionEnd() *models.Rule {
	patterns := []*models.Rule{
		bslm.KeyCommentLine.IncludeRef(),
	}

	rule := newRule(MainRegionEndKey(), patterns)
	rule.Match = `(?i)(#(КонецОбласти))(?:(\s+//\s*)([\wа-яёА-ЯЁ]+)?)?`
	rule.Captures = map[string]models.Capture{
		"1": {Name: "keyword.other.section.bsl"},
		"3": {Name: "comment.line.double-slash.bsl"},
		"4": {Name: "entity.name.section.bsl"},
	}

	return rule
}
