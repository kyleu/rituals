package uhistory

import "github.com/kyleu/rituals/app/util"

func (s *StandupHistory) Diff(sx *StandupHistory) util.Diffs {
	var diffs util.Diffs
	if s.Slug != sx.Slug {
		diffs = append(diffs, util.NewDiff("slug", s.Slug, sx.Slug))
	}
	if s.StandupID != sx.StandupID {
		diffs = append(diffs, util.NewDiff("standupID", s.StandupID.String(), sx.StandupID.String()))
	}
	if s.StandupName != sx.StandupName {
		diffs = append(diffs, util.NewDiff("standupName", s.StandupName, sx.StandupName))
	}
	if s.Created != sx.Created {
		diffs = append(diffs, util.NewDiff("created", s.Created.String(), sx.Created.String()))
	}
	return diffs
}
