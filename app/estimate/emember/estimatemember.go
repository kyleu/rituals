// Content managed by Project Forge, see [projectforge.md] for details.
package emember

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/slices"

	"github.com/kyleu/rituals/app/util"
)

type PK struct {
	EstimateID uuid.UUID `json:"estimateID"`
	UserID     uuid.UUID `json:"userID"`
}

type EstimateMember struct {
	EstimateID uuid.UUID  `json:"estimateID"`
	UserID     uuid.UUID  `json:"userID"`
	Name       string     `json:"name"`
	Picture    string     `json:"picture"`
	Role       string     `json:"role"`
	Created    time.Time  `json:"created"`
	Updated    *time.Time `json:"updated,omitempty"`
}

func New(estimateID uuid.UUID, userID uuid.UUID) *EstimateMember {
	return &EstimateMember{EstimateID: estimateID, UserID: userID}
}

func Random() *EstimateMember {
	return &EstimateMember{
		EstimateID: util.UUID(),
		UserID:     util.UUID(),
		Name:       util.RandomString(12),
		Picture:    "https://" + util.RandomString(6) + ".com/" + util.RandomString(6),
		Role:       util.RandomString(12),
		Created:    time.Now(),
		Updated:    util.NowPointer(),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*EstimateMember, error) {
	ret := &EstimateMember{}
	var err error
	if setPK {
		retEstimateID, e := m.ParseUUID("estimateID", true, true)
		if e != nil {
			return nil, e
		}
		if retEstimateID != nil {
			ret.EstimateID = *retEstimateID
		}
		retUserID, e := m.ParseUUID("userID", true, true)
		if e != nil {
			return nil, e
		}
		if retUserID != nil {
			ret.UserID = *retUserID
		}
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
	}
	ret.Name, err = m.ParseString("name", true, true)
	if err != nil {
		return nil, err
	}
	ret.Picture, err = m.ParseString("picture", true, true)
	if err != nil {
		return nil, err
	}
	ret.Role, err = m.ParseString("role", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (e *EstimateMember) Clone() *EstimateMember {
	return &EstimateMember{
		EstimateID: e.EstimateID,
		UserID:     e.UserID,
		Name:       e.Name,
		Picture:    e.Picture,
		Role:       e.Role,
		Created:    e.Created,
		Updated:    e.Updated,
	}
}

func (e *EstimateMember) String() string {
	return fmt.Sprintf("%s::%s", e.EstimateID.String(), e.UserID.String())
}

func (e *EstimateMember) TitleString() string {
	return e.EstimateID.String() + " / " + e.Name
}

func (e *EstimateMember) WebPath() string {
	return "/admin/db/estimate/member" + "/" + e.EstimateID.String() + "/" + e.UserID.String()
}

func (e *EstimateMember) Diff(ex *EstimateMember) util.Diffs {
	var diffs util.Diffs
	if e.EstimateID != ex.EstimateID {
		diffs = append(diffs, util.NewDiff("estimateID", e.EstimateID.String(), ex.EstimateID.String()))
	}
	if e.UserID != ex.UserID {
		diffs = append(diffs, util.NewDiff("userID", e.UserID.String(), ex.UserID.String()))
	}
	if e.Name != ex.Name {
		diffs = append(diffs, util.NewDiff("name", e.Name, ex.Name))
	}
	if e.Picture != ex.Picture {
		diffs = append(diffs, util.NewDiff("picture", e.Picture, ex.Picture))
	}
	if e.Role != ex.Role {
		diffs = append(diffs, util.NewDiff("role", e.Role, ex.Role))
	}
	if e.Created != ex.Created {
		diffs = append(diffs, util.NewDiff("created", e.Created.String(), ex.Created.String()))
	}
	return diffs
}

func (e *EstimateMember) ToData() []any {
	return []any{e.EstimateID, e.UserID, e.Name, e.Picture, e.Role, e.Created, e.Updated}
}

type EstimateMembers []*EstimateMember

func (e EstimateMembers) Get(estimateID uuid.UUID, userID uuid.UUID) *EstimateMember {
	for _, x := range e {
		if x.EstimateID == estimateID && x.UserID == userID {
			return x
		}
	}
	return nil
}

func (e EstimateMembers) Clone() EstimateMembers {
	return slices.Clone(e)
}
