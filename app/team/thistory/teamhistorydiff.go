// Package thistory - Content managed by Project Forge, see [projectforge.md] for details.
package thistory

import "github.com/kyleu/rituals/app/util"

func (t *TeamHistory) Diff(tx *TeamHistory) util.Diffs {
	var diffs util.Diffs
	if t.Slug != tx.Slug {
		diffs = append(diffs, util.NewDiff("slug", t.Slug, tx.Slug))
	}
	if t.TeamID != tx.TeamID {
		diffs = append(diffs, util.NewDiff("teamID", t.TeamID.String(), tx.TeamID.String()))
	}
	if t.TeamName != tx.TeamName {
		diffs = append(diffs, util.NewDiff("teamName", t.TeamName, tx.TeamName))
	}
	if t.Created != tx.Created {
		diffs = append(diffs, util.NewDiff("created", t.Created.String(), tx.Created.String()))
	}
	return diffs
}
