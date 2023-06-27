package tpermission

import (
	"github.com/kyleu/rituals/app/util"
	"github.com/samber/lo"
)

func (t TeamPermissions) ToPermissions() util.Permissions {
	return lo.Map(t, func(x *TeamPermission, _ int) *util.Permission {
		return &util.Permission{Key: x.Key, Value: x.Value}
	})
}
