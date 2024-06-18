package rpermission

import (
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

var _ svc.Model = (*RetroPermission)(nil)

type PK struct {
	RetroID uuid.UUID `json:"retroID,omitempty"`
	Key     string    `json:"key,omitempty"`
	Value   string    `json:"value,omitempty"`
}

func (p *PK) String() string {
	return fmt.Sprintf("%v::%s::%s", p.RetroID, p.Key, p.Value)
}

type RetroPermission struct {
	RetroID uuid.UUID `json:"retroID,omitempty"`
	Key     string    `json:"key,omitempty"`
	Value   string    `json:"value,omitempty"`
	Access  string    `json:"access,omitempty"`
	Created time.Time `json:"created,omitempty"`
}

func New(retroID uuid.UUID, key string, value string) *RetroPermission {
	return &RetroPermission{RetroID: retroID, Key: key, Value: value}
}

func (r *RetroPermission) Clone() *RetroPermission {
	return &RetroPermission{r.RetroID, r.Key, r.Value, r.Access, r.Created}
}

func (r *RetroPermission) String() string {
	return fmt.Sprintf("%s::%s::%s", r.RetroID.String(), r.Key, r.Value)
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

func Random() *RetroPermission {
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
	return FieldDescs.Keys(), [][]string{r.Strings()}
}

func (r *RetroPermission) WebPath() string {
	return "/admin/db/retro/permission/" + r.RetroID.String() + "/" + url.QueryEscape(r.Key) + "/" + url.QueryEscape(r.Value)
}

func (r *RetroPermission) ToData() []any {
	return []any{r.RetroID, r.Key, r.Value, r.Access, r.Created}
}

var FieldDescs = util.FieldDescs{
	{Key: "retroID", Title: "Retro ID", Description: "", Type: "uuid"},
	{Key: "key", Title: "Key", Description: "", Type: "string"},
	{Key: "value", Title: "Value", Description: "", Type: "string"},
	{Key: "access", Title: "Access", Description: "", Type: "string"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
}
