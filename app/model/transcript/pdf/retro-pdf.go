package pdf

import (
	"strings"

	pdfgen "github.com/johnfercher/maroto/pkg/pdf"
	"github.com/kyleu/rituals.dev/app/model/member"
	"github.com/kyleu/rituals.dev/app/model/retro"
	"github.com/kyleu/rituals.dev/app/model/transcript"
	"github.com/kyleu/rituals.dev/app/util"
)

func renderRetro(rsp transcript.RetroResponse, m pdfgen.Maroto) string {
	hr(m)
	caption(rsp.Session.Title, m)
	detailRow(util.Title(util.KeyOwner), rsp.Members.GetName(rsp.Session.Owner), m)
	detailRow(util.PluralTitle(util.KeyCategory), strings.Join(rsp.Session.Categories, ", "), m)
	if rsp.Team != nil {
		detailRow(util.SvcTeam.Title, rsp.Team.Title, m)
	}
	if rsp.Sprint != nil {
		detailRow(util.SvcSprint.Title, rsp.Sprint.Title, m)
	}
	detailRow(util.Title(util.KeyCreated), util.ToDateString(&rsp.Session.Created), m)

	renderPermissionList(rsp.Permissions, m)
	renderMemberList(rsp.Members, m)
	renderFeedbackLists(rsp.Session.Categories, rsp.Feedback, rsp.Members, m)
	renderCommentList(rsp.Comments, rsp.Members, m, true)

	return rsp.Session.Slug
}

func renderRetroList(sessions retro.Sessions, members member.Entries, m pdfgen.Maroto) {
	if len(sessions) > 0 {
		hr(m)
		caption(util.SvcRetro.PluralTitle, m)
		cols := []string{util.Title(util.KeyOwner), util.Title(util.KeyTitle), util.Title(util.KeyCreated)}
		var data [][]string
		for _, s := range sessions {
			data = append(data, []string{members.GetName(s.Owner), s.Title, util.ToDateString(&s.Created)})
		}
		table(cols, data, []uint{3, 6, 3}, m)
	}
}

func renderFeedbackLists(categories []string, feedbacks retro.Feedbacks, members member.Entries, m pdfgen.Maroto) {
	for _, c := range categories {
		var fs retro.Feedbacks
		for _, f := range feedbacks {
			if f.Category == c {
				fs = append(fs, f)
			}
		}
		if len(fs) > 0 {
			hr(m)
			caption(c, m)
			cols := []string{util.Title(util.KeyUser), util.Title(util.KeyContent), util.Title(util.KeyCreated)}
			var data [][]string
			for _, s := range fs {
				data = append(data, []string{members.GetName(s.UserID), s.Content, util.ToDateString(&s.Created)})
			}
			table(cols, data, []uint{3, 6, 3}, m)
		}
	}
}
