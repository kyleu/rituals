package workspace

import (
	"github.com/samber/lo"
	"strings"

	"github.com/kyleu/rituals/app/util"
)

func loadPermissionsForm(frm util.ValueMap) util.Permissions {
	var ret util.Permissions
	lo.ForEach(frm.Keys(), func(k string, _ int) {
		if strings.HasPrefix(k, "perm-") {
			k = strings.TrimPrefix(k, "perm-")
			switch k {
			case util.KeyTeam:
				ret = append(ret, &util.Permission{Key: util.KeyTeam, Value: "true"})
			case util.KeySprint:
				ret = append(ret, &util.Permission{Key: util.KeySprint, Value: "true"})
			default:
				l, r := util.StringSplit(k, '-', true)
				curr := lo.FindOrElse(ret, nil, func(x *util.Permission) bool {
					return x.Key == l
				})
				if r == "" {
					if curr == nil {
						ret = append(ret, &util.Permission{Key: l, Value: "*"})
					}
				} else {
					if curr != nil {
						curr.Value = r
					} else {
						ret = append(ret, &util.Permission{Key: l, Value: r})
					}
				}
			}
		}
	})
	return ret
}
