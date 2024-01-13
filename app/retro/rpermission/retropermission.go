// Package rpermission - Content managed by Project Forge, see [projectforge.md] for details.
package rpermission

import (
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

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

func (r *RetroPermission) WebPath() string {
	return "/admin/db/retro/permission/" + r.RetroID.String() + "/" + url.QueryEscape(r.Key) + "/" + url.QueryEscape(r.Value)
}

func (r *RetroPermission) ToData() []any {
	return []any{r.RetroID, r.Key, r.Value, r.Access, r.Created}
}
