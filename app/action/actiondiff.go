package action

import "github.com/kyleu/rituals/app/util"

func (a *Action) Diff(ax *Action) util.Diffs {
	var diffs util.Diffs
	if a.ID != ax.ID {
		diffs = append(diffs, util.NewDiff("id", a.ID.String(), ax.ID.String()))
	}
	if a.Svc != ax.Svc {
		diffs = append(diffs, util.NewDiff("svc", a.Svc.Key, ax.Svc.Key))
	}
	if a.ModelID != ax.ModelID {
		diffs = append(diffs, util.NewDiff("modelID", a.ModelID.String(), ax.ModelID.String()))
	}
	if a.UserID != ax.UserID {
		diffs = append(diffs, util.NewDiff("userID", a.UserID.String(), ax.UserID.String()))
	}
	if a.Act != ax.Act {
		diffs = append(diffs, util.NewDiff("act", a.Act, ax.Act))
	}
	diffs = append(diffs, util.DiffObjects(a.Content, ax.Content, "content")...)
	if a.Note != ax.Note {
		diffs = append(diffs, util.NewDiff("note", a.Note, ax.Note))
	}
	if a.Created != ax.Created {
		diffs = append(diffs, util.NewDiff("created", a.Created.String(), ax.Created.String()))
	}
	return diffs
}
