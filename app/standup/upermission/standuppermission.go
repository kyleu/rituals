package upermission

import (
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

const DefaultRoute = "/admin/db/standup/permission"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(paths...)
}

var _ svc.Model = (*StandupPermission)(nil)

type PK struct {
	StandupID uuid.UUID `json:"standupID,omitzero"`
	Key       string    `json:"key,omitzero"`
	Value     string    `json:"value,omitzero"`
}

func (p *PK) String() string {
	return fmt.Sprintf("%v • %s • %s", p.StandupID, p.Key, p.Value)
}

type StandupPermission struct {
	StandupID uuid.UUID `json:"standupID,omitzero"`
	Key       string    `json:"key,omitzero"`
	Value     string    `json:"value,omitzero"`
	Access    string    `json:"access,omitzero"`
	Created   time.Time `json:"created,omitzero"`
}

func NewStandupPermission(standupID uuid.UUID, key string, value string) *StandupPermission {
	return &StandupPermission{StandupID: standupID, Key: key, Value: value}
}

func (s *StandupPermission) Clone() *StandupPermission {
	return &StandupPermission{StandupID: s.StandupID, Key: s.Key, Value: s.Value, Access: s.Access, Created: s.Created}
}

func (s *StandupPermission) String() string {
	return fmt.Sprintf("%s • %s • %s", s.StandupID.String(), s.Key, s.Value)
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

func RandomStandupPermission() *StandupPermission {
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
	return StandupPermissionFieldDescs.Keys(), [][]string{s.Strings()}
}

func (s *StandupPermission) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(append(paths, url.QueryEscape(s.StandupID.String()), url.QueryEscape(s.Key), url.QueryEscape(s.Value))...)
}

func (s *StandupPermission) Breadcrumb(extra ...string) string {
	return s.TitleString() + "||" + s.WebPath(extra...) + "**permission"
}

func (s *StandupPermission) ToData() []any {
	return []any{s.StandupID, s.Key, s.Value, s.Access, s.Created}
}

var StandupPermissionFieldDescs = util.FieldDescs{
	{Key: "standupID", Title: "Standup ID", Description: "", Type: "uuid"},
	{Key: "key", Title: "Key", Description: "", Type: "string"},
	{Key: "value", Title: "Value", Description: "", Type: "string"},
	{Key: "access", Title: "Access", Description: "", Type: "string"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
}
