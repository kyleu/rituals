package user

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/util"
	"golang.org/x/text/language"
)

type SystemUser struct {
	UserID    uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	Role      string    `db:"role"`
	Theme     string    `db:"theme"`
	NavColor  string    `db:"nav_color"`
	LinkColor string    `db:"link_color"`
	Locale    string    `db:"locale"`
	Created   time.Time `db:"created"`
}

func (su *SystemUser) ToProfile() util.UserProfile {
	locale, err := language.Parse(su.Locale)
	if err != nil {
		locale = language.AmericanEnglish
	}
	return util.UserProfile{
		UserID:    su.UserID,
		Name:      su.Name,
		Role:      util.RoleFromString(su.Role),
		Theme:     util.ThemeFromString(su.Theme),
		NavColor:  su.NavColor,
		LinkColor: su.LinkColor,
		Locale:    locale,
	}
}
