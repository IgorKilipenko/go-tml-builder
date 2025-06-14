package models

import (
	"encoding/json"

	"github.com/IgorKilipenko/go-tml-builder/internal/core/regexputil"
)

type Regexp struct {
	pattern string
}

// NewRegexp создает новое Regexp с валидацией
func NewRegexp(pattern string) (Regexp, error) {
	if err := regexputil.Validate(pattern); err != nil {
		return Regexp{}, err
	}
	return Regexp{pattern: pattern}, nil
}

// MustRegexp создает Regexp без ошибки (panic при невалидном паттерне)
func MustRegexp(pattern string) Regexp {
	r, err := NewRegexp(pattern)
	if err != nil {
		panic(err)
	}
	return r
}

func (r Regexp) String() string {
	return r.pattern
}

func (r Regexp) MarshalJSON() ([]byte, error) {
	return json.Marshal(regexputil.EscapeForJSON(r.pattern))
}

func (r *Regexp) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	r.pattern = regexputil.UnescapeFromJSON(s)
	return nil
}
