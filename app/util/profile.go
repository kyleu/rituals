package util

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"golang.org/x/text/language"
)

type Role struct {
	Key string
}

var RoleGuest = Role{
	Key: "guest",
}

var RoleUser = Role{
	Key: "user",
}

var RoleAdmin = Role{
	Key: "admin",
}

var AllRoles = []Role{RoleGuest, RoleUser, RoleAdmin}

func (t *Role) String() string {
	return t.Key
}

func (t Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Key)
}

func RoleFromString(s string) Role {
	for _, t := range AllRoles {
		if t.Key == s {
			return t
		}
	}

	return RoleGuest
}

type UserProfile struct {
	UserID    uuid.UUID
	Name      string
	Role      Role
	Theme     Theme
	NavColor  string
	LinkColor string
	Picture   string
	Locale    language.Tag
}

func (p *UserProfile) LinkClass() string {
	return p.LinkColor + "-fg"
}

func NewUserProfile(id uuid.UUID) *UserProfile {
	return &UserProfile{
		UserID:    id,
		Name:      "Guest",
		Role:      RoleGuest,
		Theme:     ThemeLight,
		NavColor:  "bluegrey",
		LinkColor: "bluegrey",
		Picture:   "",
		Locale:    language.AmericanEnglish,
	}
}

type Profile struct {
	UserID    uuid.UUID `json:"userID"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	Theme     string    `json:"theme"`
	NavColor  string    `json:"navColor"`
	LinkColor string    `json:"linkColor"`
	Locale    string    `json:"locale"`
}

func (p *UserProfile) ToProfile() Profile {
	return Profile{
		UserID:    p.UserID,
		Name:      p.Name,
		Role:      p.Role.String(),
		Theme:     p.Theme.String(),
		NavColor:  p.NavColor,
		LinkColor: p.LinkColor,
		Locale:    p.Locale.String(),
	}
}
