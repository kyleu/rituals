package epermission

import "github.com/kyleu/rituals/app/util"

func (e *EstimatePermission) Diff(ex *EstimatePermission) util.Diffs {
	var diffs util.Diffs
	if e.EstimateID != ex.EstimateID {
		diffs = append(diffs, util.NewDiff("estimateID", e.EstimateID.String(), ex.EstimateID.String()))
	}
	if e.Key != ex.Key {
		diffs = append(diffs, util.NewDiff("key", e.Key, ex.Key))
	}
	if e.Value != ex.Value {
		diffs = append(diffs, util.NewDiff("value", e.Value, ex.Value))
	}
	if e.Access != ex.Access {
		diffs = append(diffs, util.NewDiff("access", e.Access, ex.Access))
	}
	if e.Created != ex.Created {
		diffs = append(diffs, util.NewDiff("created", e.Created.String(), ex.Created.String()))
	}
	return diffs
}
