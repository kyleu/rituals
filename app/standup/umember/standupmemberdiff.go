package umember

import "github.com/kyleu/rituals/app/util"

func (s *StandupMember) Diff(sx *StandupMember) util.Diffs {
	var diffs util.Diffs
	if s.StandupID != sx.StandupID {
		diffs = append(diffs, util.NewDiff("standupID", s.StandupID.String(), sx.StandupID.String()))
	}
	if s.UserID != sx.UserID {
		diffs = append(diffs, util.NewDiff("userID", s.UserID.String(), sx.UserID.String()))
	}
	if s.Name != sx.Name {
		diffs = append(diffs, util.NewDiff("name", s.Name, sx.Name))
	}
	if s.Picture != sx.Picture {
		diffs = append(diffs, util.NewDiff("picture", s.Picture, sx.Picture))
	}
	if s.Role != sx.Role {
		diffs = append(diffs, util.NewDiff("role", s.Role.Key, sx.Role.Key))
	}
	if s.Created != sx.Created {
		diffs = append(diffs, util.NewDiff("created", s.Created.String(), sx.Created.String()))
	}
	return diffs
}
