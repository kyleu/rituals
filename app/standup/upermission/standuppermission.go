package upermission

import (
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

var _ svc.Model = (*StandupPermission)(nil)

type PK struct {
	StandupID uuid.UUID `json:"standupID,omitempty"`
	Key       string    `json:"key,omitempty"`
	Value     string    `json:"value,omitempty"`
}

func (p *PK) String() string {
	return fmt.Sprintf("%v::%s::%s", p.StandupID, p.Key, p.Value)
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

func (s *StandupPermission) Strings() []string {
	return []string{s.StandupID.String(), s.Key, s.Value, s.Access, util.TimeToFull(&s.Created)}
}

func (s *StandupPermission) ToCSV() ([]string, [][]string) {
	return FieldDescs.Keys(), [][]string{s.Strings()}
}

func (s *StandupPermission) WebPath() string {
	return "/admin/db/standup/permission/" + s.StandupID.String() + "/" + url.QueryEscape(s.Key) + "/" + url.QueryEscape(s.Value)
}

func (s *StandupPermission) ToData() []any {
	return []any{s.StandupID, s.Key, s.Value, s.Access, s.Created}
}

var FieldDescs = util.FieldDescs{
	{Key: "standupID", Title: "Standup ID", Description: "", Type: "uuid"},
	{Key: "key", Title: "Key", Description: "", Type: "string"},
	{Key: "value", Title: "Value", Description: "", Type: "string"},
	{Key: "access", Title: "Access", Description: "", Type: "string"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
}
