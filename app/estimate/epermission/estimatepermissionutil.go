package epermission

import (
	"github.com/kyleu/rituals/app/util"
)

func (e EstimatePermissions) ToPermissions() util.Permissions {
	ret := make(util.Permissions, 0, len(e))
	for _, x := range e {
		ret = append(ret, &util.Permission{Key: x.Key, Value: x.Value})
	}
	return ret
}
