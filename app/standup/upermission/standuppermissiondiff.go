package upermission

import "github.com/kyleu/rituals/app/util"

func (s *StandupPermission) Diff(sx *StandupPermission) util.Diffs {
	var diffs util.Diffs
	if s.StandupID != sx.StandupID {
		diffs = append(diffs, util.NewDiff("standupID", s.StandupID.String(), sx.StandupID.String()))
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
