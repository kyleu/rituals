package vote

import (
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

const DefaultRoute = "/admin/db/estimate/story/vote"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(paths...)
}

var _ svc.Model = (*Vote)(nil)

type PK struct {
	StoryID uuid.UUID `json:"storyID,omitzero"`
	UserID  uuid.UUID `json:"userID,omitzero"`
}

func (p *PK) String() string {
	return fmt.Sprintf("%v • %v", p.StoryID, p.UserID)
}

type Vote struct {
	StoryID uuid.UUID  `json:"storyID,omitzero"`
	UserID  uuid.UUID  `json:"userID,omitzero"`
	Choice  string     `json:"choice,omitzero"`
	Created time.Time  `json:"created,omitzero"`
	Updated *time.Time `json:"updated,omitzero"`
}

func NewVote(storyID uuid.UUID, userID uuid.UUID) *Vote {
	return &Vote{StoryID: storyID, UserID: userID}
}

func (v *Vote) Clone() *Vote {
	return &Vote{StoryID: v.StoryID, UserID: v.UserID, Choice: v.Choice, Created: v.Created, Updated: v.Updated}
}

func (v *Vote) String() string {
	return fmt.Sprintf("%s • %s", v.StoryID.String(), v.UserID.String())
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

func RandomVote() *Vote {
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
	return VoteFieldDescs.Keys(), [][]string{v.Strings()}
}

func (v *Vote) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(append(paths, url.QueryEscape(v.StoryID.String()), url.QueryEscape(v.UserID.String()))...)
}

func (v *Vote) Breadcrumb(extra ...string) string {
	return v.TitleString() + "||" + v.WebPath(extra...) + "**vote-yea"
}

func (v *Vote) ToData() []any {
	return []any{v.StoryID, v.UserID, v.Choice, v.Created, v.Updated}
}

var VoteFieldDescs = util.FieldDescs{
	{Key: "storyID", Title: "Story ID", Description: "", Type: "uuid"},
	{Key: "userID", Title: "User ID", Description: "", Type: "uuid"},
	{Key: "choice", Title: "Choice", Description: "", Type: "string"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
	{Key: "updated", Title: "Updated", Description: "", Type: "timestamp"},
}
