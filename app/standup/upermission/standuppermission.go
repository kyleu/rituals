// Content managed by Project Forge, see [projectforge.md] for details.
package upermission

import (
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

type PK struct {
	StandupID uuid.UUID `json:"standupID"`
	Key       string    `json:"key"`
	Value     string    `json:"value"`
}

type StandupPermission struct {
	StandupID uuid.UUID `json:"standupID"`
	Key       string    `json:"key"`
	Value     string    `json:"value"`
	Access    string    `json:"access"`
	Created   time.Time `json:"created"`
}

func New(standupID uuid.UUID, key string, value string) *StandupPermission {
	return &StandupPermission{StandupID: standupID, Key: key, Value: value}
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

func FromMap(m util.ValueMap, setPK bool) (*StandupPermission, error) {
	ret := &StandupPermission{}
	var err error
	if setPK {
		retStandupID, e := m.ParseUUID("standupID", true, true)
		if e != nil {
			return nil, e
		}
		if retStandupID != nil {
			ret.StandupID = *retStandupID
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

func (s *StandupPermission) WebPath() string {
	return "/admin/db/standup/permission/" + s.StandupID.String() + "/" + url.QueryEscape(s.Key) + "/" + url.QueryEscape(s.Value)
}

func (s *StandupPermission) Diff(sx *StandupPermission) util.Diffs {
	var diffs util.Diffs
	if s.StandupID != sx.StandupID {
		diffs = append(diffs, util.NewDiff("standupID", s.StandupID.String(), sx.StandupID.String()))
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

func (s *StandupPermission) ToData() []any {
	return []any{s.StandupID, s.Key, s.Value, s.Access, s.Created}
}
