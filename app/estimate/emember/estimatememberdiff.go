// Package emember - Content managed by Project Forge, see [projectforge.md] for details.
package emember

import "github.com/kyleu/rituals/app/util"

func (e *EstimateMember) Diff(ex *EstimateMember) util.Diffs {
	var diffs util.Diffs
	if e.EstimateID != ex.EstimateID {
		diffs = append(diffs, util.NewDiff("estimateID", e.EstimateID.String(), ex.EstimateID.String()))
	}
	if e.UserID != ex.UserID {
		diffs = append(diffs, util.NewDiff("userID", e.UserID.String(), ex.UserID.String()))
	}
	if e.Name != ex.Name {
		diffs = append(diffs, util.NewDiff("name", e.Name, ex.Name))
	}
	if e.Picture != ex.Picture {
		diffs = append(diffs, util.NewDiff("picture", e.Picture, ex.Picture))
	}
	if e.Role != ex.Role {
		diffs = append(diffs, util.NewDiff("role", e.Role.Key, ex.Role.Key))
	}
	if e.Created != ex.Created {
		diffs = append(diffs, util.NewDiff("created", e.Created.String(), ex.Created.String()))
	}
	return diffs
}
