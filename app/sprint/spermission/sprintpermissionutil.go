package spermission

import (
	"github.com/kyleu/rituals/app/util"
	"github.com/samber/lo"
)

func (s SprintPermissions) ToPermissions() util.Permissions {
	return lo.Map(s, func(x *SprintPermission, _ int) *util.Permission {
		return &util.Permission{Key: x.Key, Value: x.Value}
	})
}
