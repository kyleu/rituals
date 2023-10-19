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

func Random() *RetroPermission {
	return &RetroPermission{
		RetroID: util.UUID(),
		Key:     util.RandomString(12),
		Value:   util.RandomString(12),
		Access:  util.RandomString(12),
		Created: util.TimeCurrent(),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*RetroPermission, error) {
	ret := &RetroPermission{}
	var err error
	if setPK {
		retRetroID, e := m.ParseUUID("retroID", true, true)
		if e != nil {
			return nil, e
		}
		if retRetroID != nil {
			ret.RetroID = *retRetroID
		}
		ret.Key, err = m.ParseString("key", true, true)
		if err != nil {
			return nil, err
		}
		ret.Value, err = m.ParseString("value", true, true)
		if err != nil {
			return nil, err
		}
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
	}
	ret.Access, err = m.ParseString("access", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
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

func (r *RetroPermission) WebPath() string {
	return "/admin/db/retro/permission/" + r.RetroID.String() + "/" + url.QueryEscape(r.Key) + "/" + url.QueryEscape(r.Value)
}

func (r *RetroPermission) Diff(rx *RetroPermission) util.Diffs {
	var diffs util.Diffs
	if r.RetroID != rx.RetroID {
		diffs = append(diffs, util.NewDiff("retroID", r.RetroID.String(), rx.RetroID.String()))
	}
	if r.Key != rx.Key {
		diffs = append(diffs, util.NewDiff("key", r.Key, rx.Key))
	}
	if r.Value != rx.Value {
		diffs = append(diffs, util.NewDiff("value", r.Value, rx.Value))
	}
	if r.Access != rx.Access {
		diffs = append(diffs, util.NewDiff("access", r.Access, rx.Access))
	}
	if r.Created != rx.Created {
		diffs = append(diffs, util.NewDiff("created", r.Created.String(), rx.Created.String()))
	}
	return diffs
}

func (r *RetroPermission) ToData() []any {
	return []any{r.RetroID, r.Key, r.Value, r.Access, r.Created}
}
