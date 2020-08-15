package xls

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/kyleu/npn/npncore"
	npnxls "github.com/kyleu/npn/npnexport/xls"
	"github.com/kyleu/rituals.dev/app/member"
)

func renderMemberList(members member.Entries, f *excelize.File) {
	key := npncore.Plural(npncore.KeyMember)
	if len(members) > 0 {
		f.NewSheet(key)

		npnxls.SetColumnHeaders(key, []string{npncore.Title(npncore.KeyTitle), npncore.Title(npncore.KeyRole), npncore.Title(npncore.KeyCreated)}, f)

		var data [][]interface{}
		for _, m := range members {
			data = append(data, []interface{}{m.Name, m.Role.String(), m.Created})
		}
		npnxls.SetData(key, 2, data, f)
		npnxls.SetColumnWidths(key, []int{16, 16, 16}, f)
	}
}
