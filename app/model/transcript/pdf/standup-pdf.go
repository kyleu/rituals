package pdf

import (
	pdfgen "github.com/johnfercher/maroto/pkg/pdf"
	"github.com/kyleu/rituals.dev/app/model/member"
	"github.com/kyleu/rituals.dev/app/model/standup"
	"github.com/kyleu/rituals.dev/app/model/transcript"
	"github.com/kyleu/rituals.dev/app/util"
	"sort"
)

func renderStandup(rsp transcript.StandupResponse, m pdfgen.Maroto) (string, error) {
	hr(m)
	caption(rsp.Session.Title, m)
	detailRow(util.Title(util.KeyOwner), rsp.Members.GetName(rsp.Session.Owner), m)
	if rsp.Team != nil {
		detailRow(util.SvcTeam.Title, rsp.Team.Title, m)
	}
	if rsp.Sprint != nil {
		detailRow(util.SvcSprint.Title, rsp.Sprint.Title, m)
	}
	detailRow(util.Title(util.KeyCreated), util.ToDateString(&rsp.Session.Created), m)

	var err error
	_, err = renderPermissionList(rsp.Permissions, m)
	if err != nil {
		return "", err
	}
	_, err = renderMemberList(rsp.Members, m)
	if err != nil {
		return "", err
	}
	_, err = renderReportLists(rsp.Reports, rsp.Members, m)
	if err != nil {
		return "", err
	}
	_, err = renderCommentList(rsp.Comments, rsp.Members, m, true)
	if err != nil {
		return "", err
	}

	return rsp.Session.Slug, nil
}

func renderStandupList(sessions standup.Sessions, members member.Entries, m pdfgen.Maroto) (string, error) {
	if len(sessions) > 0 {
		hr(m)
		caption(util.SvcStandup.PluralTitle, m)
		cols := []string{util.Title(util.KeyOwner), util.Title(util.KeyTitle), util.Title(util.KeyCreated)}
		var data [][]string
		for _, s := range sessions {
			data = append(data, []string{members.GetName(s.Owner), s.Title, util.ToDateString(&s.Created)})
		}
		table(cols, data, []uint{3, 6, 3}, m)
	}
	return "", nil
}

func renderReportLists(reports standup.Reports, members member.Entries, m pdfgen.Maroto) (string, error) {
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
			cols := []string{util.Title(util.KeyUser), util.Title(util.KeyContent), util.Title(util.KeyCreated)}
			var data [][]string
			for _, r := range reports {
				if r.D == day {
					data = append(data, []string{members.GetName(r.UserID), r.Content, util.ToDateString(&r.Created)})
				}
			}
			table(cols, data, []uint{3, 6, 3}, m)
		}
	}
	return "", nil
}
