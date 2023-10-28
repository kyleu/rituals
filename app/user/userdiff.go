// Package user - Content managed by Project Forge, see [projectforge.md] for details.
package user

import "github.com/kyleu/rituals/app/util"

func (u *User) Diff(ux *User) util.Diffs {
	var diffs util.Diffs
	if u.ID != ux.ID {
		diffs = append(diffs, util.NewDiff("id", u.ID.String(), ux.ID.String()))
	}
	if u.Name != ux.Name {
		diffs = append(diffs, util.NewDiff("name", u.Name, ux.Name))
	}
	if u.Picture != ux.Picture {
		diffs = append(diffs, util.NewDiff("picture", u.Picture, ux.Picture))
	}
	if u.Created != ux.Created {
		diffs = append(diffs, util.NewDiff("created", u.Created.String(), ux.Created.String()))
	}
	return diffs
}
