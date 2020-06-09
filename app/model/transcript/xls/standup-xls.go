package xls

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/kyleu/rituals.dev/app/model/member"
	"github.com/kyleu/rituals.dev/app/model/standup"
	"github.com/kyleu/rituals.dev/app/model/transcript"
	"github.com/kyleu/rituals.dev/app/util"
)

func renderStandup(rsp transcript.StandupResponse, f *excelize.File) (string, string, error) {
	data := [][]interface{}{
		{util.Title(util.KeyTitle), rsp.Session.Title},
		{util.Title(util.KeyOwner), rsp.Members.GetName(rsp.Session.Owner)},
	}
	if rsp.Team != nil {
		data = append(data, []interface{}{util.SvcTeam.Title, rsp.Team.Title})
	}
	if rsp.Sprint != nil {
		data = append(data, []interface{}{util.SvcSprint.Title, rsp.Sprint.Title})
	}
	data = append(data, []interface{}{util.Title(util.KeyCreated), rsp.Session.Created})

	setData(defSheet, 1, data, f)
	setColumnWidths(defSheet, []int{16, 32}, f)

	var err error
	_, _, err = renderPermissionList(rsp.Permissions, 8, f)
	if err != nil {
		return "", "", err
	}
	_, _, err = renderReportList(rsp.Reports, rsp.Members, f)
	if err != nil {
		return "", "", err
	}
	_, _, err = renderMemberList(rsp.Members, f)
	if err != nil {
		return "", "", err
	}
	_, _, err = renderCommentList(rsp.Comments, rsp.Members, f)
	if err != nil {
		return "", "", err
	}

	return rsp.Session.Slug, util.SvcStandup.Title + " export", nil
}

func renderStandupList(sessions standup.Sessions, members member.Entries, f *excelize.File) (string, string, error) {
	svc := util.SvcStandup
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
	return svc.Plural, svc.Title + " export", nil

}

func renderReportList(reports standup.Reports, members member.Entries, f *excelize.File) (string, string, error) {
	key := util.Plural(util.KeyReport)
	if len(reports) > 0 {
		f.NewSheet(key)

		setColumnHeaders(key, []string{"Day", util.Title(util.KeyUser), util.Title(util.KeyContent), util.Title(util.KeyCreated)}, f)

		var data [][]interface{}
		for _, s := range reports {
			data = append(data, []interface{}{s.D, members.GetName(s.UserID), s.Content, s.Created})
		}
		setData(key, 2, data, f)
		setColumnWidths(key, []int{16, 16, 64, 16}, f)
	}
	return key, util.Title(util.KeyStory) + " export", nil
}
