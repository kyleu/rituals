package rpermission

import "github.com/kyleu/rituals/app/util"

func (r *RetroPermission) Diff(rx *RetroPermission) util.Diffs {
	var diffs util.Diffs
	if r.RetroID != rx.RetroID {
		diffs = append(diffs, util.NewDiff("retroID", r.RetroID.String(), rx.RetroID.String()))
	}
	if r.Key != rx.Key {
		diffs = append(diffs, util.NewDiff("key", r.Key, rx.Key))
	}
	if r.Value != rx.Value {
		diffs = append(diffs, util.NewDiff("value", r.Value, rx.Value))
	}
	if r.Access != rx.Access {
		diffs = append(diffs, util.NewDiff("access", r.Access, rx.Access))
	}
	if r.Created != rx.Created {
		diffs = append(diffs, util.NewDiff("created", r.Created.String(), rx.Created.String()))
	}
	return diffs
}
