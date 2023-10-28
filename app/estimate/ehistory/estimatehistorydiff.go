// Package ehistory - Content managed by Project Forge, see [projectforge.md] for details.
package ehistory

import "github.com/kyleu/rituals/app/util"

func (e *EstimateHistory) Diff(ex *EstimateHistory) util.Diffs {
	var diffs util.Diffs
	if e.Slug != ex.Slug {
		diffs = append(diffs, util.NewDiff("slug", e.Slug, ex.Slug))
	}
	if e.EstimateID != ex.EstimateID {
		diffs = append(diffs, util.NewDiff("estimateID", e.EstimateID.String(), ex.EstimateID.String()))
	}
	if e.EstimateName != ex.EstimateName {
		diffs = append(diffs, util.NewDiff("estimateName", e.EstimateName, ex.EstimateName))
	}
	if e.Created != ex.Created {
		diffs = append(diffs, util.NewDiff("created", e.Created.String(), ex.Created.String()))
	}
	return diffs
}
