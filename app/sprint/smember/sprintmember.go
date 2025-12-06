package smember

import (
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

const DefaultRoute = "/admin/db/sprint/member"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(paths...)
}

var _ svc.Model = (*SprintMember)(nil)

type PK struct {
	SprintID uuid.UUID `json:"sprintID,omitzero"`
	UserID   uuid.UUID `json:"userID,omitzero"`
}

func (p *PK) String() string {
	return fmt.Sprintf("%v • %v", p.SprintID, p.UserID)
}

type SprintMember struct {
	SprintID uuid.UUID         `json:"sprintID,omitzero"`
	UserID   uuid.UUID         `json:"userID,omitzero"`
	Name     string            `json:"name,omitzero"`
	Picture  string            `json:"picture,omitzero"`
	Role     enum.MemberStatus `json:"role,omitzero"`
	Created  time.Time         `json:"created,omitzero"`
	Updated  *time.Time        `json:"updated,omitzero"`
}

func NewSprintMember(sprintID uuid.UUID, userID uuid.UUID) *SprintMember {
	return &SprintMember{SprintID: sprintID, UserID: userID}
}

func (s *SprintMember) Clone() *SprintMember {
	if s == nil {
		return nil
	}
	return &SprintMember{
		SprintID: s.SprintID, UserID: s.UserID, Name: s.Name, Picture: s.Picture, Role: s.Role, Created: s.Created,
		Updated: s.Updated,
	}
}

func (s *SprintMember) String() string {
	return fmt.Sprintf("%s • %s", s.SprintID.String(), s.UserID.String())
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

func RandomSprintMember() *SprintMember {
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
	return SprintMemberFieldDescs.Keys(), [][]string{s.Strings()}
}

func (s *SprintMember) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(append(paths, url.QueryEscape(s.SprintID.String()), url.QueryEscape(s.UserID.String()))...)
}

func (s *SprintMember) Breadcrumb(extra ...string) string {
	return s.TitleString() + "||" + s.WebPath(extra...) + "**users"
}

func (s *SprintMember) ToData() []any {
	return []any{s.SprintID, s.UserID, s.Name, s.Picture, s.Role, s.Created, s.Updated}
}

var SprintMemberFieldDescs = util.FieldDescs{
	{Key: "sprintID", Title: "Sprint ID", Type: "uuid"},
	{Key: "userID", Title: "User ID", Type: "uuid"},
	{Key: "name", Title: "Name", Type: "string"},
	{Key: "picture", Title: "Picture", Type: "string"},
	{Key: "role", Title: "Role", Type: "enum(member_status)"},
	{Key: "created", Title: "Created", Type: "timestamp"},
	{Key: "updated", Title: "Updated", Type: "timestamp"},
}
