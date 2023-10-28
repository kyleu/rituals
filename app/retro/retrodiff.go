// Package retro - Content managed by Project Forge, see [projectforge.md] for details.
package retro

import (
	"fmt"

	"github.com/kyleu/rituals/app/util"
)

//nolint:lll,gocognit
func (r *Retro) Diff(rx *Retro) util.Diffs {
	var diffs util.Diffs
	if r.ID != rx.ID {
		diffs = append(diffs, util.NewDiff("id", r.ID.String(), rx.ID.String()))
	}
	if r.Slug != rx.Slug {
		diffs = append(diffs, util.NewDiff("slug", r.Slug, rx.Slug))
	}
	if r.Title != rx.Title {
		diffs = append(diffs, util.NewDiff("title", r.Title, rx.Title))
	}
	if r.Icon != rx.Icon {
		diffs = append(diffs, util.NewDiff("icon", r.Icon, rx.Icon))
	}
	if r.Status != rx.Status {
		diffs = append(diffs, util.NewDiff("status", r.Status.Key, rx.Status.Key))
	}
	if (r.TeamID == nil && rx.TeamID != nil) || (r.TeamID != nil && rx.TeamID == nil) || (r.TeamID != nil && rx.TeamID != nil && *r.TeamID != *rx.TeamID) {
		diffs = append(diffs, util.NewDiff("teamID", fmt.Sprint(r.TeamID), fmt.Sprint(rx.TeamID))) //nolint:gocritic // it's nullable
	}
	if (r.SprintID == nil && rx.SprintID != nil) || (r.SprintID != nil && rx.SprintID == nil) || (r.SprintID != nil && rx.SprintID != nil && *r.SprintID != *rx.SprintID) {
		diffs = append(diffs, util.NewDiff("sprintID", fmt.Sprint(r.SprintID), fmt.Sprint(rx.SprintID))) //nolint:gocritic // it's nullable
	}
	diffs = append(diffs, util.DiffObjects(r.Categories, rx.Categories, "categories")...)
	if r.Created != rx.Created {
		diffs = append(diffs, util.NewDiff("created", r.Created.String(), rx.Created.String()))
	}
	return diffs
}
