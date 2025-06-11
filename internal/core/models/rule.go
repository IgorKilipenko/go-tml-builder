package models

type Rule struct {
	Name       string      `json:"name,omitempty"`
	Match      string      `json:"match,omitempty"`
	Begin      string      `json:"begin,omitempty"`
	End        string      `json:"end,omitempty"`
	While      string      `json:"while,omitempty"`
	Patterns   []*Rule     `json:"patterns,omitempty"`
	Captures   map[string]Capture `json:"captures,omitempty"`
	BeginCaptures map[string]Capture `json:"beginCaptures,omitempty"`
	EndCaptures   map[string]Capture `json:"endCaptures,omitempty"`
	ContentName string     `json:"contentName,omitempty"`
}

type Capture struct {
	Name  string `json:"name,omitempty"`
	Patterns []*Rule `json:"patterns,omitempty"`
}
