package estimate

import (
	"fmt"

	"github.com/kyleu/rituals/app/util"
)

//nolint:lll
func (e *Estimate) Diff(ex *Estimate) util.Diffs {
	var diffs util.Diffs
	if e.ID != ex.ID {
		diffs = append(diffs, util.NewDiff("id", e.ID.String(), ex.ID.String()))
	}
	if e.Slug != ex.Slug {
		diffs = append(diffs, util.NewDiff("slug", e.Slug, ex.Slug))
	}
	if e.Title != ex.Title {
		diffs = append(diffs, util.NewDiff("title", e.Title, ex.Title))
	}
	if e.Icon != ex.Icon {
		diffs = append(diffs, util.NewDiff("icon", e.Icon, ex.Icon))
	}
	if e.Status != ex.Status {
		diffs = append(diffs, util.NewDiff("status", e.Status.Key, ex.Status.Key))
	}
	if (e.TeamID == nil && ex.TeamID != nil) || (e.TeamID != nil && ex.TeamID == nil) || (e.TeamID != nil && ex.TeamID != nil && *e.TeamID != *ex.TeamID) {
		diffs = append(diffs, util.NewDiff("teamID", fmt.Sprint(e.TeamID), fmt.Sprint(ex.TeamID))) //nolint:gocritic // it's nullable
	}
	if (e.SprintID == nil && ex.SprintID != nil) || (e.SprintID != nil && ex.SprintID == nil) || (e.SprintID != nil && ex.SprintID != nil && *e.SprintID != *ex.SprintID) {
		diffs = append(diffs, util.NewDiff("sprintID", fmt.Sprint(e.SprintID), fmt.Sprint(ex.SprintID))) //nolint:gocritic // it's nullable
	}
	diffs = append(diffs, util.DiffObjects(e.Choices, ex.Choices, "choices")...)
	if e.Created != ex.Created {
		diffs = append(diffs, util.NewDiff("created", e.Created.String(), ex.Created.String()))
	}
	return diffs
}
