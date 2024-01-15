// Package user - Content managed by Project Forge, see [projectforge.md] for details.
package user

import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

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
	return u.Name
}

func Random() *User {
	return &User{
		ID:      util.UUID(),
		Name:    util.RandomString(12),
		Picture: "https://" + util.RandomString(6) + ".com/" + util.RandomString(6),
		Created: util.TimeCurrent(),
		Updated: util.TimeCurrentP(),
	}
}

func (u *User) WebPath() string {
	return "/admin/db/user/" + u.ID.String()
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
