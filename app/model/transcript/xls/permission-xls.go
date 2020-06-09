package xls

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/kyleu/rituals.dev/app/model/permission"
	"github.com/kyleu/rituals.dev/app/util"
)

func renderPermissionList(permissions permission.Permissions, firstRow int, f *excelize.File) (string, string, error) {
	if len(permissions) > 0 {
		var data [][]interface{}
		for _, p := range permissions {
			data = append(data, []interface{}{"", p.Message()})
		}
		setData(defSheet, firstRow, data, f)
	}
	return util.KeyPermission, util.Title(util.KeyPermission) + " export", nil
}
