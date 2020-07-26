package pdf

import (
	pdfgen "github.com/johnfercher/maroto/pkg/pdf"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/util"
)

func renderMemberList(members member.Entries, m pdfgen.Maroto) {
	if len(members) > 0 {
		hr(m)
		cols := []string{util.Title(util.KeyMember), util.Title(util.KeyRole), util.Title(util.KeyCreated)}
		var data [][]string
		for _, c := range members {
			data = append(data, []string{c.Name, c.Role.String(), util.ToDateString(&c.Created)})
		}
		table(cols, data, []uint{3, 6, 3}, m)
	}
}
