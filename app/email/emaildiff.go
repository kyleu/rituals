// Package email - Content managed by Project Forge, see [projectforge.md] for details.
package email

import "github.com/kyleu/rituals/app/util"

func (e *Email) Diff(ex *Email) util.Diffs {
	var diffs util.Diffs
	if e.ID != ex.ID {
		diffs = append(diffs, util.NewDiff("id", e.ID.String(), ex.ID.String()))
	}
	diffs = append(diffs, util.DiffObjects(e.Recipients, ex.Recipients, "recipients")...)
	if e.Subject != ex.Subject {
		diffs = append(diffs, util.NewDiff("subject", e.Subject, ex.Subject))
	}
	diffs = append(diffs, util.DiffObjects(e.Data, ex.Data, "data")...)
	if e.Plain != ex.Plain {
		diffs = append(diffs, util.NewDiff("plain", e.Plain, ex.Plain))
	}
	if e.HTML != ex.HTML {
		diffs = append(diffs, util.NewDiff("html", e.HTML, ex.HTML))
	}
	if e.UserID != ex.UserID {
		diffs = append(diffs, util.NewDiff("userID", e.UserID.String(), ex.UserID.String()))
	}
	if e.Status != ex.Status {
		diffs = append(diffs, util.NewDiff("status", e.Status, ex.Status))
	}
	if e.Created != ex.Created {
		diffs = append(diffs, util.NewDiff("created", e.Created.String(), ex.Created.String()))
	}
	return diffs
}
