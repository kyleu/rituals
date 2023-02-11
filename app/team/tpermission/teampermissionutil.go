package tpermission

import (
	"github.com/kyleu/rituals/app/util"
)

func (t TeamPermissions) ToPermissions() util.Permissions {
	ret := make(util.Permissions, 0, len(t))
	for _, x := range t {
		ret = append(ret, &util.Permission{Key: x.Key, Value: x.Value})
	}
	return ret
}
