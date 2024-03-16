// Package vote - Content managed by Project Forge, see [projectforge.md] for details.
package vote

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

type PK struct {
	StoryID uuid.UUID `json:"storyID,omitempty"`
	UserID  uuid.UUID `json:"userID,omitempty"`
}

func (p *PK) String() string {
	return fmt.Sprintf("%v::%v", p.StoryID, p.UserID)
}

type Vote struct {
	StoryID uuid.UUID  `json:"storyID,omitempty"`
	UserID  uuid.UUID  `json:"userID,omitempty"`
	Choice  string     `json:"choice,omitempty"`
	Created time.Time  `json:"created,omitempty"`
	Updated *time.Time `json:"updated,omitempty"`
}

func New(storyID uuid.UUID, userID uuid.UUID) *Vote {
	return &Vote{StoryID: storyID, UserID: userID}
}

func (v *Vote) Clone() *Vote {
	return &Vote{v.StoryID, v.UserID, v.Choice, v.Created, v.Updated}
}

func (v *Vote) String() string {
	return fmt.Sprintf("%s::%s", v.StoryID.String(), v.UserID.String())
}

func (v *Vote) TitleString() string {
	return v.String()
}

func (v *Vote) ToPK() *PK {
	return &PK{
		StoryID: v.StoryID,
		UserID:  v.UserID,
	}
}

func Random() *Vote {
	return &Vote{
		StoryID: util.UUID(),
		UserID:  util.UUID(),
		Choice:  util.RandomString(12),
		Created: util.TimeCurrent(),
		Updated: util.TimeCurrentP(),
	}
}

func (v *Vote) Strings() []string {
	return []string{v.StoryID.String(), v.UserID.String(), v.Choice, util.TimeToFull(&v.Created), util.TimeToFull(v.Updated)}
}

func (v *Vote) ToCSV() ([]string, [][]string) {
	return FieldDescs.Keys(), [][]string{v.Strings()}
}

func (v *Vote) WebPath() string {
	return "/admin/db/estimate/story/vote/" + v.StoryID.String() + "/" + v.UserID.String()
}

func (v *Vote) ToData() []any {
	return []any{v.StoryID, v.UserID, v.Choice, v.Created, v.Updated}
}

var FieldDescs = util.FieldDescs{
	{Key: "storyID", Title: "Story ID", Description: "", Type: "uuid"},
	{Key: "userID", Title: "User ID", Description: "", Type: "uuid"},
	{Key: "choice", Title: "Choice", Description: "", Type: "string"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
	{Key: "updated", Title: "Updated", Description: "", Type: "timestamp"},
}
