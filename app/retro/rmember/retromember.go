// Content managed by Project Forge, see [projectforge.md] for details.
package rmember

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

type PK struct {
	RetroID uuid.UUID `json:"retroID"`
	UserID  uuid.UUID `json:"userID"`
}

type RetroMember struct {
	RetroID uuid.UUID         `json:"retroID"`
	UserID  uuid.UUID         `json:"userID"`
	Name    string            `json:"name"`
	Picture string            `json:"picture"`
	Role    enum.MemberStatus `json:"role"`
	Created time.Time         `json:"created"`
	Updated *time.Time        `json:"updated,omitempty"`
}

func New(retroID uuid.UUID, userID uuid.UUID) *RetroMember {
	return &RetroMember{RetroID: retroID, UserID: userID}
}

func Random() *RetroMember {
	return &RetroMember{
		RetroID: util.UUID(),
		UserID:  util.UUID(),
		Name:    util.RandomString(12),
		Picture: "https://" + util.RandomString(6) + ".com/" + util.RandomString(6),
		Role:    enum.MemberStatus(util.RandomString(12)),
		Created: util.TimeCurrent(),
		Updated: util.TimeCurrentP(),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*RetroMember, error) {
	ret := &RetroMember{}
	var err error
	if setPK {
		retRetroID, e := m.ParseUUID("retroID", true, true)
		if e != nil {
			return nil, e
		}
		if retRetroID != nil {
			ret.RetroID = *retRetroID
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
	retRole, err := m.ParseString("role", true, true)
	if err != nil {
		return nil, err
	}
	ret.Role = enum.MemberStatus(retRole)
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (r *RetroMember) Clone() *RetroMember {
	return &RetroMember{r.RetroID, r.UserID, r.Name, r.Picture, r.Role, r.Created, r.Updated}
}

func (r *RetroMember) String() string {
	return fmt.Sprintf("%s::%s", r.RetroID.String(), r.UserID.String())
}

func (r *RetroMember) TitleString() string {
	return r.RetroID.String() + " / " + r.Name
}

func (r *RetroMember) ToPK() *PK {
	return &PK{
		RetroID: r.RetroID,
		UserID:  r.UserID,
	}
}

func (r *RetroMember) WebPath() string {
	return "/admin/db/retro/member/" + r.RetroID.String() + "/" + r.UserID.String()
}

func (r *RetroMember) Diff(rx *RetroMember) util.Diffs {
	var diffs util.Diffs
	if r.RetroID != rx.RetroID {
		diffs = append(diffs, util.NewDiff("retroID", r.RetroID.String(), rx.RetroID.String()))
	}
	if r.UserID != rx.UserID {
		diffs = append(diffs, util.NewDiff("userID", r.UserID.String(), rx.UserID.String()))
	}
	if r.Name != rx.Name {
		diffs = append(diffs, util.NewDiff("name", r.Name, rx.Name))
	}
	if r.Picture != rx.Picture {
		diffs = append(diffs, util.NewDiff("picture", r.Picture, rx.Picture))
	}
	if r.Role != rx.Role {
		diffs = append(diffs, util.NewDiff("role", string(r.Role), string(rx.Role)))
	}
	if r.Created != rx.Created {
		diffs = append(diffs, util.NewDiff("created", r.Created.String(), rx.Created.String()))
	}
	return diffs
}

func (r *RetroMember) ToData() []any {
	return []any{r.RetroID, r.UserID, r.Name, r.Picture, r.Role, r.Created, r.Updated}
}
