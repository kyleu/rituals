package xls

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/kyleu/npn/npncore"
	npnxls "github.com/kyleu/npn/npnexport/xls"
	"github.com/kyleu/rituals.dev/app/comment"
	"github.com/kyleu/rituals.dev/app/member"
)

func renderCommentList(comments comment.Comments, members member.Entries, f *excelize.File) {
	key := npncore.Plural(npncore.KeyComment)
	if len(comments) > 0 {
		f.NewSheet(key)

		npnxls.SetColumnHeaders(key, []string{npncore.Title(npncore.KeyUser), npncore.Title(npncore.KeyContent), npncore.Title(npncore.KeyCreated)}, f)

		var data [][]interface{}
		for _, c := range comments {
			data = append(data, []interface{}{members.GetName(c.UserID), c.Content, c.Created})
		}
		npnxls.SetData(key, 2, data, f)
		npnxls.SetColumnWidths(key, []int{16, 64, 16}, f)
	}
}
