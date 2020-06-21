package xls

import (
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/kyleu/rituals.dev/app/model/member"
	"github.com/kyleu/rituals.dev/app/model/retro"
	"github.com/kyleu/rituals.dev/app/model/transcript"
	"github.com/kyleu/rituals.dev/app/util"
)

func renderRetro(rsp transcript.RetroResponse, f *excelize.File) (string, string, error) {
	data := [][]interface{}{
		{util.Title(util.KeyTitle), rsp.Session.Title},
		{util.Title(util.KeyOwner), rsp.Members.GetName(rsp.Session.Owner)},
		{util.PluralTitle(util.KeyCategory), strings.Join(rsp.Session.Categories, ", ")},
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

	renderPermissionList(rsp.Permissions, 8, f)
	renderFeedbackList(rsp.Feedback, rsp.Members, f)
	renderMemberList(rsp.Members, f)
	renderCommentList(rsp.Comments, rsp.Members, f)

	return rsp.Session.Slug, util.SvcRetro.Title + " export", nil
}

func renderRetroList(sessions retro.Sessions, members member.Entries, f *excelize.File) (string, string) {
	svc := util.SvcRetro
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

func renderFeedbackList(feedbacks retro.Feedbacks, members member.Entries, f *excelize.File) (string, string) {
	key := util.KeyFeedback
	if len(feedbacks) > 0 {
		f.NewSheet(key)

		setColumnHeaders(key, []string{util.Title(util.KeyUser), util.Title(util.KeyCategory), util.Title(util.KeyContent), util.Title(util.KeyCreated)}, f)

		var data [][]interface{}
		for _, f := range feedbacks {
			data = append(data, []interface{}{members.GetName(f.UserID), f.Category, f.Content, f.Created})
		}
		setData(key, 2, data, f)
		setColumnWidths(key, []int{16, 16, 64, 16}, f)
	}
	return key, key + " export"
}
