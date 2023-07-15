package spermission

import (
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

func (s SprintPermissions) ToPermissions() util.Permissions {
	return lo.Map(s, func(x *SprintPermission, _ int) *util.Permission {
		return &util.Permission{Key: x.Key, Value: x.Value}
	})
}
