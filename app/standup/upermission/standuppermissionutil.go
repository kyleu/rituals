package upermission

import (
	"github.com/kyleu/rituals/app/util"
)

func (s StandupPermissions) ToPermissions() util.Permissions {
	ret := make(util.Permissions, 0, len(s))
	for _, x := range s {
		ret = append(ret, &util.Permission{Key: x.Key, Value: x.Value})
	}
	return ret
}
