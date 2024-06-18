package tpermission

import "github.com/kyleu/rituals/app/util"

func (t *TeamPermission) Diff(tx *TeamPermission) util.Diffs {
	var diffs util.Diffs
	if t.TeamID != tx.TeamID {
		diffs = append(diffs, util.NewDiff("teamID", t.TeamID.String(), tx.TeamID.String()))
	}
	if t.Key != tx.Key {
		diffs = append(diffs, util.NewDiff("key", t.Key, tx.Key))
	}
	if t.Value != tx.Value {
		diffs = append(diffs, util.NewDiff("value", t.Value, tx.Value))
	}
	if t.Access != tx.Access {
		diffs = append(diffs, util.NewDiff("access", t.Access, tx.Access))
	}
	if t.Created != tx.Created {
		diffs = append(diffs, util.NewDiff("created", t.Created.String(), tx.Created.String()))
	}
	return diffs
}
