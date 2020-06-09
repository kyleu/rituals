package pdf

import (
	pdfgen "github.com/johnfercher/maroto/pkg/pdf"
	"github.com/kyleu/rituals.dev/app/model/member"
	"github.com/kyleu/rituals.dev/app/model/sprint"
	"github.com/kyleu/rituals.dev/app/model/transcript"
	"github.com/kyleu/rituals.dev/app/util"
)

func renderSprint(rsp transcript.SprintResponse, m pdfgen.Maroto) (string, error) {
	hr(m)
	caption(rsp.Session.Title, m)
	detailRow(util.Title(util.KeyOwner), rsp.Members.GetName(rsp.Session.Owner), m)
	if rsp.Team != nil {
		detailRow(util.SvcTeam.Title, rsp.Team.Title, m)
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
	_, err = renderEstimateList(rsp.Estimates, rsp.Members, m)
	if err != nil {
		return "", err
	}
	_, err = renderStandupList(rsp.Standups, rsp.Members, m)
	if err != nil {
		return "", err
	}
	_, err = renderRetroList(rsp.Retros, rsp.Members, m)
	if err != nil {
		return "", err
	}
	_, err = renderCommentList(rsp.Comments, rsp.Members, m, true)
	if err != nil {
		return "", err
	}

	return rsp.Session.Slug, nil
}

func renderSprintList(sessions sprint.Sessions, members member.Entries, m pdfgen.Maroto) (string, error) {
	if len(sessions) > 0 {
		hr(m)
		caption(util.SvcSprint.PluralTitle, m)
		cols := []string{util.Title(util.KeyOwner), util.Title(util.KeyTitle), util.Title(util.KeyCreated)}
		var data [][]string
		for _, s := range sessions {
			data = append(data, []string{members.GetName(s.Owner), s.Title, util.ToDateString(&s.Created)})
		}
		table(cols, data, []uint{3, 6, 3}, m)
	}
	return "", nil
}
