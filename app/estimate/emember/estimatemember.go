// Package emember - Content managed by Project Forge, see [projectforge.md] for details.
package emember

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

type PK struct {
	EstimateID uuid.UUID `json:"estimateID,omitempty"`
	UserID     uuid.UUID `json:"userID,omitempty"`
}

func (p *PK) String() string {
	return fmt.Sprintf("%v::%v", p.EstimateID, p.UserID)
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

func New(estimateID uuid.UUID, userID uuid.UUID) *EstimateMember {
	return &EstimateMember{EstimateID: estimateID, UserID: userID}
}

func (e *EstimateMember) Clone() *EstimateMember {
	return &EstimateMember{e.EstimateID, e.UserID, e.Name, e.Picture, e.Role, e.Created, e.Updated}
}

func (e *EstimateMember) String() string {
	return fmt.Sprintf("%s::%s", e.EstimateID.String(), e.UserID.String())
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

func Random() *EstimateMember {
	return &EstimateMember{
		EstimateID: util.UUID(),
		UserID:     util.UUID(),
		Name:       util.RandomString(12),
		Picture:    "https://" + util.RandomString(6) + ".com/" + util.RandomString(6),
		Role:       enum.AllMemberStatuses.Random(),
		Created:    util.TimeCurrent(),
		Updated:    util.TimeCurrentP(),
	}
}

func (e *EstimateMember) WebPath() string {
	return "/admin/db/estimate/member/" + e.EstimateID.String() + "/" + e.UserID.String()
}

func (e *EstimateMember) ToData() []any {
	return []any{e.EstimateID, e.UserID, e.Name, e.Picture, e.Role, e.Created, e.Updated}
}

var FieldDescs = util.FieldDescs{
	{Key: "estimateID", Title: "Estimate ID", Description: "", Type: "uuid"},
	{Key: "userID", Title: "User ID", Description: "", Type: "uuid"},
	{Key: "name", Title: "Name", Description: "", Type: "string"},
	{Key: "picture", Title: "Picture", Description: "", Type: "string"},
	{Key: "role", Title: "Role", Description: "", Type: "enum(member_status)"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
	{Key: "updated", Title: "Updated", Description: "", Type: "timestamp"},
}
