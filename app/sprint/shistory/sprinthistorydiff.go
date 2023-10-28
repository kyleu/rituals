// Package shistory - Content managed by Project Forge, see [projectforge.md] for details.
package shistory

import "github.com/kyleu/rituals/app/util"

func (s *SprintHistory) Diff(sx *SprintHistory) util.Diffs {
	var diffs util.Diffs
	if s.Slug != sx.Slug {
		diffs = append(diffs, util.NewDiff("slug", s.Slug, sx.Slug))
	}
	if s.SprintID != sx.SprintID {
		diffs = append(diffs, util.NewDiff("sprintID", s.SprintID.String(), sx.SprintID.String()))
	}
	if s.SprintName != sx.SprintName {
		diffs = append(diffs, util.NewDiff("sprintName", s.SprintName, sx.SprintName))
	}
	if s.Created != sx.Created {
		diffs = append(diffs, util.NewDiff("created", s.Created.String(), sx.Created.String()))
	}
	return diffs
}
