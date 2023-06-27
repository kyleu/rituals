package upermission

import (
	"github.com/kyleu/rituals/app/util"
	"github.com/samber/lo"
)

func (s StandupPermissions) ToPermissions() util.Permissions {
	return lo.Map(s, func(x *StandupPermission, _ int) *util.Permission {
		return &util.Permission{Key: x.Key, Value: x.Value}
	})
}
