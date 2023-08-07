// Content managed by Project Forge, see [projectforge.md] for details.
package upermission

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type StandupPermissions []*StandupPermission

func (s StandupPermissions) Get(standupID uuid.UUID, key string, value string) *StandupPermission {
	return lo.FindOrElse(s, nil, func(x *StandupPermission) bool {
		return x.StandupID == standupID && x.Key == key && x.Value == value
	})
}

func (s StandupPermissions) ToPKs() []*PK {
	return lo.Map(s, func(x *StandupPermission, _ int) *PK {
		return x.ToPK()
	})
}

func (s StandupPermissions) GetByStandupIDs(standupIDs ...uuid.UUID) StandupPermissions {
	return lo.Filter(s, func(x *StandupPermission, _ int) bool {
		return lo.Contains(standupIDs, x.StandupID)
	})
}

func (s StandupPermissions) StandupIDs() []uuid.UUID {
	return lo.Map(s, func(x *StandupPermission, _ int) uuid.UUID {
		return x.StandupID
	})
}

func (s StandupPermissions) StandupIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(s, func(x *StandupPermission, _ int) {
		ret = append(ret, x.StandupID.String())
	})
	return ret
}

func (s StandupPermissions) GetByKeys(keys ...string) StandupPermissions {
	return lo.Filter(s, func(x *StandupPermission, _ int) bool {
		return lo.Contains(keys, x.Key)
	})
}

func (s StandupPermissions) Keys() []string {
	return lo.Map(s, func(x *StandupPermission, _ int) string {
		return x.Key
	})
}

func (s StandupPermissions) KeyStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(s, func(x *StandupPermission, _ int) {
		ret = append(ret, x.Key)
	})
	return ret
}

func (s StandupPermissions) GetByValues(values ...string) StandupPermissions {
	return lo.Filter(s, func(x *StandupPermission, _ int) bool {
		return lo.Contains(values, x.Value)
	})
}

func (s StandupPermissions) Values() []string {
	return lo.Map(s, func(x *StandupPermission, _ int) string {
		return x.Value
	})
}

func (s StandupPermissions) ValueStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(s, func(x *StandupPermission, _ int) {
		ret = append(ret, x.Value)
	})
	return ret
}

func (s StandupPermissions) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(s)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(s, func(x *StandupPermission, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (s StandupPermissions) Clone() StandupPermissions {
	return slices.Clone(s)
}
