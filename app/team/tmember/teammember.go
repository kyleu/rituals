package tmember

import (
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

const DefaultRoute = "/admin/db/team/member"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(paths...)
}

var _ svc.Model = (*TeamMember)(nil)

type PK struct {
	TeamID uuid.UUID `json:"teamID,omitzero"`
	UserID uuid.UUID `json:"userID,omitzero"`
}

func (p *PK) String() string {
	return fmt.Sprintf("%v • %v", p.TeamID, p.UserID)
}

type TeamMember struct {
	TeamID  uuid.UUID         `json:"teamID,omitzero"`
	UserID  uuid.UUID         `json:"userID,omitzero"`
	Name    string            `json:"name,omitzero"`
	Picture string            `json:"picture,omitzero"`
	Role    enum.MemberStatus `json:"role,omitzero"`
	Created time.Time         `json:"created,omitzero"`
	Updated *time.Time        `json:"updated,omitzero"`
}

func NewTeamMember(teamID uuid.UUID, userID uuid.UUID) *TeamMember {
	return &TeamMember{TeamID: teamID, UserID: userID}
}

func (t *TeamMember) Clone() *TeamMember {
	return &TeamMember{
		TeamID: t.TeamID, UserID: t.UserID, Name: t.Name, Picture: t.Picture, Role: t.Role, Created: t.Created,
		Updated: t.Updated,
	}
}

func (t *TeamMember) String() string {
	return fmt.Sprintf("%s • %s", t.TeamID.String(), t.UserID.String())
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

func RandomTeamMember() *TeamMember {
	return &TeamMember{
		TeamID:  util.UUID(),
		UserID:  util.UUID(),
		Name:    util.RandomString(12),
		Picture: util.RandomURL().String(),
		Role:    enum.AllMemberStatuses.Random(),
		Created: util.TimeCurrent(),
		Updated: util.TimeCurrentP(),
	}
}

func (t *TeamMember) Strings() []string {
	return []string{t.TeamID.String(), t.UserID.String(), t.Name, t.Picture, t.Role.String(), util.TimeToFull(&t.Created), util.TimeToFull(t.Updated)}
}

func (t *TeamMember) ToCSV() ([]string, [][]string) {
	return TeamMemberFieldDescs.Keys(), [][]string{t.Strings()}
}

func (t *TeamMember) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(append(paths, url.QueryEscape(t.TeamID.String()), url.QueryEscape(t.UserID.String()))...)
}

func (t *TeamMember) Breadcrumb(extra ...string) string {
	return t.TitleString() + "||" + t.WebPath(extra...) + "**users"
}

func (t *TeamMember) ToData() []any {
	return []any{t.TeamID, t.UserID, t.Name, t.Picture, t.Role, t.Created, t.Updated}
}

var TeamMemberFieldDescs = util.FieldDescs{
	{Key: "teamID", Title: "Team ID", Description: "", Type: "uuid"},
	{Key: "userID", Title: "User ID", Description: "", Type: "uuid"},
	{Key: "name", Title: "Name", Description: "", Type: "string"},
	{Key: "picture", Title: "Picture", Description: "", Type: "string"},
	{Key: "role", Title: "Role", Description: "", Type: "enum(member_status)"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
	{Key: "updated", Title: "Updated", Description: "", Type: "timestamp"},
}
