package epermission

import (
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

func (e EstimatePermissions) ToPermissions() util.Permissions {
	return lo.Map(e, func(x *EstimatePermission, _ int) *util.Permission {
		return &util.Permission{Key: x.Key, Value: x.Value}
	})
}
