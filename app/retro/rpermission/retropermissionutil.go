package rpermission

import (
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

func (r RetroPermissions) ToPermissions() util.Permissions {
	return lo.Map(r, func(x *RetroPermission, _ int) *util.Permission {
		return &util.Permission{Key: x.Key, Value: x.Value}
	})
}
