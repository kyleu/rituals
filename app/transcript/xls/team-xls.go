package xls

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/team"
	"github.com/kyleu/rituals.dev/app/transcript"
	"github.com/kyleu/rituals.dev/app/util"
)

func renderTeam(rsp transcript.TeamResponse, f *excelize.File) (string, string, error) {
	data := [][]interface{}{
		{npncore.Title(npncore.KeyTitle), rsp.Session.Title},
		{npncore.Title(npncore.KeyOwner), rsp.Members.GetName(rsp.Session.Owner)},
		{npncore.Title(npncore.KeyCreated), rsp.Session.Created},
	}

	setData(defSheet, 1, data, f)
	setColumnWidths(defSheet, []int{16, 32}, f)

	renderPermissionList(rsp.Permissions, 8, f)
	renderSprintList(rsp.Sprints, rsp.Members, f)
	renderEstimateList(rsp.Estimates, rsp.Members, f)
	renderStandupList(rsp.Standups, rsp.Members, f)
	renderRetroList(rsp.Retros, rsp.Members, f)
	renderMemberList(rsp.Members, f)
	renderCommentList(rsp.Comments, rsp.Members, f)

	return rsp.Session.Slug, util.SvcTeam.Title + " export", nil
}

func renderTeamList(sessions team.Sessions, members member.Entries, f *excelize.File) (string, string) {
	svc := util.SvcTeam
	if len(sessions) > 0 {
		f.NewSheet(svc.Plural)

		setColumnHeaders(svc.Plural, []string{npncore.Title(npncore.KeyTitle), npncore.Title(npncore.KeyOwner), npncore.Title(npncore.KeyCreated)}, f)

		var data [][]interface{}
		for _, s := range sessions {
			data = append(data, []interface{}{s.Title, members.GetName(s.Owner), s.Created})
		}
		setData(svc.Plural, 2, data, f)
		setColumnWidths(svc.Plural, []int{16, 16, 16}, f)
	}
	return svc.Plural, svc.Title + " export"
}
