package xls

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/kyleu/rituals.dev/app/model/comment"
	"github.com/kyleu/rituals.dev/app/model/member"
	"github.com/kyleu/rituals.dev/app/util"
)

func renderCommentList(comments comment.Comments, members member.Entries, f *excelize.File) {
	key := util.Plural(util.KeyComment)
	if len(comments) > 0 {
		f.NewSheet(key)

		setColumnHeaders(key, []string{util.Title(util.KeyUser), util.Title(util.KeyContent), util.Title(util.KeyCreated)}, f)

		var data [][]interface{}
		for _, c := range comments {
			data = append(data, []interface{}{members.GetName(c.UserID), c.Content, c.Created})
		}
		setData(key, 2, data, f)
		setColumnWidths(key, []int{16, 64, 16}, f)
	}
}
