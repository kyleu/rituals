package xls

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/kyleu/rituals.dev/app/model/member"
	"github.com/kyleu/rituals.dev/app/model/sprint"
	"github.com/kyleu/rituals.dev/app/model/transcript"
	"github.com/kyleu/rituals.dev/app/util"
)

func renderSprint(rsp transcript.SprintResponse, f *excelize.File) (string, string, error) {
	data := [][]interface{}{
		{util.Title(util.KeyTitle), rsp.Session.Title},
		{util.Title(util.KeyOwner), rsp.Members.GetName(rsp.Session.Owner)},
	}
	if rsp.Team != nil {
		data = append(data, []interface{}{util.SvcTeam.Title, rsp.Team.Title})
	}
	data = append(data, []interface{}{util.Title(util.KeyCreated), rsp.Session.Created})

	setData(defSheet, 1, data, f)
	setColumnWidths(defSheet, []int{16, 32}, f)

	renderPermissionList(rsp.Permissions, 8, f)
	renderEstimateList(rsp.Estimates, rsp.Members, f)
	renderStandupList(rsp.Standups, rsp.Members, f)
	renderRetroList(rsp.Retros, rsp.Members, f)
	renderMemberList(rsp.Members, f)
	renderCommentList(rsp.Comments, rsp.Members, f)

	return rsp.Session.Slug, util.SvcSprint.Title + " export", nil
}

func renderSprintList(sessions sprint.Sessions, members member.Entries, f *excelize.File) (string, string) {
	svc := util.SvcSprint
	if len(sessions) > 0 {
		f.NewSheet(svc.Plural)

		setColumnHeaders(svc.Plural, []string{util.Title(util.KeyTitle), util.Title(util.KeyOwner), util.Title(util.KeyCreated)}, f)

		var data [][]interface{}
		for _, s := range sessions {
			data = append(data, []interface{}{s.Title, members.GetName(s.Owner), s.Created})
		}
		setData(svc.Plural, 2, data, f)
		setColumnWidths(svc.Plural, []int{16, 16, 16}, f)
	}
	return svc.Plural, svc.Title + " export"
}
