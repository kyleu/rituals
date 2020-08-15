package pdf

import (
	pdfgen "github.com/johnfercher/maroto/pkg/pdf"
	"github.com/kyleu/npn/npncore"
	npnpdf "github.com/kyleu/npn/npnexport/pdf"
	"github.com/kyleu/rituals.dev/app/member"
)

func renderMemberList(members member.Entries, m pdfgen.Maroto) {
	if len(members) > 0 {
		npnpdf.HR(m)
		cols := []string{npncore.Title(npncore.KeyMember), npncore.Title(npncore.KeyRole), npncore.Title(npncore.KeyCreated)}
		var data [][]string
		for _, c := range members {
			data = append(data, []string{c.Name, c.Role.String(), npncore.ToDateString(&c.Created)})
		}
		npnpdf.Table(cols, data, []uint{3, 6, 3}, m)
	}
}
