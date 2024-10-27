package rpermission

import (
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

type RetroPermissions []*RetroPermission

func (r RetroPermissions) Get(retroID uuid.UUID, key string, value string) *RetroPermission {
	return lo.FindOrElse(r, nil, func(x *RetroPermission) bool {
		return x.RetroID == retroID && x.Key == key && x.Value == value
	})
}

func (r RetroPermissions) RetroIDs() []uuid.UUID {
	return lo.Map(r, func(xx *RetroPermission, _ int) uuid.UUID {
		return xx.RetroID
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

func (r RetroPermissions) Keys() []string {
	return lo.Map(r, func(xx *RetroPermission, _ int) string {
		return xx.Key
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

func (r RetroPermissions) Values() []string {
	return lo.Map(r, func(xx *RetroPermission, _ int) string {
		return xx.Value
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

func (r RetroPermissions) ToPKs() []*PK {
	return lo.Map(r, func(x *RetroPermission, _ int) *PK {
		return x.ToPK()
	})
}

func (r RetroPermissions) GetByRetroID(retroID uuid.UUID) RetroPermissions {
	return lo.Filter(r, func(xx *RetroPermission, _ int) bool {
		return xx.RetroID == retroID
	})
}

func (r RetroPermissions) GetByRetroIDs(retroIDs ...uuid.UUID) RetroPermissions {
	return lo.Filter(r, func(xx *RetroPermission, _ int) bool {
		return lo.Contains(retroIDs, xx.RetroID)
	})
}

func (r RetroPermissions) GetByKey(key string) RetroPermissions {
	return lo.Filter(r, func(xx *RetroPermission, _ int) bool {
		return xx.Key == key
	})
}

func (r RetroPermissions) GetByKeys(keys ...string) RetroPermissions {
	return lo.Filter(r, func(xx *RetroPermission, _ int) bool {
		return lo.Contains(keys, xx.Key)
	})
}

func (r RetroPermissions) GetByValue(value string) RetroPermissions {
	return lo.Filter(r, func(xx *RetroPermission, _ int) bool {
		return xx.Value == value
	})
}

func (r RetroPermissions) GetByValues(values ...string) RetroPermissions {
	return lo.Filter(r, func(xx *RetroPermission, _ int) bool {
		return lo.Contains(values, xx.Value)
	})
}

func (r RetroPermissions) ToCSV() ([]string, [][]string) {
	return RetroPermissionFieldDescs.Keys(), lo.Map(r, func(x *RetroPermission, _ int) []string {
		return x.Strings()
	})
}

func (r RetroPermissions) Random() *RetroPermission {
	return util.RandomElement(r)
}

func (r RetroPermissions) Clone() RetroPermissions {
	return slices.Clone(r)
}
