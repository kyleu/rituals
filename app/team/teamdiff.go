// Package team - Content managed by Project Forge, see [projectforge.md] for details.
package team

import "github.com/kyleu/rituals/app/util"

func (t *Team) Diff(tx *Team) util.Diffs {
	var diffs util.Diffs
	if t.ID != tx.ID {
		diffs = append(diffs, util.NewDiff("id", t.ID.String(), tx.ID.String()))
	}
	if t.Slug != tx.Slug {
		diffs = append(diffs, util.NewDiff("slug", t.Slug, tx.Slug))
	}
	if t.Title != tx.Title {
		diffs = append(diffs, util.NewDiff("title", t.Title, tx.Title))
	}
	if t.Icon != tx.Icon {
		diffs = append(diffs, util.NewDiff("icon", t.Icon, tx.Icon))
	}
	if t.Status != tx.Status {
		diffs = append(diffs, util.NewDiff("status", t.Status.Key, tx.Status.Key))
	}
	if t.Created != tx.Created {
		diffs = append(diffs, util.NewDiff("created", t.Created.String(), tx.Created.String()))
	}
	return diffs
}
