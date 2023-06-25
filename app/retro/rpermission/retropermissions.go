// Content managed by Project Forge, see [projectforge.md] for details.
package rpermission

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type RetroPermissions []*RetroPermission

func (r RetroPermissions) Get(retroID uuid.UUID, key string, value string) *RetroPermission {
	return lo.FindOrElse(r, nil, func(x *RetroPermission) bool {
		return x.RetroID == retroID && x.Key == key && x.Value == value
	})
}

func (r RetroPermissions) GetByRetroIDs(retroIDs ...uuid.UUID) RetroPermissions {
	return lo.Filter(r, func(x *RetroPermission, _ int) bool {
		return lo.Contains(retroIDs, x.RetroID)
	})
}

func (r RetroPermissions) RetroIDs() []uuid.UUID {
	return lo.Map(r, func(x *RetroPermission, _ int) uuid.UUID {
		return x.RetroID
	})
}

func (r RetroPermissions) RetroIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(r)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(r, func(x *RetroPermission, _ int) {
		ret = append(ret, x.RetroID.String())
	})
	return ret
}

func (r RetroPermissions) GetByKeys(keys ...string) RetroPermissions {
	return lo.Filter(r, func(x *RetroPermission, _ int) bool {
		return lo.Contains(keys, x.Key)
	})
}

func (r RetroPermissions) Keys() []string {
	return lo.Map(r, func(x *RetroPermission, _ int) string {
		return x.Key
	})
}

func (r RetroPermissions) KeyStrings(includeNil bool) []string {
	ret := make([]string, 0, len(r)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(r, func(x *RetroPermission, _ int) {
		ret = append(ret, x.Key)
	})
	return ret
}

func (r RetroPermissions) GetByValues(values ...string) RetroPermissions {
	return lo.Filter(r, func(x *RetroPermission, _ int) bool {
		return lo.Contains(values, x.Value)
	})
}

func (r RetroPermissions) Values() []string {
	return lo.Map(r, func(x *RetroPermission, _ int) string {
		return x.Value
	})
}

func (r RetroPermissions) ValueStrings(includeNil bool) []string {
	ret := make([]string, 0, len(r)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(r, func(x *RetroPermission, _ int) {
		ret = append(ret, x.Value)
	})
	return ret
}

func (r RetroPermissions) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(r)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(r, func(x *RetroPermission, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (r RetroPermissions) Clone() RetroPermissions {
	return slices.Clone(r)
}
