package xls

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/rituals.dev/app/member"
)

func renderMemberList(members member.Entries, f *excelize.File) {
	key := npncore.Plural(npncore.KeyMember)
	if len(members) > 0 {
		f.NewSheet(key)

		setColumnHeaders(key, []string{npncore.Title(npncore.KeyTitle), npncore.Title(npncore.KeyRole), npncore.Title(npncore.KeyCreated)}, f)

		var data [][]interface{}
		for _, m := range members {
			data = append(data, []interface{}{m.Name, m.Role.String(), m.Created})
		}
		setData(key, 2, data, f)
		setColumnWidths(key, []int{16, 16, 16}, f)
	}
}
