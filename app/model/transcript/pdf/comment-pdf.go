package pdf

import (
	pdfgen "github.com/johnfercher/maroto/pkg/pdf"
	"github.com/kyleu/rituals.dev/app/model/comment"
	"github.com/kyleu/rituals.dev/app/model/member"
	"github.com/kyleu/rituals.dev/app/util"
)

func renderCommentList(comments comment.Comments, members member.Entries, m pdfgen.Maroto, showTitle bool) (string, error) {
	if len(comments) > 0 {
		if showTitle {
			hr(m)
			caption(util.PluralTitle(util.KeyComment), m)
		} else {
			tr(func() { td(util.PluralTitle(util.KeyComment), 12, m) }, 8, m)
		}
		cols := []string{util.Title(util.KeyUser), util.Title(util.KeyContent), util.Title(util.KeyCreated)}
		var data [][]string
		for _, c := range comments {
			data = append(data, []string{members.GetName(c.UserID), c.Content, util.ToDateString(&c.Created)})
		}
		table(cols, data, []uint{3, 6, 3}, m)
	}
	return "", nil
}
