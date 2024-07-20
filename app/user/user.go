package user

import (
	"net/url"
	"path"
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
	return path.Join(paths...)
}

var _ svc.Model = (*User)(nil)

type User struct {
	ID      uuid.UUID  `json:"id,omitempty"`
	Name    string     `json:"name,omitempty"`
	Picture string     `json:"picture,omitempty"`
	Created time.Time  `json:"created,omitempty"`
	Updated *time.Time `json:"updated,omitempty"`
}

func New(id uuid.UUID) *User {
	return &User{ID: id}
}

func (u *User) Clone() *User {
	return &User{u.ID, u.Name, u.Picture, u.Created, u.Updated}
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

func Random() *User {
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
	return FieldDescs.Keys(), [][]string{u.Strings()}
}

func (u *User) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return path.Join(append(paths, url.QueryEscape(u.ID.String()))...)
}

func (u *User) ToData() []any {
	return []any{u.ID, u.Name, u.Picture, u.Created, u.Updated}
}

var FieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Description: "", Type: "uuid"},
	{Key: "name", Title: "Name", Description: "", Type: "string"},
	{Key: "picture", Title: "Picture", Description: "", Type: "string"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
	{Key: "updated", Title: "Updated", Description: "", Type: "timestamp"},
}
