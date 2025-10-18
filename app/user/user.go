package user

import (
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

const DefaultRoute = "/admin/db/user"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(paths...)
}

var _ svc.Model = (*User)(nil)

type User struct {
	ID      uuid.UUID  `json:"id,omitzero"`
	Name    string     `json:"name,omitzero"`
	Picture string     `json:"picture,omitzero"`
	Created time.Time  `json:"created,omitzero"`
	Updated *time.Time `json:"updated,omitzero"`
}

func NewUser(id uuid.UUID) *User {
	return &User{ID: id}
}

func (u *User) Clone() *User {
	return &User{ID: u.ID, Name: u.Name, Picture: u.Picture, Created: u.Created, Updated: u.Updated}
}

func (u *User) String() string {
	return u.ID.String()
}

func (u *User) TitleString() string {
	if xx := u.Name; xx != "" {
		return xx
	}
	return u.String()
}

func RandomUser() *User {
	return &User{
		ID:      util.UUID(),
		Name:    util.RandomString(12),
		Picture: util.RandomURL().String(),
		Created: util.TimeCurrent(),
		Updated: util.TimeCurrentP(),
	}
}

func (u *User) Strings() []string {
	return []string{u.ID.String(), u.Name, u.Picture, util.TimeToFull(&u.Created), util.TimeToFull(u.Updated)}
}

func (u *User) ToCSV() ([]string, [][]string) {
	return UserFieldDescs.Keys(), [][]string{u.Strings()}
}

func (u *User) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(append(paths, url.QueryEscape(u.ID.String()))...)
}

func (u *User) Breadcrumb(extra ...string) string {
	return u.TitleString() + "||" + u.WebPath(extra...) + "**profile"
}

func (u *User) ToData() []any {
	return []any{u.ID, u.Name, u.Picture, u.Created, u.Updated}
}

var UserFieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Type: "uuid"},
	{Key: "name", Title: "Name", Type: "string"},
	{Key: "picture", Title: "Picture", Type: "string"},
	{Key: "created", Title: "Created", Type: "timestamp"},
	{Key: "updated", Title: "Updated", Type: "timestamp"},
}
