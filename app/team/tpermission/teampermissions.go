package tpermission

import (
	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

type TeamPermissions []*TeamPermission

func (t TeamPermissions) Get(teamID uuid.UUID, key string, value string) *TeamPermission {
	return lo.FindOrElse(t, nil, func(x *TeamPermission) bool {
		return x.TeamID == teamID && x.Key == key && x.Value == value
	})
}

func (t TeamPermissions) TeamIDs() []uuid.UUID {
	return lo.Map(t, func(xx *TeamPermission, _ int) uuid.UUID {
		return xx.TeamID
	})
}

func (t TeamPermissions) TeamIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(t)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(t, func(x *TeamPermission, _ int) {
		ret = append(ret, x.TeamID.String())
	})
	return ret
}

func (t TeamPermissions) Keys() []string {
	return lo.Map(t, func(xx *TeamPermission, _ int) string {
		return xx.Key
	})
}

func (t TeamPermissions) KeyStrings(includeNil bool) []string {
	ret := make([]string, 0, len(t)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(t, func(x *TeamPermission, _ int) {
		ret = append(ret, x.Key)
	})
	return ret
}

func (t TeamPermissions) Values() []string {
	return lo.Map(t, func(xx *TeamPermission, _ int) string {
		return xx.Value
	})
}

func (t TeamPermissions) ValueStrings(includeNil bool) []string {
	ret := make([]string, 0, len(t)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(t, func(x *TeamPermission, _ int) {
		ret = append(ret, x.Value)
	})
	return ret
}

func (t TeamPermissions) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(t)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(t, func(x *TeamPermission, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (t TeamPermissions) ToPKs() []*PK {
	return lo.Map(t, func(x *TeamPermission, _ int) *PK {
		return x.ToPK()
	})
}

func (t TeamPermissions) GetByTeamID(teamID uuid.UUID) TeamPermissions {
	return lo.Filter(t, func(xx *TeamPermission, _ int) bool {
		return xx.TeamID == teamID
	})
}

func (t TeamPermissions) GetByTeamIDs(teamIDs ...uuid.UUID) TeamPermissions {
	return lo.Filter(t, func(xx *TeamPermission, _ int) bool {
		return lo.Contains(teamIDs, xx.TeamID)
	})
}

func (t TeamPermissions) GetByKey(key string) TeamPermissions {
	return lo.Filter(t, func(xx *TeamPermission, _ int) bool {
		return xx.Key == key
	})
}

func (t TeamPermissions) GetByKeys(keys ...string) TeamPermissions {
	return lo.Filter(t, func(xx *TeamPermission, _ int) bool {
		return lo.Contains(keys, xx.Key)
	})
}

func (t TeamPermissions) GetByValue(value string) TeamPermissions {
	return lo.Filter(t, func(xx *TeamPermission, _ int) bool {
		return xx.Value == value
	})
}

func (t TeamPermissions) GetByValues(values ...string) TeamPermissions {
	return lo.Filter(t, func(xx *TeamPermission, _ int) bool {
		return lo.Contains(values, xx.Value)
	})
}

func (t TeamPermissions) ToMaps() []util.ValueMap {
	return lo.Map(t, func(xx *TeamPermission, _ int) util.ValueMap {
		return xx.ToMap()
	})
}

func (t TeamPermissions) ToOrderedMaps() util.OrderedMaps[any] {
	return lo.Map(t, func(x *TeamPermission, _ int) *util.OrderedMap[any] {
		return x.ToOrderedMap()
	})
}

func (t TeamPermissions) ToCSV() ([]string, [][]string) {
	return TeamPermissionFieldDescs.Keys(), lo.Map(t, func(x *TeamPermission, _ int) []string {
		return x.Strings()
	})
}

func (t TeamPermissions) Random() *TeamPermission {
	return util.RandomElement(t)
}

func (t TeamPermissions) Clone() TeamPermissions {
	return lo.Map(t, func(xx *TeamPermission, _ int) *TeamPermission {
		return xx.Clone()
	})
}
