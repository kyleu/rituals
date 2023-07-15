package upermission

import (
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

func (s StandupPermissions) ToPermissions() util.Permissions {
	return lo.Map(s, func(x *StandupPermission, _ int) *util.Permission {
		return &util.Permission{Key: x.Key, Value: x.Value}
	})
}
