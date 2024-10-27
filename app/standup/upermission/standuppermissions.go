package upermission

import (
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

type StandupPermissions []*StandupPermission

func (s StandupPermissions) Get(standupID uuid.UUID, key string, value string) *StandupPermission {
	return lo.FindOrElse(s, nil, func(x *StandupPermission) bool {
		return x.StandupID == standupID && x.Key == key && x.Value == value
	})
}

func (s StandupPermissions) StandupIDs() []uuid.UUID {
	return lo.Map(s, func(xx *StandupPermission, _ int) uuid.UUID {
		return xx.StandupID
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

func (s StandupPermissions) Keys() []string {
	return lo.Map(s, func(xx *StandupPermission, _ int) string {
		return xx.Key
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

func (s StandupPermissions) Values() []string {
	return lo.Map(s, func(xx *StandupPermission, _ int) string {
		return xx.Value
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

func (s StandupPermissions) ToPKs() []*PK {
	return lo.Map(s, func(x *StandupPermission, _ int) *PK {
		return x.ToPK()
	})
}

func (s StandupPermissions) GetByStandupID(standupID uuid.UUID) StandupPermissions {
	return lo.Filter(s, func(xx *StandupPermission, _ int) bool {
		return xx.StandupID == standupID
	})
}

func (s StandupPermissions) GetByStandupIDs(standupIDs ...uuid.UUID) StandupPermissions {
	return lo.Filter(s, func(xx *StandupPermission, _ int) bool {
		return lo.Contains(standupIDs, xx.StandupID)
	})
}

func (s StandupPermissions) GetByKey(key string) StandupPermissions {
	return lo.Filter(s, func(xx *StandupPermission, _ int) bool {
		return xx.Key == key
	})
}

func (s StandupPermissions) GetByKeys(keys ...string) StandupPermissions {
	return lo.Filter(s, func(xx *StandupPermission, _ int) bool {
		return lo.Contains(keys, xx.Key)
	})
}

func (s StandupPermissions) GetByValue(value string) StandupPermissions {
	return lo.Filter(s, func(xx *StandupPermission, _ int) bool {
		return xx.Value == value
	})
}

func (s StandupPermissions) GetByValues(values ...string) StandupPermissions {
	return lo.Filter(s, func(xx *StandupPermission, _ int) bool {
		return lo.Contains(values, xx.Value)
	})
}

func (s StandupPermissions) ToCSV() ([]string, [][]string) {
	return StandupPermissionFieldDescs.Keys(), lo.Map(s, func(x *StandupPermission, _ int) []string {
		return x.Strings()
	})
}

func (s StandupPermissions) Random() *StandupPermission {
	return util.RandomElement(s)
}

func (s StandupPermissions) Clone() StandupPermissions {
	return slices.Clone(s)
}
