package rpermission

import (
	"github.com/kyleu/rituals/app/util"
)

func (r RetroPermissions) ToPermissions() util.Permissions {
	ret := make(util.Permissions, 0, len(r))
	for _, x := range r {
		ret = append(ret, &util.Permission{Key: x.Key, Value: x.Value})
	}
	return ret
}
