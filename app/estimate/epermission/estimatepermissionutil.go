package epermission

import (
	"github.com/kyleu/rituals/app/util"
	"github.com/samber/lo"
)

func (e EstimatePermissions) ToPermissions() util.Permissions {
	return lo.Map(e, func(x *EstimatePermission, _ int) *util.Permission {
		return &util.Permission{Key: x.Key, Value: x.Value}
	})
}
