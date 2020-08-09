package pdf

import (
	pdfgen "github.com/johnfercher/maroto/pkg/pdf"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/rituals.dev/app/comment"
	"github.com/kyleu/rituals.dev/app/member"
)

func renderCommentList(comments comment.Comments, members member.Entries, m pdfgen.Maroto, showTitle bool) {
	if len(comments) > 0 {
		if showTitle {
			hr(m)
			caption(npncore.PluralTitle(npncore.KeyComment), m)
		} else {
			tr(func() { td(npncore.PluralTitle(npncore.KeyComment), 12, m) }, 8, m)
		}
		cols := []string{npncore.Title(npncore.KeyUser), npncore.Title(npncore.KeyContent), npncore.Title(npncore.KeyCreated)}
		var data [][]string
		for _, c := range comments {
			data = append(data, []string{members.GetName(c.UserID), c.Content, npncore.ToDateString(&c.Created)})
		}
		table(cols, data, []uint{3, 6, 3}, m)
	}
}
