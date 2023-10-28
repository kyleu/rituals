// Package report - Content managed by Project Forge, see [projectforge.md] for details.
package report

import "github.com/kyleu/rituals/app/util"

func (r *Report) Diff(rx *Report) util.Diffs {
	var diffs util.Diffs
	if r.ID != rx.ID {
		diffs = append(diffs, util.NewDiff("id", r.ID.String(), rx.ID.String()))
	}
	if r.StandupID != rx.StandupID {
		diffs = append(diffs, util.NewDiff("standupID", r.StandupID.String(), rx.StandupID.String()))
	}
	if r.Day != rx.Day {
		diffs = append(diffs, util.NewDiff("day", r.Day.String(), rx.Day.String()))
	}
	if r.UserID != rx.UserID {
		diffs = append(diffs, util.NewDiff("userID", r.UserID.String(), rx.UserID.String()))
	}
	if r.Content != rx.Content {
		diffs = append(diffs, util.NewDiff("content", r.Content, rx.Content))
	}
	if r.HTML != rx.HTML {
		diffs = append(diffs, util.NewDiff("html", r.HTML, rx.HTML))
	}
	if r.Created != rx.Created {
		diffs = append(diffs, util.NewDiff("created", r.Created.String(), rx.Created.String()))
	}
	return diffs
}
