// Content managed by Project Forge, see [projectforge.md] for details.
package spermission

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

type PK struct {
	SprintID uuid.UUID `json:"sprintID"`
	K        string    `json:"k"`
	V        string    `json:"v"`
}

type SprintPermission struct {
	SprintID uuid.UUID `json:"sprintID"`
	K        string    `json:"k"`
	V        string    `json:"v"`
	Access   string    `json:"access"`
	Created  time.Time `json:"created"`
}

func New(sprintID uuid.UUID, k string, v string) *SprintPermission {
	return &SprintPermission{SprintID: sprintID, K: k, V: v}
}

func Random() *SprintPermission {
	return &SprintPermission{
		SprintID: util.UUID(),
		K:        util.RandomString(12),
		V:        util.RandomString(12),
		Access:   util.RandomString(12),
		Created:  time.Now(),
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

func (s *SprintPermission) Clone() *SprintPermission {
	return &SprintPermission{
		SprintID: s.SprintID,
		K:        s.K,
		V:        s.V,
		Access:   s.Access,
		Created:  s.Created,
	}
}

func (s *SprintPermission) String() string {
	return fmt.Sprintf("%s::%s::%s", s.SprintID.String(), s.K, s.V)
}

func (s *SprintPermission) TitleString() string {
	return s.String()
}

func (s *SprintPermission) WebPath() string {
	return "/admin/db/sprint/permission" + "/" + s.SprintID.String() + "/" + s.K + "/" + s.V
}

func (s *SprintPermission) Diff(sx *SprintPermission) util.Diffs {
	var diffs util.Diffs
	if s.SprintID != sx.SprintID {
		diffs = append(diffs, util.NewDiff("sprintID", s.SprintID.String(), sx.SprintID.String()))
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

func (s *SprintPermission) ToData() []any {
	return []any{s.SprintID, s.K, s.V, s.Access, s.Created}
}
