// Package tmember - Content managed by Project Forge, see [projectforge.md] for details.
package tmember

import "github.com/kyleu/rituals/app/util"

func (t *TeamMember) Diff(tx *TeamMember) util.Diffs {
	var diffs util.Diffs
	if t.TeamID != tx.TeamID {
		diffs = append(diffs, util.NewDiff("teamID", t.TeamID.String(), tx.TeamID.String()))
	}
	if t.UserID != tx.UserID {
		diffs = append(diffs, util.NewDiff("userID", t.UserID.String(), tx.UserID.String()))
	}
	if t.Name != tx.Name {
		diffs = append(diffs, util.NewDiff("name", t.Name, tx.Name))
	}
	if t.Picture != tx.Picture {
		diffs = append(diffs, util.NewDiff("picture", t.Picture, tx.Picture))
	}
	if t.Role != tx.Role {
		diffs = append(diffs, util.NewDiff("role", t.Role.Key, tx.Role.Key))
	}
	if t.Created != tx.Created {
		diffs = append(diffs, util.NewDiff("created", t.Created.String(), tx.Created.String()))
	}
	return diffs
}
