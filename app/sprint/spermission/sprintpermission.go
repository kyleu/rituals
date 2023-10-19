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

func Random() *SprintPermission {
	return &SprintPermission{
		SprintID: util.UUID(),
		Key:      util.RandomString(12),
		Value:    util.RandomString(12),
		Access:   util.RandomString(12),
		Created:  util.TimeCurrent(),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*SprintPermission, error) {
	ret := &SprintPermission{}
	var err error
	if setPK {
		retSprintID, e := m.ParseUUID("sprintID", true, true)
		if e != nil {
			return nil, e
		}
		if retSprintID != nil {
			ret.SprintID = *retSprintID
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

func (s *SprintPermission) WebPath() string {
	return "/admin/db/sprint/permission/" + s.SprintID.String() + "/" + url.QueryEscape(s.Key) + "/" + url.QueryEscape(s.Value)
}

func (s *SprintPermission) Diff(sx *SprintPermission) util.Diffs {
	var diffs util.Diffs
	if s.SprintID != sx.SprintID {
		diffs = append(diffs, util.NewDiff("sprintID", s.SprintID.String(), sx.SprintID.String()))
	}
	if s.Key != sx.Key {
		diffs = append(diffs, util.NewDiff("key", s.Key, sx.Key))
	}
	if s.Value != sx.Value {
		diffs = append(diffs, util.NewDiff("value", s.Value, sx.Value))
	}
	if s.Access != sx.Access {
		diffs = append(diffs, util.NewDiff("access", s.Access, sx.Access))
	}
	if s.Created != sx.Created {
		diffs = append(diffs, util.NewDiff("created", s.Created.String(), sx.Created.String()))
	}
	return diffs
}

func (s *SprintPermission) ToData() []any {
	return []any{s.SprintID, s.Key, s.Value, s.Access, s.Created}
}
