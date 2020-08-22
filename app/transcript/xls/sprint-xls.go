package xls

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/kyleu/npn/npncore"
	npnxls "github.com/kyleu/npn/npnexport/xls"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/transcript"
	"github.com/kyleu/rituals.dev/app/util"
)

func renderSprint(rsp transcript.SprintResponse, f *excelize.File) (string, string, error) {
	data := [][]interface{}{
		{npncore.Title(npncore.KeyTitle), rsp.Session.Title},
		{npncore.Title(npncore.KeyOwner), rsp.Members.GetName(rsp.Session.Owner)},
	}
	if rsp.Team != nil {
		data = append(data, []interface{}{util.SvcTeam.Title, rsp.Team.Title})
	}
	data = append(data, []interface{}{npncore.Title(npncore.KeyCreated), rsp.Session.Created})

	npnxls.SetData(npnxls.DefSheet, 1, data, f)
	npnxls.SetColumnWidths(npnxls.DefSheet, []int{16, 32}, f)

	renderPermissionList(rsp.Permissions, f)
	renderEstimateList(rsp.Estimates, rsp.Members, f)
	renderStandupList(rsp.Standups, rsp.Members, f)
	renderRetroList(rsp.Retros, rsp.Members, f)
	renderMemberList(rsp.Members, f)
	renderCommentList(rsp.Comments, rsp.Members, f)

	return rsp.Session.Slug, util.SvcSprint.Title + exportSuffix, nil
}

func renderSprintList(sessions sprint.Sessions, members member.Entries, f *excelize.File) (string, string) {
	svc := util.SvcSprint
	if len(sessions) > 0 {
		f.NewSheet(svc.Plural)

		npnxls.SetColumnHeaders(svc.Plural, []string{npncore.Title(npncore.KeyTitle), npncore.Title(npncore.KeyOwner), npncore.Title(npncore.KeyCreated)}, f)

		var data [][]interface{}
		for _, s := range sessions {
			data = append(data, []interface{}{s.Title, members.GetName(s.Owner), s.Created})
		}
		npnxls.SetData(svc.Plural, 2, data, f)
		npnxls.SetColumnWidths(svc.Plural, []int{16, 16, 16}, f)
	}
	return svc.Plural, svc.Title + exportSuffix
}
