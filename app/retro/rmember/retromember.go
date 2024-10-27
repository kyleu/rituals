package rmember

import (
	"fmt"
	"net/url"
	"path"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

const DefaultRoute = "/admin/db/retro/member"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return path.Join(paths...)
}

var _ svc.Model = (*RetroMember)(nil)

type PK struct {
	RetroID uuid.UUID `json:"retroID,omitempty"`
	UserID  uuid.UUID `json:"userID,omitempty"`
}

func (p *PK) String() string {
	return fmt.Sprintf("%v::%v", p.RetroID, p.UserID)
}

type RetroMember struct {
	RetroID uuid.UUID         `json:"retroID,omitempty"`
	UserID  uuid.UUID         `json:"userID,omitempty"`
	Name    string            `json:"name,omitempty"`
	Picture string            `json:"picture,omitempty"`
	Role    enum.MemberStatus `json:"role,omitempty"`
	Created time.Time         `json:"created,omitempty"`
	Updated *time.Time        `json:"updated,omitempty"`
}

func NewRetroMember(retroID uuid.UUID, userID uuid.UUID) *RetroMember {
	return &RetroMember{RetroID: retroID, UserID: userID}
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

func RandomRetroMember() *RetroMember {
	return &RetroMember{
		RetroID: util.UUID(),
		UserID:  util.UUID(),
		Name:    util.RandomString(12),
		Picture: util.RandomURL().String(),
		Role:    enum.AllMemberStatuses.Random(),
		Created: util.TimeCurrent(),
		Updated: util.TimeCurrentP(),
	}
}

func (r *RetroMember) Strings() []string {
	return []string{r.RetroID.String(), r.UserID.String(), r.Name, r.Picture, r.Role.String(), util.TimeToFull(&r.Created), util.TimeToFull(r.Updated)}
}

func (r *RetroMember) ToCSV() ([]string, [][]string) {
	return RetroMemberFieldDescs.Keys(), [][]string{r.Strings()}
}

func (r *RetroMember) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return path.Join(append(paths, url.QueryEscape(r.RetroID.String()), url.QueryEscape(r.UserID.String()))...)
}

func (r *RetroMember) ToData() []any {
	return []any{r.RetroID, r.UserID, r.Name, r.Picture, r.Role, r.Created, r.Updated}
}

var RetroMemberFieldDescs = util.FieldDescs{
	{Key: "retroID", Title: "Retro ID", Description: "", Type: "uuid"},
	{Key: "userID", Title: "User ID", Description: "", Type: "uuid"},
	{Key: "name", Title: "Name", Description: "", Type: "string"},
	{Key: "picture", Title: "Picture", Description: "", Type: "string"},
	{Key: "role", Title: "Role", Description: "", Type: "enum(member_status)"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
	{Key: "updated", Title: "Updated", Description: "", Type: "timestamp"},
}
