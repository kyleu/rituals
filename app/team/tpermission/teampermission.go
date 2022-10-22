// Content managed by Project Forge, see [projectforge.md] for details.
package tpermission

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/slices"

	"github.com/kyleu/rituals/app/util"
)

type PK struct {
	TeamID uuid.UUID `json:"teamID"`
	K      string    `json:"k"`
	V      string    `json:"v"`
}

type TeamPermission struct {
	TeamID  uuid.UUID `json:"teamID"`
	K       string    `json:"k"`
	V       string    `json:"v"`
	Access  string    `json:"access"`
	Created time.Time `json:"created"`
}

func New(teamID uuid.UUID, k string, v string) *TeamPermission {
	return &TeamPermission{TeamID: teamID, K: k, V: v}
}

func Random() *TeamPermission {
	return &TeamPermission{
		TeamID:  util.UUID(),
		K:       util.RandomString(12),
		V:       util.RandomString(12),
		Access:  util.RandomString(12),
		Created: time.Now(),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*TeamPermission, error) {
	ret := &TeamPermission{}
	var err error
	if setPK {
		retTeamID, e := m.ParseUUID("teamID", true, true)
		if e != nil {
			return nil, e
		}
		if retTeamID != nil {
			ret.TeamID = *retTeamID
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

func (t *TeamPermission) Clone() *TeamPermission {
	return &TeamPermission{
		TeamID:  t.TeamID,
		K:       t.K,
		V:       t.V,
		Access:  t.Access,
		Created: t.Created,
	}
}

func (t *TeamPermission) String() string {
	return fmt.Sprintf("%s::%s::%s", t.TeamID.String(), t.K, t.V)
}

func (t *TeamPermission) TitleString() string {
	return t.String()
}

func (t *TeamPermission) WebPath() string {
	return "/admin/db/team/permission" + "/" + t.TeamID.String() + "/" + t.K + "/" + t.V
}

func (t *TeamPermission) Diff(tx *TeamPermission) util.Diffs {
	var diffs util.Diffs
	if t.TeamID != tx.TeamID {
		diffs = append(diffs, util.NewDiff("teamID", t.TeamID.String(), tx.TeamID.String()))
	}
	if t.K != tx.K {
		diffs = append(diffs, util.NewDiff("k", t.K, tx.K))
	}
	if t.V != tx.V {
		diffs = append(diffs, util.NewDiff("v", t.V, tx.V))
	}
	if t.Access != tx.Access {
		diffs = append(diffs, util.NewDiff("access", t.Access, tx.Access))
	}
	if t.Created != tx.Created {
		diffs = append(diffs, util.NewDiff("created", t.Created.String(), tx.Created.String()))
	}
	return diffs
}

func (t *TeamPermission) ToData() []any {
	return []any{t.TeamID, t.K, t.V, t.Access, t.Created}
}

type TeamPermissions []*TeamPermission

func (t TeamPermissions) Get(teamID uuid.UUID, k string, v string) *TeamPermission {
	for _, x := range t {
		if x.TeamID == teamID && x.K == k && x.V == v {
			return x
		}
	}
	return nil
}

func (t TeamPermissions) Clone() TeamPermissions {
	return slices.Clone(t)
}
