package pdf

import (
	"github.com/kyleu/npn/npncore"
	"sort"

	pdfgen "github.com/johnfercher/maroto/pkg/pdf"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/standup"
	"github.com/kyleu/rituals.dev/app/transcript"
	"github.com/kyleu/rituals.dev/app/util"
)

func renderStandup(rsp transcript.StandupResponse, m pdfgen.Maroto) string {
	hr(m)
	caption(rsp.Session.Title, m)
	detailRow(npncore.Title(npncore.KeyOwner), rsp.Members.GetName(rsp.Session.Owner), m)
	if rsp.Team != nil {
		detailRow(util.SvcTeam.Title, rsp.Team.Title, m)
	}
	if rsp.Sprint != nil {
		detailRow(util.SvcSprint.Title, rsp.Sprint.Title, m)
	}
	detailRow(npncore.Title(npncore.KeyCreated), npncore.ToDateString(&rsp.Session.Created), m)

	renderPermissionList(rsp.Permissions, m)
	renderMemberList(rsp.Members, m)
	renderReportLists(rsp.Reports, rsp.Members, m)
	renderCommentList(rsp.Comments, rsp.Members, m, true)

	return rsp.Session.Slug
}

func renderStandupList(sessions standup.Sessions, members member.Entries, m pdfgen.Maroto) {
	if len(sessions) > 0 {
		hr(m)
		caption(util.SvcStandup.PluralTitle, m)
		cols := []string{npncore.Title(npncore.KeyOwner), npncore.Title(npncore.KeyTitle), npncore.Title(npncore.KeyCreated)}
		var data [][]string
		for _, s := range sessions {
			data = append(data, []string{members.GetName(s.Owner), s.Title, npncore.ToDateString(&s.Created)})
		}
		table(cols, data, []uint{3, 6, 3}, m)
	}
}

func renderReportLists(reports standup.Reports, members member.Entries, m pdfgen.Maroto) {
	if len(reports) > 0 {
		dayMap := make(map[string]bool)
		for _, r := range reports {
			dayMap[r.D] = true
		}
		days := make([]string, 0, len(dayMap))
		for k := range dayMap {
			days = append(days, k)
		}
		sort.Strings(days)
		for _, day := range days {
			hr(m)
			caption(day, m)
			cols := []string{npncore.Title(npncore.KeyUser), npncore.Title(npncore.KeyContent), npncore.Title(npncore.KeyCreated)}
			var data [][]string
			for _, r := range reports {
				if r.D == day {
					data = append(data, []string{members.GetName(r.UserID), r.Content, npncore.ToDateString(&r.Created)})
				}
			}
			table(cols, data, []uint{3, 6, 3}, m)
		}
	}
}
