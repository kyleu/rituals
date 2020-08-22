package xls

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/kyleu/npn/npncore"
	npnxls "github.com/kyleu/npn/npnexport/xls"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/standup"
	"github.com/kyleu/rituals.dev/app/transcript"
	"github.com/kyleu/rituals.dev/app/util"
)

func renderStandup(rsp transcript.StandupResponse, f *excelize.File) (string, string, error) {
	data := [][]interface{}{
		{npncore.Title(npncore.KeyTitle), rsp.Session.Title},
		{npncore.Title(npncore.KeyOwner), rsp.Members.GetName(rsp.Session.Owner)},
	}
	if rsp.Team != nil {
		data = append(data, []interface{}{util.SvcTeam.Title, rsp.Team.Title})
	}
	if rsp.Sprint != nil {
		data = append(data, []interface{}{util.SvcSprint.Title, rsp.Sprint.Title})
	}
	data = append(data, []interface{}{npncore.Title(npncore.KeyCreated), rsp.Session.Created})

	npnxls.SetData(npnxls.DefSheet, 1, data, f)
	npnxls.SetColumnWidths(npnxls.DefSheet, []int{16, 32}, f)

	renderPermissionList(rsp.Permissions, f)
	renderReportList(rsp.Reports, rsp.Members, f)
	renderMemberList(rsp.Members, f)
	renderCommentList(rsp.Comments, rsp.Members, f)

	return rsp.Session.Slug, util.SvcStandup.Title + exportSuffix, nil
}

func renderStandupList(sessions standup.Sessions, members member.Entries, f *excelize.File) (string, string) {
	svc := util.SvcStandup
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

func renderReportList(reports standup.Reports, members member.Entries, f *excelize.File) (string, string) {
	key := npncore.Plural(npncore.KeyReport)
	if len(reports) > 0 {
		f.NewSheet(key)

		npnxls.SetColumnHeaders(key, []string{"Day", npncore.Title(npncore.KeyUser), npncore.Title(npncore.KeyContent), npncore.Title(npncore.KeyCreated)}, f)

		var data [][]interface{}
		for _, s := range reports {
			data = append(data, []interface{}{s.D, members.GetName(s.UserID), s.Content, s.Created})
		}
		npnxls.SetData(key, 2, data, f)
		npnxls.SetColumnWidths(key, []int{16, 16, 64, 16}, f)
	}
	return key, key + exportSuffix
}
