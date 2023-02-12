package workspace

import (
	"strings"

	"github.com/kyleu/rituals/app/util"
)

func loadPermissionsForm(frm util.ValueMap) util.Permissions {
	var ret util.Permissions
	for k := range frm {
		if strings.HasPrefix(k, "perm-") {
			k = strings.TrimPrefix(k, "perm-")
			if k == util.KeyTeam {
				ret = append(ret, &util.Permission{Key: util.KeyTeam, Value: "true"})
			} else if k == util.KeySprint {
				ret = append(ret, &util.Permission{Key: util.KeySprint, Value: "true"})
			} else {
				l, r := util.StringSplit(k, '-', true)
				var curr *util.Permission
				for _, x := range ret {
					if x.Key == l {
						curr = x
						break
					}
				}
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
	}
	return ret
}
