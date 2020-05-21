package util

import (
	"encoding/json"
)

type Theme struct {
	Name            string
	BackgroundClass string
	CardClass       string
	LogoPath        string
}

var ThemeLight = Theme{
	Name:            "light",
	BackgroundClass: "uk-dark",
	CardClass:       "uk-card-default",
	LogoPath:        "/assets/logo.png",
}

var ThemeDark = Theme{
	Name:            "dark",
	BackgroundClass: "uk-light",
	CardClass:       "uk-card-secondary",
	LogoPath:        "/assets/logo-white.png",
}

var AllThemes = []Theme{ThemeLight, ThemeDark}

func (t *Theme) String() string {
	return t.Name
}

func (t Theme) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Name)
}

func ThemeFromString(s string) Theme {
	for _, t := range AllThemes {
		if t.Name == s {
			return t
		}
	}

	return ThemeLight
}

var AllColors = []string{"clear", "grey", "bluegrey", "red", "orange", "yellow", "green", "blue", "purple"}
