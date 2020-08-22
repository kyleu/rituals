package xls

import (
	npnxls "github.com/kyleu/npn/npnexport/xls"
	"strings"

	"github.com/kyleu/npn/npncore"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/kyleu/rituals.dev/app/estimate"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/transcript"
	"github.com/kyleu/rituals.dev/app/util"
)

func renderEstimate(rsp transcript.EstimateResponse, f *excelize.File) (string, string, error) {
	data := [][]interface{}{
		{npncore.Title(npncore.KeyTitle), rsp.Session.Title},
		{npncore.Title(npncore.KeyOwner), rsp.Members.GetName(rsp.Session.Owner)},
		{npncore.PluralTitle(npncore.KeyChoice), strings.Join(rsp.Session.Choices, ", ")},
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
	renderStoryList(rsp.Stories, rsp.Members, f)
	renderMemberList(rsp.Members, f)
	renderCommentList(rsp.Comments, rsp.Members, f)

	return rsp.Session.Slug, util.SvcEstimate.Title + exportSuffix, nil
}

func renderEstimateList(sessions estimate.Sessions, members member.Entries, f *excelize.File) (string, string) {
	svc := util.SvcEstimate
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

func renderStoryList(stories estimate.Stories, members member.Entries, f *excelize.File) (string, string) {
	key := npncore.Plural(util.KeyStory)
	if len(stories) > 0 {
		f.NewSheet(key)

		npnxls.SetColumnHeaders(key, []string{npncore.Title(npncore.KeyTitle), npncore.Title(npncore.KeyUser), npncore.Title(npncore.KeyCreated)}, f)

		var data [][]interface{}
		for _, s := range stories {
			data = append(data, []interface{}{s.Title, members.GetName(s.UserID), s.Created})
		}
		npnxls.SetData(key, 2, data, f)
		npnxls.SetColumnWidths(key, []int{32, 16, 16}, f)
	}
	return key, npncore.Title(util.KeyStory) + exportSuffix
}
