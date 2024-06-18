package story

import (
	"fmt"

	"github.com/kyleu/rituals/app/util"
)

func (s *Story) Diff(sx *Story) util.Diffs {
	var diffs util.Diffs
	if s.ID != sx.ID {
		diffs = append(diffs, util.NewDiff("id", s.ID.String(), sx.ID.String()))
	}
	if s.EstimateID != sx.EstimateID {
		diffs = append(diffs, util.NewDiff("estimateID", s.EstimateID.String(), sx.EstimateID.String()))
	}
	if s.Idx != sx.Idx {
		diffs = append(diffs, util.NewDiff("idx", fmt.Sprint(s.Idx), fmt.Sprint(sx.Idx)))
	}
	if s.UserID != sx.UserID {
		diffs = append(diffs, util.NewDiff("userID", s.UserID.String(), sx.UserID.String()))
	}
	if s.Title != sx.Title {
		diffs = append(diffs, util.NewDiff("title", s.Title, sx.Title))
	}
	if s.Status != sx.Status {
		diffs = append(diffs, util.NewDiff("status", s.Status.Key, sx.Status.Key))
	}
	if s.FinalVote != sx.FinalVote {
		diffs = append(diffs, util.NewDiff("finalVote", s.FinalVote, sx.FinalVote))
	}
	if s.Created != sx.Created {
		diffs = append(diffs, util.NewDiff("created", s.Created.String(), sx.Created.String()))
	}
	return diffs
}
