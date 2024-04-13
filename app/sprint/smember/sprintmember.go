// Package smember - Content managed by Project Forge, see [projectforge.md] for details.
package smember

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

var _ svc.Model = (*SprintMember)(nil)

type PK struct {
	SprintID uuid.UUID `json:"sprintID,omitempty"`
	UserID   uuid.UUID `json:"userID,omitempty"`
}

func (p *PK) String() string {
	return fmt.Sprintf("%v::%v", p.SprintID, p.UserID)
}

type SprintMember struct {
	SprintID uuid.UUID         `json:"sprintID,omitempty"`
	UserID   uuid.UUID         `json:"userID,omitempty"`
	Name     string            `json:"name,omitempty"`
	Picture  string            `json:"picture,omitempty"`
	Role     enum.MemberStatus `json:"role,omitempty"`
	Created  time.Time         `json:"created,omitempty"`
	Updated  *time.Time        `json:"updated,omitempty"`
}

func New(sprintID uuid.UUID, userID uuid.UUID) *SprintMember {
	return &SprintMember{SprintID: sprintID, UserID: userID}
}

func (s *SprintMember) Clone() *SprintMember {
	return &SprintMember{s.SprintID, s.UserID, s.Name, s.Picture, s.Role, s.Created, s.Updated}
}

func (s *SprintMember) String() string {
	return fmt.Sprintf("%s::%s", s.SprintID.String(), s.UserID.String())
}

func (s *SprintMember) TitleString() string {
	return s.SprintID.String() + " / " + s.Name
}

func (s *SprintMember) ToPK() *PK {
	return &PK{
		SprintID: s.SprintID,
		UserID:   s.UserID,
	}
}

func Random() *SprintMember {
	return &SprintMember{
		SprintID: util.UUID(),
		UserID:   util.UUID(),
		Name:     util.RandomString(12),
		Picture:  util.RandomURL().String(),
		Role:     enum.AllMemberStatuses.Random(),
		Created:  util.TimeCurrent(),
		Updated:  util.TimeCurrentP(),
	}
}

func (s *SprintMember) Strings() []string {
	return []string{s.SprintID.String(), s.UserID.String(), s.Name, s.Picture, s.Role.String(), util.TimeToFull(&s.Created), util.TimeToFull(s.Updated)}
}

func (s *SprintMember) ToCSV() ([]string, [][]string) {
	return FieldDescs.Keys(), [][]string{s.Strings()}
}

func (s *SprintMember) WebPath() string {
	return "/admin/db/sprint/member/" + s.SprintID.String() + "/" + s.UserID.String()
}

func (s *SprintMember) ToData() []any {
	return []any{s.SprintID, s.UserID, s.Name, s.Picture, s.Role, s.Created, s.Updated}
}

var FieldDescs = util.FieldDescs{
	{Key: "sprintID", Title: "Sprint ID", Description: "", Type: "uuid"},
	{Key: "userID", Title: "User ID", Description: "", Type: "uuid"},
	{Key: "name", Title: "Name", Description: "", Type: "string"},
	{Key: "picture", Title: "Picture", Description: "", Type: "string"},
	{Key: "role", Title: "Role", Description: "", Type: "enum(member_status)"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
	{Key: "updated", Title: "Updated", Description: "", Type: "timestamp"},
}
