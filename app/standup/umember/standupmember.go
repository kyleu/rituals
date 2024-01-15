// Package umember - Content managed by Project Forge, see [projectforge.md] for details.
package umember

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

type PK struct {
	StandupID uuid.UUID `json:"standupID,omitempty"`
	UserID    uuid.UUID `json:"userID,omitempty"`
}

func (p *PK) String() string {
	return fmt.Sprintf("%v::%v", p.StandupID, p.UserID)
}

type StandupMember struct {
	StandupID uuid.UUID         `json:"standupID,omitempty"`
	UserID    uuid.UUID         `json:"userID,omitempty"`
	Name      string            `json:"name,omitempty"`
	Picture   string            `json:"picture,omitempty"`
	Role      enum.MemberStatus `json:"role,omitempty"`
	Created   time.Time         `json:"created,omitempty"`
	Updated   *time.Time        `json:"updated,omitempty"`
}

func New(standupID uuid.UUID, userID uuid.UUID) *StandupMember {
	return &StandupMember{StandupID: standupID, UserID: userID}
}

func (s *StandupMember) Clone() *StandupMember {
	return &StandupMember{s.StandupID, s.UserID, s.Name, s.Picture, s.Role, s.Created, s.Updated}
}

func (s *StandupMember) String() string {
	return fmt.Sprintf("%s::%s", s.StandupID.String(), s.UserID.String())
}

func (s *StandupMember) TitleString() string {
	return s.StandupID.String() + " / " + s.Name
}

func (s *StandupMember) ToPK() *PK {
	return &PK{
		StandupID: s.StandupID,
		UserID:    s.UserID,
	}
}

func Random() *StandupMember {
	return &StandupMember{
		StandupID: util.UUID(),
		UserID:    util.UUID(),
		Name:      util.RandomString(12),
		Picture:   "https://" + util.RandomString(6) + ".com/" + util.RandomString(6),
		Role:      enum.AllMemberStatuses.Random(),
		Created:   util.TimeCurrent(),
		Updated:   util.TimeCurrentP(),
	}
}

func (s *StandupMember) WebPath() string {
	return "/admin/db/standup/member/" + s.StandupID.String() + "/" + s.UserID.String()
}

func (s *StandupMember) ToData() []any {
	return []any{s.StandupID, s.UserID, s.Name, s.Picture, s.Role, s.Created, s.Updated}
}

var FieldDescs = util.FieldDescs{
	{Key: "standupID", Title: "Standup ID", Description: "", Type: "uuid"},
	{Key: "userID", Title: "User ID", Description: "", Type: "uuid"},
	{Key: "name", Title: "Name", Description: "", Type: "string"},
	{Key: "picture", Title: "Picture", Description: "", Type: "string"},
	{Key: "role", Title: "Role", Description: "", Type: "enum(member_status)"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
	{Key: "updated", Title: "Updated", Description: "", Type: "timestamp"},
}
