package vote

import "github.com/kyleu/rituals/app/util"

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
