package spermission

import (
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

const DefaultRoute = "/admin/db/sprint/permission"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(paths...)
}

var _ svc.Model = (*SprintPermission)(nil)

type PK struct {
	SprintID uuid.UUID `json:"sprintID,omitempty"`
	Key      string    `json:"key,omitempty"`
	Value    string    `json:"value,omitempty"`
}

func (p *PK) String() string {
	return fmt.Sprintf("%v • %s • %s", p.SprintID, p.Key, p.Value)
}

type SprintPermission struct {
	SprintID uuid.UUID `json:"sprintID,omitempty"`
	Key      string    `json:"key,omitempty"`
	Value    string    `json:"value,omitempty"`
	Access   string    `json:"access,omitempty"`
	Created  time.Time `json:"created,omitempty"`
}

func NewSprintPermission(sprintID uuid.UUID, key string, value string) *SprintPermission {
	return &SprintPermission{SprintID: sprintID, Key: key, Value: value}
}

func (s *SprintPermission) Clone() *SprintPermission {
	return &SprintPermission{s.SprintID, s.Key, s.Value, s.Access, s.Created}
}

func (s *SprintPermission) String() string {
	return fmt.Sprintf("%s • %s • %s", s.SprintID.String(), s.Key, s.Value)
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

func RandomSprintPermission() *SprintPermission {
	return &SprintPermission{
		SprintID: util.UUID(),
		Key:      util.RandomString(12),
		Value:    util.RandomString(12),
		Access:   util.RandomString(12),
		Created:  util.TimeCurrent(),
	}
}

func (s *SprintPermission) Strings() []string {
	return []string{s.SprintID.String(), s.Key, s.Value, s.Access, util.TimeToFull(&s.Created)}
}

func (s *SprintPermission) ToCSV() ([]string, [][]string) {
	return SprintPermissionFieldDescs.Keys(), [][]string{s.Strings()}
}

func (s *SprintPermission) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(append(paths, url.QueryEscape(s.SprintID.String()), url.QueryEscape(s.Key), url.QueryEscape(s.Value))...)
}

func (s *SprintPermission) Breadcrumb(extra ...string) string {
	return s.TitleString() + "||" + s.WebPath(extra...) + "**permission"
}

func (s *SprintPermission) ToData() []any {
	return []any{s.SprintID, s.Key, s.Value, s.Access, s.Created}
}

var SprintPermissionFieldDescs = util.FieldDescs{
	{Key: "sprintID", Title: "Sprint ID", Description: "", Type: "uuid"},
	{Key: "key", Title: "Key", Description: "", Type: "string"},
	{Key: "value", Title: "Value", Description: "", Type: "string"},
	{Key: "access", Title: "Access", Description: "", Type: "string"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
}
