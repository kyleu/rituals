// Package rhistory - Content managed by Project Forge, see [projectforge.md] for details.
package rhistory

import "github.com/kyleu/rituals/app/util"

func (r *RetroHistory) Diff(rx *RetroHistory) util.Diffs {
	var diffs util.Diffs
	if r.Slug != rx.Slug {
		diffs = append(diffs, util.NewDiff("slug", r.Slug, rx.Slug))
	}
	if r.RetroID != rx.RetroID {
		diffs = append(diffs, util.NewDiff("retroID", r.RetroID.String(), rx.RetroID.String()))
	}
	if r.RetroName != rx.RetroName {
		diffs = append(diffs, util.NewDiff("retroName", r.RetroName, rx.RetroName))
	}
	if r.Created != rx.Created {
		diffs = append(diffs, util.NewDiff("created", r.Created.String(), rx.Created.String()))
	}
	return diffs
}
