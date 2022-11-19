// Content managed by Project Forge, see [projectforge.md] for details.
package umember

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

type PK struct {
	StandupID uuid.UUID `json:"standupID"`
	UserID    uuid.UUID `json:"userID"`
}

type StandupMember struct {
	StandupID uuid.UUID         `json:"standupID"`
	UserID    uuid.UUID         `json:"userID"`
	Name      string            `json:"name"`
	Picture   string            `json:"picture"`
	Role      enum.MemberStatus `json:"role"`
	Created   time.Time         `json:"created"`
	Updated   *time.Time        `json:"updated,omitempty"`
}

func New(standupID uuid.UUID, userID uuid.UUID) *StandupMember {
	return &StandupMember{StandupID: standupID, UserID: userID}
}

func Random() *StandupMember {
	return &StandupMember{
		StandupID: util.UUID(),
		UserID:    util.UUID(),
		Name:      util.RandomString(12),
		Picture:   "https://" + util.RandomString(6) + ".com/" + util.RandomString(6),
		Role:      enum.MemberStatus(util.RandomString(12)),
		Created:   time.Now(),
		Updated:   util.NowPointer(),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*StandupMember, error) {
	ret := &StandupMember{}
	var err error
	if setPK {
		retStandupID, e := m.ParseUUID("standupID", true, true)
		if e != nil {
			return nil, e
		}
		if retStandupID != nil {
			ret.StandupID = *retStandupID
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

func (s *StandupMember) Clone() *StandupMember {
	return &StandupMember{
		StandupID: s.StandupID,
		UserID:    s.UserID,
		Name:      s.Name,
		Picture:   s.Picture,
		Role:      s.Role,
		Created:   s.Created,
		Updated:   s.Updated,
	}
}

func (s *StandupMember) String() string {
	return fmt.Sprintf("%s::%s", s.StandupID.String(), s.UserID.String())
}

func (s *StandupMember) TitleString() string {
	return s.StandupID.String() + " / " + s.Name
}

func (s *StandupMember) WebPath() string {
	return "/admin/db/standup/member/" + s.StandupID.String() + "/" + s.UserID.String()
}

func (s *StandupMember) Diff(sx *StandupMember) util.Diffs {
	var diffs util.Diffs
	if s.StandupID != sx.StandupID {
		diffs = append(diffs, util.NewDiff("standupID", s.StandupID.String(), sx.StandupID.String()))
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

func (s *StandupMember) ToData() []any {
	return []any{s.StandupID, s.UserID, s.Name, s.Picture, s.Role, s.Created, s.Updated}
}
