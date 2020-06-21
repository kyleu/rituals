package xls

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/kyleu/rituals.dev/app/model/permission"
)

func renderPermissionList(permissions permission.Permissions, firstRow int, f *excelize.File) {
	if len(permissions) > 0 {
		var data [][]interface{}
		for _, p := range permissions {
			data = append(data, []interface{}{"", p.Message()})
		}
		setData(defSheet, firstRow, data, f)
	}
}
