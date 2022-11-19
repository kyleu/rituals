// Content managed by Project Forge, see [projectforge.md] for details.
package smember

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

type PK struct {
	SprintID uuid.UUID `json:"sprintID"`
	UserID   uuid.UUID `json:"userID"`
}

type SprintMember struct {
	SprintID uuid.UUID         `json:"sprintID"`
	UserID   uuid.UUID         `json:"userID"`
	Name     string            `json:"name"`
	Picture  string            `json:"picture"`
	Role     enum.MemberStatus `json:"role"`
	Created  time.Time         `json:"created"`
	Updated  *time.Time        `json:"updated,omitempty"`
}

func New(sprintID uuid.UUID, userID uuid.UUID) *SprintMember {
	return &SprintMember{SprintID: sprintID, UserID: userID}
}

func Random() *SprintMember {
	return &SprintMember{
		SprintID: util.UUID(),
		UserID:   util.UUID(),
		Name:     util.RandomString(12),
		Picture:  "https://" + util.RandomString(6) + ".com/" + util.RandomString(6),
		Role:     enum.MemberStatus(util.RandomString(12)),
		Created:  time.Now(),
		Updated:  util.NowPointer(),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*SprintMember, error) {
	ret := &SprintMember{}
	var err error
	if setPK {
		retSprintID, e := m.ParseUUID("sprintID", true, true)
		if e != nil {
			return nil, e
		}
		if retSprintID != nil {
			ret.SprintID = *retSprintID
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

func (s *SprintMember) Clone() *SprintMember {
	return &SprintMember{
		SprintID: s.SprintID,
		UserID:   s.UserID,
		Name:     s.Name,
		Picture:  s.Picture,
		Role:     s.Role,
		Created:  s.Created,
		Updated:  s.Updated,
	}
}

func (s *SprintMember) String() string {
	return fmt.Sprintf("%s::%s", s.SprintID.String(), s.UserID.String())
}

func (s *SprintMember) TitleString() string {
	return s.SprintID.String() + " / " + s.Name
}

func (s *SprintMember) WebPath() string {
	return "/admin/db/sprint/member/" + s.SprintID.String() + "/" + s.UserID.String()
}

func (s *SprintMember) Diff(sx *SprintMember) util.Diffs {
	var diffs util.Diffs
	if s.SprintID != sx.SprintID {
		diffs = append(diffs, util.NewDiff("sprintID", s.SprintID.String(), sx.SprintID.String()))
	}
	if s.UserID != sx.UserID {
		diffs = append(diffs, util.NewDiff("userID", s.UserID.String(), sx.UserID.String()))
	}
	if s.Name != sx.Name {
		diffs = append(diffs, util.NewDiff("name", s.Name, sx.Name))
	}
	if s.Picture != sx.Picture {
		diffs = append(diffs, util.NewDiff("picture", s.Picture, sx.Picture))
	}
	if s.Role != sx.Role {
		diffs = append(diffs, util.NewDiff("role", string(s.Role), string(sx.Role)))
	}
	if s.Created != sx.Created {
		diffs = append(diffs, util.NewDiff("created", s.Created.String(), sx.Created.String()))
	}
	return diffs
}

func (s *SprintMember) ToData() []any {
	return []any{s.SprintID, s.UserID, s.Name, s.Picture, s.Role, s.Created, s.Updated}
}
