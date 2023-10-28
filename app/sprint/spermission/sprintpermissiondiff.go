// Package spermission - Content managed by Project Forge, see [projectforge.md] for details.
package spermission

import "github.com/kyleu/rituals/app/util"

func (s *SprintPermission) Diff(sx *SprintPermission) util.Diffs {
	var diffs util.Diffs
	if s.SprintID != sx.SprintID {
		diffs = append(diffs, util.NewDiff("sprintID", s.SprintID.String(), sx.SprintID.String()))
	}
	if s.Key != sx.Key {
		diffs = append(diffs, util.NewDiff("key", s.Key, sx.Key))
	}
	if s.Value != sx.Value {
		diffs = append(diffs, util.NewDiff("value", s.Value, sx.Value))
	}
	if s.Access != sx.Access {
		diffs = append(diffs, util.NewDiff("access", s.Access, sx.Access))
	}
	if s.Created != sx.Created {
		diffs = append(diffs, util.NewDiff("created", s.Created.String(), sx.Created.String()))
	}
	return diffs
}
