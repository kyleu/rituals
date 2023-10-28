// Package spermission - Content managed by Project Forge, see [projectforge.md] for details.
package spermission

import (
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

type PK struct {
	SprintID uuid.UUID `json:"sprintID,omitempty"`
	Key      string    `json:"key,omitempty"`
	Value    string    `json:"value,omitempty"`
}

type SprintPermission struct {
	SprintID uuid.UUID `json:"sprintID,omitempty"`
	Key      string    `json:"key,omitempty"`
	Value    string    `json:"value,omitempty"`
	Access   string    `json:"access,omitempty"`
	Created  time.Time `json:"created,omitempty"`
}

func New(sprintID uuid.UUID, key string, value string) *SprintPermission {
	return &SprintPermission{SprintID: sprintID, Key: key, Value: value}
}

func (s *SprintPermission) Clone() *SprintPermission {
	return &SprintPermission{s.SprintID, s.Key, s.Value, s.Access, s.Created}
}

func (s *SprintPermission) String() string {
	return fmt.Sprintf("%s::%s::%s", s.SprintID.String(), s.Key, s.Value)
}

func (s *SprintPermission) TitleString() string {
	return s.String()
}

func (s *SprintPermission) ToPK() *PK {
	return &PK{
		SprintID: s.SprintID,
		Key:      s.Key,
		Value:    s.Value,
	}
}

func Random() *SprintPermission {
	return &SprintPermission{
		SprintID: util.UUID(),
		Key:      util.RandomString(12),
		Value:    util.RandomString(12),
		Access:   util.RandomString(12),
		Created:  util.TimeCurrent(),
	}
}

func (s *SprintPermission) WebPath() string {
	return "/admin/db/sprint/permission/" + s.SprintID.String() + "/" + url.QueryEscape(s.Key) + "/" + url.QueryEscape(s.Value)
}

func (s *SprintPermission) ToData() []any {
	return []any{s.SprintID, s.Key, s.Value, s.Access, s.Created}
}
