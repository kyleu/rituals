package emember

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

const DefaultRoute = "/admin/db/estimate/member"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return path.Join(paths...)
}

var _ svc.Model = (*EstimateMember)(nil)

type PK struct {
	EstimateID uuid.UUID `json:"estimateID,omitempty"`
	UserID     uuid.UUID `json:"userID,omitempty"`
}

func (p *PK) String() string {
	return fmt.Sprintf("%v • %v", p.EstimateID, p.UserID)
}

type EstimateMember struct {
	EstimateID uuid.UUID         `json:"estimateID,omitempty"`
	UserID     uuid.UUID         `json:"userID,omitempty"`
	Name       string            `json:"name,omitempty"`
	Picture    string            `json:"picture,omitempty"`
	Role       enum.MemberStatus `json:"role,omitempty"`
	Created    time.Time         `json:"created,omitempty"`
	Updated    *time.Time        `json:"updated,omitempty"`
}

func NewEstimateMember(estimateID uuid.UUID, userID uuid.UUID) *EstimateMember {
	return &EstimateMember{EstimateID: estimateID, UserID: userID}
}

func (e *EstimateMember) Clone() *EstimateMember {
	return &EstimateMember{e.EstimateID, e.UserID, e.Name, e.Picture, e.Role, e.Created, e.Updated}
}

func (e *EstimateMember) String() string {
	return fmt.Sprintf("%s • %s", e.EstimateID.String(), e.UserID.String())
}

func (e *EstimateMember) TitleString() string {
	return e.EstimateID.String() + " / " + e.Name
}

func (e *EstimateMember) ToPK() *PK {
	return &PK{
		EstimateID: e.EstimateID,
		UserID:     e.UserID,
	}
}

func RandomEstimateMember() *EstimateMember {
	return &EstimateMember{
		EstimateID: util.UUID(),
		UserID:     util.UUID(),
		Name:       util.RandomString(12),
		Picture:    util.RandomURL().String(),
		Role:       enum.AllMemberStatuses.Random(),
		Created:    util.TimeCurrent(),
		Updated:    util.TimeCurrentP(),
	}
}

func (e *EstimateMember) Strings() []string {
	return []string{e.EstimateID.String(), e.UserID.String(), e.Name, e.Picture, e.Role.String(), util.TimeToFull(&e.Created), util.TimeToFull(e.Updated)}
}

func (e *EstimateMember) ToCSV() ([]string, [][]string) {
	return EstimateMemberFieldDescs.Keys(), [][]string{e.Strings()}
}

func (e *EstimateMember) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return path.Join(append(paths, url.QueryEscape(e.EstimateID.String()), url.QueryEscape(e.UserID.String()))...)
}

func (e *EstimateMember) ToData() []any {
	return []any{e.EstimateID, e.UserID, e.Name, e.Picture, e.Role, e.Created, e.Updated}
}

var EstimateMemberFieldDescs = util.FieldDescs{
	{Key: "estimateID", Title: "Estimate ID", Description: "", Type: "uuid"},
	{Key: "userID", Title: "User ID", Description: "", Type: "uuid"},
	{Key: "name", Title: "Name", Description: "", Type: "string"},
	{Key: "picture", Title: "Picture", Description: "", Type: "string"},
	{Key: "role", Title: "Role", Description: "", Type: "enum(member_status)"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
	{Key: "updated", Title: "Updated", Description: "", Type: "timestamp"},
}
