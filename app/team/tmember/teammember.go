// Package tmember - Content managed by Project Forge, see [projectforge.md] for details.
package tmember

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

type PK struct {
	TeamID uuid.UUID `json:"teamID,omitempty"`
	UserID uuid.UUID `json:"userID,omitempty"`
}

func (p *PK) String() string {
	return fmt.Sprintf("%v::%v", p.TeamID, p.UserID)
}

type TeamMember struct {
	TeamID  uuid.UUID         `json:"teamID,omitempty"`
	UserID  uuid.UUID         `json:"userID,omitempty"`
	Name    string            `json:"name,omitempty"`
	Picture string            `json:"picture,omitempty"`
	Role    enum.MemberStatus `json:"role,omitempty"`
	Created time.Time         `json:"created,omitempty"`
	Updated *time.Time        `json:"updated,omitempty"`
}

func New(teamID uuid.UUID, userID uuid.UUID) *TeamMember {
	return &TeamMember{TeamID: teamID, UserID: userID}
}

func (t *TeamMember) Clone() *TeamMember {
	return &TeamMember{t.TeamID, t.UserID, t.Name, t.Picture, t.Role, t.Created, t.Updated}
}

func (t *TeamMember) String() string {
	return fmt.Sprintf("%s::%s", t.TeamID.String(), t.UserID.String())
}

func (t *TeamMember) TitleString() string {
	return t.TeamID.String() + " / " + t.Name
}

func (t *TeamMember) ToPK() *PK {
	return &PK{
		TeamID: t.TeamID,
		UserID: t.UserID,
	}
}

func Random() *TeamMember {
	return &TeamMember{
		TeamID:  util.UUID(),
		UserID:  util.UUID(),
		Name:    util.RandomString(12),
		Picture: "https://" + util.RandomString(6) + ".com/" + util.RandomString(6),
		Role:    enum.AllMemberStatuses.Random(),
		Created: util.TimeCurrent(),
		Updated: util.TimeCurrentP(),
	}
}

func (t *TeamMember) WebPath() string {
	return "/admin/db/team/member/" + t.TeamID.String() + "/" + t.UserID.String()
}

func (t *TeamMember) ToData() []any {
	return []any{t.TeamID, t.UserID, t.Name, t.Picture, t.Role, t.Created, t.Updated}
}
