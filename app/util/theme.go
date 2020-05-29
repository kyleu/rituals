package util

import (
	"encoding/json"
)

type Theme struct {
	Name string
}

var ThemeDefault = Theme{
	Name: "default",
}

var ThemeLight = Theme{
	Name: "light",
}

var ThemeDark = Theme{
	Name: "dark",
}

var AllThemes = []Theme{ThemeDefault, ThemeLight, ThemeDark}

func ThemeFromString(s string) Theme {
	for _, t := range AllThemes {
		if t.Name == s {
			return t
		}
	}
	return ThemeDefault
}

func (t *Theme) String() string {
	return t.Name
}

func (t Theme) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Name)
}

func (t *Theme) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*t = ThemeFromString(s)
	return nil
}

var AllColors = []string{"clear", "grey", "bluegrey", "red", "orange", "yellow", "green", "blue", "purple"}
