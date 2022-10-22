// Content managed by Project Forge, see [projectforge.md] for details.
package rpermission

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/slices"

	"github.com/kyleu/rituals/app/util"
)

type PK struct {
	RetroID uuid.UUID `json:"retroID"`
	K       string    `json:"k"`
	V       string    `json:"v"`
}

type RetroPermission struct {
	RetroID uuid.UUID `json:"retroID"`
	K       string    `json:"k"`
	V       string    `json:"v"`
	Access  string    `json:"access"`
	Created time.Time `json:"created"`
}

func New(retroID uuid.UUID, k string, v string) *RetroPermission {
	return &RetroPermission{RetroID: retroID, K: k, V: v}
}

func Random() *RetroPermission {
	return &RetroPermission{
		RetroID: util.UUID(),
		K:       util.RandomString(12),
		V:       util.RandomString(12),
		Access:  util.RandomString(12),
		Created: time.Now(),
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

func (r *RetroPermission) Clone() *RetroPermission {
	return &RetroPermission{
		RetroID: r.RetroID,
		K:       r.K,
		V:       r.V,
		Access:  r.Access,
		Created: r.Created,
	}
}

func (r *RetroPermission) String() string {
	return fmt.Sprintf("%s::%s::%s", r.RetroID.String(), r.K, r.V)
}

func (r *RetroPermission) TitleString() string {
	return r.String()
}

func (r *RetroPermission) WebPath() string {
	return "/retro/rpermission" + "/" + r.RetroID.String() + "/" + r.K + "/" + r.V
}

func (r *RetroPermission) Diff(rx *RetroPermission) util.Diffs {
	var diffs util.Diffs
	if r.RetroID != rx.RetroID {
		diffs = append(diffs, util.NewDiff("retroID", r.RetroID.String(), rx.RetroID.String()))
	}
	if r.K != rx.K {
		diffs = append(diffs, util.NewDiff("k", r.K, rx.K))
	}
	if r.V != rx.V {
		diffs = append(diffs, util.NewDiff("v", r.V, rx.V))
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
	return []any{r.RetroID, r.K, r.V, r.Access, r.Created}
}

type RetroPermissions []*RetroPermission

func (r RetroPermissions) Get(retroID uuid.UUID, k string, v string) *RetroPermission {
	for _, x := range r {
		if x.RetroID == retroID && x.K == k && x.V == v {
			return x
		}
	}
	return nil
}

func (r RetroPermissions) Clone() RetroPermissions {
	return slices.Clone(r)
}
