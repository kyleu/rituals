package xls

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/kyleu/rituals.dev/app/permission"
)

func renderPermissionList(permissions permission.Permissions, f *excelize.File) {
	if len(permissions) > 0 {
		var data [][]interface{}
		for _, p := range permissions {
			data = append(data, []interface{}{"", p.Message()})
		}
		setData(defSheet, 8, data, f)
	}
}
