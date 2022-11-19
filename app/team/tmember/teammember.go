// Content managed by Project Forge, see [projectforge.md] for details.
package tmember

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

type PK struct {
	TeamID uuid.UUID `json:"teamID"`
	UserID uuid.UUID `json:"userID"`
}

type TeamMember struct {
	TeamID  uuid.UUID         `json:"teamID"`
	UserID  uuid.UUID         `json:"userID"`
	Name    string            `json:"name"`
	Picture string            `json:"picture"`
	Role    enum.MemberStatus `json:"role"`
	Created time.Time         `json:"created"`
	Updated *time.Time        `json:"updated,omitempty"`
}

func New(teamID uuid.UUID, userID uuid.UUID) *TeamMember {
	return &TeamMember{TeamID: teamID, UserID: userID}
}

func Random() *TeamMember {
	return &TeamMember{
		TeamID:  util.UUID(),
		UserID:  util.UUID(),
		Name:    util.RandomString(12),
		Picture: "https://" + util.RandomString(6) + ".com/" + util.RandomString(6),
		Role:    enum.MemberStatus(util.RandomString(12)),
		Created: time.Now(),
		Updated: util.NowPointer(),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*TeamMember, error) {
	ret := &TeamMember{}
	var err error
	if setPK {
		retTeamID, e := m.ParseUUID("teamID", true, true)
		if e != nil {
			return nil, e
		}
		if retTeamID != nil {
			ret.TeamID = *retTeamID
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

func (t *TeamMember) Clone() *TeamMember {
	return &TeamMember{
		TeamID:  t.TeamID,
		UserID:  t.UserID,
		Name:    t.Name,
		Picture: t.Picture,
		Role:    t.Role,
		Created: t.Created,
		Updated: t.Updated,
	}
}

func (t *TeamMember) String() string {
	return fmt.Sprintf("%s::%s", t.TeamID.String(), t.UserID.String())
}

func (t *TeamMember) TitleString() string {
	return t.TeamID.String() + " / " + t.Name
}

func (t *TeamMember) WebPath() string {
	return "/admin/db/team/member/" + t.TeamID.String() + "/" + t.UserID.String()
}

func (t *TeamMember) Diff(tx *TeamMember) util.Diffs {
	var diffs util.Diffs
	if t.TeamID != tx.TeamID {
		diffs = append(diffs, util.NewDiff("teamID", t.TeamID.String(), tx.TeamID.String()))
	}
	if t.UserID != tx.UserID {
		diffs = append(diffs, util.NewDiff("userID", t.UserID.String(), tx.UserID.String()))
	}
	if t.Name != tx.Name {
		diffs = append(diffs, util.NewDiff("name", t.Name, tx.Name))
	}
	if t.Picture != tx.Picture {
		diffs = append(diffs, util.NewDiff("picture", t.Picture, tx.Picture))
	}
	if t.Role != tx.Role {
		diffs = append(diffs, util.NewDiff("role", string(t.Role), string(tx.Role)))
	}
	if t.Created != tx.Created {
		diffs = append(diffs, util.NewDiff("created", t.Created.String(), tx.Created.String()))
	}
	return diffs
}

func (t *TeamMember) ToData() []any {
	return []any{t.TeamID, t.UserID, t.Name, t.Picture, t.Role, t.Created, t.Updated}
}
