package tpermission

import (
	"fmt"
	"net/url"
	"path"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

const DefaultRoute = "/admin/db/team/permission"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return path.Join(paths...)
}

var _ svc.Model = (*TeamPermission)(nil)

type PK struct {
	TeamID uuid.UUID `json:"teamID,omitempty"`
	Key    string    `json:"key,omitempty"`
	Value  string    `json:"value,omitempty"`
}

func (p *PK) String() string {
	return fmt.Sprintf("%v • %s • %s", p.TeamID, p.Key, p.Value)
}

type TeamPermission struct {
	TeamID  uuid.UUID `json:"teamID,omitempty"`
	Key     string    `json:"key,omitempty"`
	Value   string    `json:"value,omitempty"`
	Access  string    `json:"access,omitempty"`
	Created time.Time `json:"created,omitempty"`
}

func NewTeamPermission(teamID uuid.UUID, key string, value string) *TeamPermission {
	return &TeamPermission{TeamID: teamID, Key: key, Value: value}
}

func (t *TeamPermission) Clone() *TeamPermission {
	return &TeamPermission{t.TeamID, t.Key, t.Value, t.Access, t.Created}
}

func (t *TeamPermission) String() string {
	return fmt.Sprintf("%s • %s • %s", t.TeamID.String(), t.Key, t.Value)
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

func RandomTeamPermission() *TeamPermission {
	return &TeamPermission{
		TeamID:  util.UUID(),
		Key:     util.RandomString(12),
		Value:   util.RandomString(12),
		Access:  util.RandomString(12),
		Created: util.TimeCurrent(),
	}
}

func (t *TeamPermission) Strings() []string {
	return []string{t.TeamID.String(), t.Key, t.Value, t.Access, util.TimeToFull(&t.Created)}
}

func (t *TeamPermission) ToCSV() ([]string, [][]string) {
	return TeamPermissionFieldDescs.Keys(), [][]string{t.Strings()}
}

func (t *TeamPermission) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return path.Join(append(paths, url.QueryEscape(t.TeamID.String()), url.QueryEscape(t.Key), url.QueryEscape(t.Value))...)
}

func (t *TeamPermission) Breadcrumb(extra ...string) string {
	return t.TitleString() + "||" + t.WebPath(extra...) + "**permission"
}

func (t *TeamPermission) ToData() []any {
	return []any{t.TeamID, t.Key, t.Value, t.Access, t.Created}
}

var TeamPermissionFieldDescs = util.FieldDescs{
	{Key: "teamID", Title: "Team ID", Description: "", Type: "uuid"},
	{Key: "key", Title: "Key", Description: "", Type: "string"},
	{Key: "value", Title: "Value", Description: "", Type: "string"},
	{Key: "access", Title: "Access", Description: "", Type: "string"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
}
