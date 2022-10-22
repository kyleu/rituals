// Content managed by Project Forge, see [projectforge.md] for details.
package upermission

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/slices"

	"github.com/kyleu/rituals/app/util"
)

type PK struct {
	StandupID uuid.UUID `json:"standupID"`
	K         string    `json:"k"`
	V         string    `json:"v"`
}

type StandupPermission struct {
	StandupID uuid.UUID `json:"standupID"`
	K         string    `json:"k"`
	V         string    `json:"v"`
	Access    string    `json:"access"`
	Created   time.Time `json:"created"`
}

func New(standupID uuid.UUID, k string, v string) *StandupPermission {
	return &StandupPermission{StandupID: standupID, K: k, V: v}
}

func Random() *StandupPermission {
	return &StandupPermission{
		StandupID: util.UUID(),
		K:         util.RandomString(12),
		V:         util.RandomString(12),
		Access:    util.RandomString(12),
		Created:   time.Now(),
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
		ret.K, err = m.ParseString("k", true, true)
		if err != nil {
			return nil, err
		}
		ret.V, err = m.ParseString("v", true, true)
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
	return &StandupPermission{
		StandupID: s.StandupID,
		K:         s.K,
		V:         s.V,
		Access:    s.Access,
		Created:   s.Created,
	}
}

func (s *StandupPermission) String() string {
	return fmt.Sprintf("%s::%s::%s", s.StandupID.String(), s.K, s.V)
}

func (s *StandupPermission) TitleString() string {
	return s.String()
}

func (s *StandupPermission) WebPath() string {
	return "/admin/db/standup/permission" + "/" + s.StandupID.String() + "/" + s.K + "/" + s.V
}

func (s *StandupPermission) Diff(sx *StandupPermission) util.Diffs {
	var diffs util.Diffs
	if s.StandupID != sx.StandupID {
		diffs = append(diffs, util.NewDiff("standupID", s.StandupID.String(), sx.StandupID.String()))
	}
	if s.K != sx.K {
		diffs = append(diffs, util.NewDiff("k", s.K, sx.K))
	}
	if s.V != sx.V {
		diffs = append(diffs, util.NewDiff("v", s.V, sx.V))
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
	return []any{s.StandupID, s.K, s.V, s.Access, s.Created}
}

type StandupPermissions []*StandupPermission

func (s StandupPermissions) Get(standupID uuid.UUID, k string, v string) *StandupPermission {
	for _, x := range s {
		if x.StandupID == standupID && x.K == k && x.V == v {
			return x
		}
	}
	return nil
}

func (s StandupPermissions) Clone() StandupPermissions {
	return slices.Clone(s)
}
