package sprint

import (
	"fmt"

	"github.com/kyleu/rituals/app/util"
)

//nolint:lll
func (s *Sprint) Diff(sx *Sprint) util.Diffs {
	var diffs util.Diffs
	if s.ID != sx.ID {
		diffs = append(diffs, util.NewDiff("id", s.ID.String(), sx.ID.String()))
	}
	if s.Slug != sx.Slug {
		diffs = append(diffs, util.NewDiff("slug", s.Slug, sx.Slug))
	}
	if s.Title != sx.Title {
		diffs = append(diffs, util.NewDiff("title", s.Title, sx.Title))
	}
	if s.Icon != sx.Icon {
		diffs = append(diffs, util.NewDiff("icon", s.Icon, sx.Icon))
	}
	if s.Status != sx.Status {
		diffs = append(diffs, util.NewDiff("status", s.Status.Key, sx.Status.Key))
	}
	if (s.TeamID == nil && sx.TeamID != nil) || (s.TeamID != nil && sx.TeamID == nil) || (s.TeamID != nil && sx.TeamID != nil && *s.TeamID != *sx.TeamID) {
		diffs = append(diffs, util.NewDiff("teamID", fmt.Sprint(s.TeamID), fmt.Sprint(sx.TeamID))) //nolint:gocritic // it's nullable
	}
	if (s.StartDate == nil && sx.StartDate != nil) || (s.StartDate != nil && sx.StartDate == nil) || (s.StartDate != nil && sx.StartDate != nil && *s.StartDate != *sx.StartDate) {
		diffs = append(diffs, util.NewDiff("startDate", fmt.Sprint(s.StartDate), fmt.Sprint(sx.StartDate))) //nolint:gocritic // it's nullable
	}
	if (s.EndDate == nil && sx.EndDate != nil) || (s.EndDate != nil && sx.EndDate == nil) || (s.EndDate != nil && sx.EndDate != nil && *s.EndDate != *sx.EndDate) {
		diffs = append(diffs, util.NewDiff("endDate", fmt.Sprint(s.EndDate), fmt.Sprint(sx.EndDate))) //nolint:gocritic // it's nullable
	}
	if s.Created != sx.Created {
		diffs = append(diffs, util.NewDiff("created", s.Created.String(), sx.Created.String()))
	}
	return diffs
}
