package pdf

import (
	pdfgen "github.com/johnfercher/maroto/pkg/pdf"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/team"
	"github.com/kyleu/rituals.dev/app/transcript"
	"github.com/kyleu/rituals.dev/app/util"
)

func renderTeam(rsp transcript.TeamResponse, m pdfgen.Maroto) (string, error) {
	hr(m)
	caption(rsp.Session.Title, m)
	detailRow(util.Title(util.KeyOwner), rsp.Members.GetName(rsp.Session.Owner), m)
	detailRow(util.Title(util.KeyCreated), util.ToDateString(&rsp.Session.Created), m)

	renderPermissionList(rsp.Permissions, m)
	renderMemberList(rsp.Members, m)
	renderSprintList(rsp.Sprints, rsp.Members, m)
	renderEstimateList(rsp.Estimates, rsp.Members, m)
	renderStandupList(rsp.Standups, rsp.Members, m)
	renderRetroList(rsp.Retros, rsp.Members, m)
	renderCommentList(rsp.Comments, rsp.Members, m, true)

	return rsp.Session.Slug, nil
}

func renderTeamList(sessions team.Sessions, members member.Entries, m pdfgen.Maroto) {
	hr(m)
	if len(sessions) > 0 {
		caption(util.SvcTeam.PluralTitle, m)
		cols := []string{util.Title(util.KeyOwner), util.Title(util.KeyTitle), util.Title(util.KeyCreated)}
		var data [][]string
		for _, s := range sessions {
			data = append(data, []string{members.GetName(s.Owner), s.Title, util.ToDateString(&s.Created)})
		}
		table(cols, data, []uint{3, 6, 3}, m)
	}
}
