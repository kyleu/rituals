// Package rmember - Content managed by Project Forge, see [projectforge.md] for details.
package rmember

import "github.com/kyleu/rituals/app/util"

func (r *RetroMember) Diff(rx *RetroMember) util.Diffs {
	var diffs util.Diffs
	if r.RetroID != rx.RetroID {
		diffs = append(diffs, util.NewDiff("retroID", r.RetroID.String(), rx.RetroID.String()))
	}
	if r.UserID != rx.UserID {
		diffs = append(diffs, util.NewDiff("userID", r.UserID.String(), rx.UserID.String()))
	}
	if r.Name != rx.Name {
		diffs = append(diffs, util.NewDiff("name", r.Name, rx.Name))
	}
	if r.Picture != rx.Picture {
		diffs = append(diffs, util.NewDiff("picture", r.Picture, rx.Picture))
	}
	if r.Role != rx.Role {
		diffs = append(diffs, util.NewDiff("role", r.Role.Key, rx.Role.Key))
	}
	if r.Created != rx.Created {
		diffs = append(diffs, util.NewDiff("created", r.Created.String(), rx.Created.String()))
	}
	return diffs
}
