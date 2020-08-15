package pdf

import (
	pdfgen "github.com/johnfercher/maroto/pkg/pdf"
	"github.com/kyleu/npn/npncore"
	npnpdf "github.com/kyleu/npn/npnexport/pdf"
	"github.com/kyleu/rituals.dev/app/comment"
	"github.com/kyleu/rituals.dev/app/member"
)

func renderCommentList(comments comment.Comments, members member.Entries, m pdfgen.Maroto, showTitle bool) {
	if len(comments) > 0 {
		if showTitle {
			npnpdf.HR(m)
			npnpdf.Caption(npncore.PluralTitle(npncore.KeyComment), m)
		} else {
			npnpdf.TR(func() { npnpdf.TD(npncore.PluralTitle(npncore.KeyComment), 12, m) }, 8, m)
		}
		cols := []string{npncore.Title(npncore.KeyUser), npncore.Title(npncore.KeyContent), npncore.Title(npncore.KeyCreated)}
		var data [][]string
		for _, c := range comments {
			data = append(data, []string{members.GetName(c.UserID), c.Content, npncore.ToDateString(&c.Created)})
		}
		npnpdf.Table(cols, data, []uint{3, 6, 3}, m)
	}
}
