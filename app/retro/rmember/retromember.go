package rmember

import (
	"fmt"
	"net/url"
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
	return util.StringPath(paths...)
}

var _ svc.Model = (*RetroMember)(nil)

type PK struct {
	RetroID uuid.UUID `json:"retroID,omitzero"`
	UserID  uuid.UUID `json:"userID,omitzero"`
}

func (p *PK) String() string {
	return fmt.Sprintf("%v • %v", p.RetroID, p.UserID)
}

type RetroMember struct {
	RetroID uuid.UUID         `json:"retroID,omitzero"`
	UserID  uuid.UUID         `json:"userID,omitzero"`
	Name    string            `json:"name,omitzero"`
	Picture string            `json:"picture,omitzero"`
	Role    enum.MemberStatus `json:"role,omitzero"`
	Created time.Time         `json:"created,omitzero"`
	Updated *time.Time        `json:"updated,omitzero"`
}

func NewRetroMember(retroID uuid.UUID, userID uuid.UUID) *RetroMember {
	return &RetroMember{RetroID: retroID, UserID: userID}
}

func (r *RetroMember) Clone() *RetroMember {
	if r == nil {
		return nil
	}
	return &RetroMember{
		RetroID: r.RetroID, UserID: r.UserID, Name: r.Name, Picture: r.Picture, Role: r.Role, Created: r.Created,
		Updated: r.Updated,
	}
}

func (r *RetroMember) String() string {
	return fmt.Sprintf("%s • %s", r.RetroID.String(), r.UserID.String())
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
	return util.StringPath(append(paths, url.QueryEscape(r.RetroID.String()), url.QueryEscape(r.UserID.String()))...)
}

func (r *RetroMember) Breadcrumb(extra ...string) string {
	return r.TitleString() + "||" + r.WebPath(extra...) + "**users"
}

func (r *RetroMember) ToData() []any {
	return []any{r.RetroID, r.UserID, r.Name, r.Picture, r.Role, r.Created, r.Updated}
}

var RetroMemberFieldDescs = util.FieldDescs{
	{Key: "retroID", Title: "Retro ID", Type: "uuid"},
	{Key: "userID", Title: "User ID", Type: "uuid"},
	{Key: "name", Title: "Name", Type: "string"},
	{Key: "picture", Title: "Picture", Type: "string"},
	{Key: "role", Title: "Role", Type: "enum(member_status)"},
	{Key: "created", Title: "Created", Type: "timestamp"},
	{Key: "updated", Title: "Updated", Type: "timestamp"},
}
