// Package vote - Content managed by Project Forge, see [projectforge.md] for details.
package vote

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

type PK struct {
	StoryID uuid.UUID `json:"storyID"`
	UserID  uuid.UUID `json:"userID"`
}

type Vote struct {
	StoryID uuid.UUID  `json:"storyID"`
	UserID  uuid.UUID  `json:"userID"`
	Choice  string     `json:"choice"`
	Created time.Time  `json:"created"`
	Updated *time.Time `json:"updated,omitempty"`
}

func New(storyID uuid.UUID, userID uuid.UUID) *Vote {
	return &Vote{StoryID: storyID, UserID: userID}
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

func FromMap(m util.ValueMap, setPK bool) (*Vote, error) {
	ret := &Vote{}
	var err error
	if setPK {
		retStoryID, e := m.ParseUUID("storyID", true, true)
		if e != nil {
			return nil, e
		}
		if retStoryID != nil {
			ret.StoryID = *retStoryID
		}
		retUserID, e := m.ParseUUID("userID", true, true)
		if e != nil {
			return nil, e
		}
		if retUserID != nil {
			ret.UserID = *retUserID
		}
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
	}
	ret.Choice, err = m.ParseString("choice", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
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

func (v *Vote) WebPath() string {
	return "/admin/db/estimate/story/vote/" + v.StoryID.String() + "/" + v.UserID.String()
}

func (v *Vote) Diff(vx *Vote) util.Diffs {
	var diffs util.Diffs
	if v.StoryID != vx.StoryID {
		diffs = append(diffs, util.NewDiff("storyID", v.StoryID.String(), vx.StoryID.String()))
	}
	if v.UserID != vx.UserID {
		diffs = append(diffs, util.NewDiff("userID", v.UserID.String(), vx.UserID.String()))
	}
	if v.Choice != vx.Choice {
		diffs = append(diffs, util.NewDiff("choice", v.Choice, vx.Choice))
	}
	if v.Created != vx.Created {
		diffs = append(diffs, util.NewDiff("created", v.Created.String(), vx.Created.String()))
	}
	return diffs
}

func (v *Vote) ToData() []any {
	return []any{v.StoryID, v.UserID, v.Choice, v.Created, v.Updated}
}
