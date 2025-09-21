package umember

import (
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

const DefaultRoute = "/admin/db/standup/member"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(paths...)
}

var _ svc.Model = (*StandupMember)(nil)

type PK struct {
	StandupID uuid.UUID `json:"standupID,omitzero"`
	UserID    uuid.UUID `json:"userID,omitzero"`
}

func (p *PK) String() string {
	return fmt.Sprintf("%v • %v", p.StandupID, p.UserID)
}

type StandupMember struct {
	StandupID uuid.UUID         `json:"standupID,omitzero"`
	UserID    uuid.UUID         `json:"userID,omitzero"`
	Name      string            `json:"name,omitzero"`
	Picture   string            `json:"picture,omitzero"`
	Role      enum.MemberStatus `json:"role,omitzero"`
	Created   time.Time         `json:"created,omitzero"`
	Updated   *time.Time        `json:"updated,omitzero"`
}

func NewStandupMember(standupID uuid.UUID, userID uuid.UUID) *StandupMember {
	return &StandupMember{StandupID: standupID, UserID: userID}
}

func (s *StandupMember) Clone() *StandupMember {
	return &StandupMember{s.StandupID, s.UserID, s.Name, s.Picture, s.Role, s.Created, s.Updated}
}

func (s *StandupMember) String() string {
	return fmt.Sprintf("%s • %s", s.StandupID.String(), s.UserID.String())
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

func RandomStandupMember() *StandupMember {
	return &StandupMember{
		StandupID: util.UUID(),
		UserID:    util.UUID(),
		Name:      util.RandomString(12),
		Picture:   util.RandomURL().String(),
		Role:      enum.AllMemberStatuses.Random(),
		Created:   util.TimeCurrent(),
		Updated:   util.TimeCurrentP(),
	}
}

func (s *StandupMember) Strings() []string {
	return []string{s.StandupID.String(), s.UserID.String(), s.Name, s.Picture, s.Role.String(), util.TimeToFull(&s.Created), util.TimeToFull(s.Updated)}
}

func (s *StandupMember) ToCSV() ([]string, [][]string) {
	return StandupMemberFieldDescs.Keys(), [][]string{s.Strings()}
}

func (s *StandupMember) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(append(paths, url.QueryEscape(s.StandupID.String()), url.QueryEscape(s.UserID.String()))...)
}

func (s *StandupMember) Breadcrumb(extra ...string) string {
	return s.TitleString() + "||" + s.WebPath(extra...) + "**users"
}

func (s *StandupMember) ToData() []any {
	return []any{s.StandupID, s.UserID, s.Name, s.Picture, s.Role, s.Created, s.Updated}
}

var StandupMemberFieldDescs = util.FieldDescs{
	{Key: "standupID", Title: "Standup ID", Description: "", Type: "uuid"},
	{Key: "userID", Title: "User ID", Description: "", Type: "uuid"},
	{Key: "name", Title: "Name", Description: "", Type: "string"},
	{Key: "picture", Title: "Picture", Description: "", Type: "string"},
	{Key: "role", Title: "Role", Description: "", Type: "enum(member_status)"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
	{Key: "updated", Title: "Updated", Description: "", Type: "timestamp"},
}
