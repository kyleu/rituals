package rpermission

import (
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

const DefaultRoute = "/admin/db/retro/permission"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(paths...)
}

var _ svc.Model = (*RetroPermission)(nil)

type PK struct {
	RetroID uuid.UUID `json:"retroID,omitzero"`
	Key     string    `json:"key,omitzero"`
	Value   string    `json:"value,omitzero"`
}

func (p *PK) String() string {
	return fmt.Sprintf("%v • %s • %s", p.RetroID, p.Key, p.Value)
}

type RetroPermission struct {
	RetroID uuid.UUID `json:"retroID,omitzero"`
	Key     string    `json:"key,omitzero"`
	Value   string    `json:"value,omitzero"`
	Access  string    `json:"access,omitzero"`
	Created time.Time `json:"created,omitzero"`
}

func NewRetroPermission(retroID uuid.UUID, key string, value string) *RetroPermission {
	return &RetroPermission{RetroID: retroID, Key: key, Value: value}
}

func (r *RetroPermission) Clone() *RetroPermission {
	if r == nil {
		return nil
	}
	return &RetroPermission{RetroID: r.RetroID, Key: r.Key, Value: r.Value, Access: r.Access, Created: r.Created}
}

func (r *RetroPermission) String() string {
	return fmt.Sprintf("%s • %s • %s", r.RetroID.String(), r.Key, r.Value)
}

func (r *RetroPermission) TitleString() string {
	return r.String()
}

func (r *RetroPermission) ToPK() *PK {
	return &PK{
		RetroID: r.RetroID,
		Key:     r.Key,
		Value:   r.Value,
	}
}

func RandomRetroPermission() *RetroPermission {
	return &RetroPermission{
		RetroID: util.UUID(),
		Key:     util.RandomString(12),
		Value:   util.RandomString(12),
		Access:  util.RandomString(12),
		Created: util.TimeCurrent(),
	}
}

func (r *RetroPermission) Strings() []string {
	return []string{r.RetroID.String(), r.Key, r.Value, r.Access, util.TimeToFull(&r.Created)}
}

func (r *RetroPermission) ToCSV() ([]string, [][]string) {
	return RetroPermissionFieldDescs.Keys(), [][]string{r.Strings()}
}

func (r *RetroPermission) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(append(paths, url.QueryEscape(r.RetroID.String()), url.QueryEscape(r.Key), url.QueryEscape(r.Value))...)
}

func (r *RetroPermission) Breadcrumb(extra ...string) string {
	return r.TitleString() + "||" + r.WebPath(extra...) + "**permission"
}

func (r *RetroPermission) ToData() []any {
	return []any{r.RetroID, r.Key, r.Value, r.Access, r.Created}
}

var RetroPermissionFieldDescs = util.FieldDescs{
	{Key: "retroID", Title: "Retro ID", Type: "uuid"},
	{Key: "key", Title: "Key", Type: "string"},
	{Key: "value", Title: "Value", Type: "string"},
	{Key: "access", Title: "Access", Type: "string"},
	{Key: "created", Title: "Created", Type: "timestamp"},
}
