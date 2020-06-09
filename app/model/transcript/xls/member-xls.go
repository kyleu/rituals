package xls

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/kyleu/rituals.dev/app/model/member"
	"github.com/kyleu/rituals.dev/app/util"
)

func renderMemberList(members member.Entries, f *excelize.File) (string, string, error) {
	key := util.Plural(util.KeyMember)
	if len(members) > 0 {
		f.NewSheet(key)

		setColumnHeaders(key, []string{util.Title(util.KeyTitle), util.Title(util.KeyRole), util.Title(util.KeyCreated)}, f)

		var data [][]interface{}
		for _, m := range members {
			data = append(data, []interface{}{m.Name, m.Role.String(), m.Created})
		}
		setData(key, 2, data, f)
		setColumnWidths(key, []int{16, 16, 16}, f)
	}
	return key, util.Title(util.KeyMember) + " export", nil
}
