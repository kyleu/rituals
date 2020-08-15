package pdf

import (
	pdfgen "github.com/johnfercher/maroto/pkg/pdf"
	"github.com/kyleu/npn/npncore"
	npnpdf "github.com/kyleu/npn/npnexport/pdf"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/transcript"
	"github.com/kyleu/rituals.dev/app/util"
)

func renderSprint(rsp transcript.SprintResponse, m pdfgen.Maroto) (string, error) {
	npnpdf.HR(m)
	npnpdf.Caption(rsp.Session.Title, m)
	npnpdf.DetailRow(npncore.Title(npncore.KeyOwner), rsp.Members.GetName(rsp.Session.Owner), m)
	if rsp.Team != nil {
		npnpdf.DetailRow(util.SvcTeam.Title, rsp.Team.Title, m)
	}
	npnpdf.DetailRow(npncore.Title(npncore.KeyCreated), npncore.ToDateString(&rsp.Session.Created), m)

	renderPermissionList(rsp.Permissions, m)
	renderMemberList(rsp.Members, m)
	renderEstimateList(rsp.Estimates, rsp.Members, m)
	renderStandupList(rsp.Standups, rsp.Members, m)
	renderRetroList(rsp.Retros, rsp.Members, m)
	renderCommentList(rsp.Comments, rsp.Members, m, true)

	return rsp.Session.Slug, nil
}

func renderSprintList(sessions sprint.Sessions, members member.Entries, m pdfgen.Maroto) {
	if len(sessions) > 0 {
		npnpdf.HR(m)
		npnpdf.Caption(util.SvcSprint.PluralTitle, m)
		cols := []string{npncore.Title(npncore.KeyOwner), npncore.Title(npncore.KeyTitle), npncore.Title(npncore.KeyCreated)}
		var data [][]string
		for _, s := range sessions {
			data = append(data, []string{members.GetName(s.Owner), s.Title, npncore.ToDateString(&s.Created)})
		}
		npnpdf.Table(cols, data, []uint{3, 6, 3}, m)
	}
}
