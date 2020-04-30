package util

import (
	"github.com/gofrs/uuid"
	"golang.org/x/text/language"
)

var AllColors = []string{"clear", "grey", "bluegrey", "red", "orange", "yellow", "green", "blue", "purple"}

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
	UserID    uuid.UUID
	Name      string
	Role      string
	Theme     Theme
	NavColor  string
	LinkColor string
	Locale    language.Tag
}

func (p *UserProfile) LinkClass() string {
	return p.LinkColor + "-fg"
}

var SystemProfile = NewUserProfile(uuid.UUID{})

func NewUserProfile(id uuid.UUID) UserProfile {
	return UserProfile{
		UserID:    id,
		Name:      "Guest",
		Role:      "user",
		Theme:     ThemeLight,
		NavColor:  "bluegrey",
		LinkColor: "bluegrey",
		Locale:    language.AmericanEnglish,
	}
}
