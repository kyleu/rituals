// Package standup - Content managed by Project Forge, see [projectforge.md] for details.
package standup

import (
	"fmt"

	"github.com/kyleu/rituals/app/util"
)

//nolint:lll,gocognit
func (s *Standup) Diff(sx *Standup) util.Diffs {
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
	if (s.SprintID == nil && sx.SprintID != nil) || (s.SprintID != nil && sx.SprintID == nil) || (s.SprintID != nil && sx.SprintID != nil && *s.SprintID != *sx.SprintID) {
		diffs = append(diffs, util.NewDiff("sprintID", fmt.Sprint(s.SprintID), fmt.Sprint(sx.SprintID))) //nolint:gocritic // it's nullable
	}
	if s.Created != sx.Created {
		diffs = append(diffs, util.NewDiff("created", s.Created.String(), sx.Created.String()))
	}
	return diffs
}
