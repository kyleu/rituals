package spermission

import (
	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

type SprintPermissions []*SprintPermission

func (s SprintPermissions) Get(sprintID uuid.UUID, key string, value string) *SprintPermission {
	return lo.FindOrElse(s, nil, func(x *SprintPermission) bool {
		return x.SprintID == sprintID && x.Key == key && x.Value == value
	})
}

func (s SprintPermissions) SprintIDs() []uuid.UUID {
	return lo.Map(s, func(xx *SprintPermission, _ int) uuid.UUID {
		return xx.SprintID
	})
}

func (s SprintPermissions) SprintIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(s, func(x *SprintPermission, _ int) {
		ret = append(ret, x.SprintID.String())
	})
	return ret
}

func (s SprintPermissions) Keys() []string {
	return lo.Map(s, func(xx *SprintPermission, _ int) string {
		return xx.Key
	})
}

func (s SprintPermissions) KeyStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(s, func(x *SprintPermission, _ int) {
		ret = append(ret, x.Key)
	})
	return ret
}

func (s SprintPermissions) Values() []string {
	return lo.Map(s, func(xx *SprintPermission, _ int) string {
		return xx.Value
	})
}

func (s SprintPermissions) ValueStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(s, func(x *SprintPermission, _ int) {
		ret = append(ret, x.Value)
	})
	return ret
}

func (s SprintPermissions) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(s)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(s, func(x *SprintPermission, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (s SprintPermissions) ToPKs() []*PK {
	return lo.Map(s, func(x *SprintPermission, _ int) *PK {
		return x.ToPK()
	})
}

func (s SprintPermissions) GetBySprintID(sprintID uuid.UUID) SprintPermissions {
	return lo.Filter(s, func(xx *SprintPermission, _ int) bool {
		return xx.SprintID == sprintID
	})
}

func (s SprintPermissions) GetBySprintIDs(sprintIDs ...uuid.UUID) SprintPermissions {
	return lo.Filter(s, func(xx *SprintPermission, _ int) bool {
		return lo.Contains(sprintIDs, xx.SprintID)
	})
}

func (s SprintPermissions) GetByKey(key string) SprintPermissions {
	return lo.Filter(s, func(xx *SprintPermission, _ int) bool {
		return xx.Key == key
	})
}

func (s SprintPermissions) GetByKeys(keys ...string) SprintPermissions {
	return lo.Filter(s, func(xx *SprintPermission, _ int) bool {
		return lo.Contains(keys, xx.Key)
	})
}

func (s SprintPermissions) GetByValue(value string) SprintPermissions {
	return lo.Filter(s, func(xx *SprintPermission, _ int) bool {
		return xx.Value == value
	})
}

func (s SprintPermissions) GetByValues(values ...string) SprintPermissions {
	return lo.Filter(s, func(xx *SprintPermission, _ int) bool {
		return lo.Contains(values, xx.Value)
	})
}

func (s SprintPermissions) ToCSV() ([]string, [][]string) {
	return SprintPermissionFieldDescs.Keys(), lo.Map(s, func(x *SprintPermission, _ int) []string {
		return x.Strings()
	})
}

func (s SprintPermissions) Random() *SprintPermission {
	return util.RandomElement(s)
}

func (s SprintPermissions) Clone() SprintPermissions {
	return lo.Map(s, func(xx *SprintPermission, _ int) *SprintPermission {
		return xx.Clone()
	})
}
