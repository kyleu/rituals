package util

import "golang.org/x/text/language"

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

func (t Theme) String() string {
	return t.Name
}

func ThemeFromString(s string) Theme {
	for _, t := range AllThemes {
		if t.String() == s {
			return t
		}
	}
	return ThemeLight
}

type UserProfile struct {
	Name      string
	Theme     Theme
	NavColor  string
	LinkColor string
	Locale    language.Tag
}

func (p *UserProfile) LinkClass() string {
	return p.LinkColor + "-fg"
}

var SystemProfile = NewUserProfile()

func NewUserProfile() UserProfile {
	return UserProfile{
		Name:      "System",
		Theme:     ThemeLight,
		NavColor:  "bluegrey",
		LinkColor: "bluegrey",
		Locale:    language.AmericanEnglish,
	}
}
