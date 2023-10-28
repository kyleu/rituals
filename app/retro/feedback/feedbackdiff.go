// Package feedback - Content managed by Project Forge, see [projectforge.md] for details.
package feedback

import (
	"fmt"

	"github.com/kyleu/rituals/app/util"
)

func (f *Feedback) Diff(fx *Feedback) util.Diffs {
	var diffs util.Diffs
	if f.ID != fx.ID {
		diffs = append(diffs, util.NewDiff("id", f.ID.String(), fx.ID.String()))
	}
	if f.RetroID != fx.RetroID {
		diffs = append(diffs, util.NewDiff("retroID", f.RetroID.String(), fx.RetroID.String()))
	}
	if f.Idx != fx.Idx {
		diffs = append(diffs, util.NewDiff("idx", fmt.Sprint(f.Idx), fmt.Sprint(fx.Idx)))
	}
	if f.UserID != fx.UserID {
		diffs = append(diffs, util.NewDiff("userID", f.UserID.String(), fx.UserID.String()))
	}
	if f.Category != fx.Category {
		diffs = append(diffs, util.NewDiff("category", f.Category, fx.Category))
	}
	if f.Content != fx.Content {
		diffs = append(diffs, util.NewDiff("content", f.Content, fx.Content))
	}
	if f.HTML != fx.HTML {
		diffs = append(diffs, util.NewDiff("html", f.HTML, fx.HTML))
	}
	if f.Created != fx.Created {
		diffs = append(diffs, util.NewDiff("created", f.Created.String(), fx.Created.String()))
	}
	return diffs
}
