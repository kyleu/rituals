// Package upermission - Content managed by Project Forge, see [projectforge.md] for details.
package upermission

import (
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

type PK struct {
	StandupID uuid.UUID `json:"standupID,omitempty"`
	Key       string    `json:"key,omitempty"`
	Value     string    `json:"value,omitempty"`
}

type StandupPermission struct {
	StandupID uuid.UUID `json:"standupID,omitempty"`
	Key       string    `json:"key,omitempty"`
	Value     string    `json:"value,omitempty"`
	Access    string    `json:"access,omitempty"`
	Created   time.Time `json:"created,omitempty"`
}

func New(standupID uuid.UUID, key string, value string) *StandupPermission {
	return &StandupPermission{StandupID: standupID, Key: key, Value: value}
}

func (s *StandupPermission) Clone() *StandupPermission {
	return &StandupPermission{s.StandupID, s.Key, s.Value, s.Access, s.Created}
}

func (s *StandupPermission) String() string {
	return fmt.Sprintf("%s::%s::%s", s.StandupID.String(), s.Key, s.Value)
}

func (s *StandupPermission) TitleString() string {
	return s.String()
}

func (s *StandupPermission) ToPK() *PK {
	return &PK{
		StandupID: s.StandupID,
		Key:       s.Key,
		Value:     s.Value,
	}
}

func Random() *StandupPermission {
	return &StandupPermission{
		StandupID: util.UUID(),
		Key:       util.RandomString(12),
		Value:     util.RandomString(12),
		Access:    util.RandomString(12),
		Created:   util.TimeCurrent(),
	}
}

func (s *StandupPermission) WebPath() string {
	return "/admin/db/standup/permission/" + s.StandupID.String() + "/" + url.QueryEscape(s.Key) + "/" + url.QueryEscape(s.Value)
}

func (s *StandupPermission) ToData() []any {
	return []any{s.StandupID, s.Key, s.Value, s.Access, s.Created}
}
