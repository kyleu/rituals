// Package tpermission - Content managed by Project Forge, see [projectforge.md] for details.
package tpermission

import (
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

type PK struct {
	TeamID uuid.UUID `json:"teamID,omitempty"`
	Key    string    `json:"key,omitempty"`
	Value  string    `json:"value,omitempty"`
}

type TeamPermission struct {
	TeamID  uuid.UUID `json:"teamID,omitempty"`
	Key     string    `json:"key,omitempty"`
	Value   string    `json:"value,omitempty"`
	Access  string    `json:"access,omitempty"`
	Created time.Time `json:"created,omitempty"`
}

func New(teamID uuid.UUID, key string, value string) *TeamPermission {
	return &TeamPermission{TeamID: teamID, Key: key, Value: value}
}

func (t *TeamPermission) Clone() *TeamPermission {
	return &TeamPermission{t.TeamID, t.Key, t.Value, t.Access, t.Created}
}

func (t *TeamPermission) String() string {
	return fmt.Sprintf("%s::%s::%s", t.TeamID.String(), t.Key, t.Value)
}

func (t *TeamPermission) TitleString() string {
	return t.String()
}

func (t *TeamPermission) ToPK() *PK {
	return &PK{
		TeamID: t.TeamID,
		Key:    t.Key,
		Value:  t.Value,
	}
}

func Random() *TeamPermission {
	return &TeamPermission{
		TeamID:  util.UUID(),
		Key:     util.RandomString(12),
		Value:   util.RandomString(12),
		Access:  util.RandomString(12),
		Created: util.TimeCurrent(),
	}
}

func (t *TeamPermission) WebPath() string {
	return "/admin/db/team/permission/" + t.TeamID.String() + "/" + url.QueryEscape(t.Key) + "/" + url.QueryEscape(t.Value)
}

func (t *TeamPermission) ToData() []any {
	return []any{t.TeamID, t.Key, t.Value, t.Access, t.Created}
}
