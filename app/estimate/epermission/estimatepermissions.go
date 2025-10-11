package epermission

import (
	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

type EstimatePermissions []*EstimatePermission

func (e EstimatePermissions) Get(estimateID uuid.UUID, key string, value string) *EstimatePermission {
	return lo.FindOrElse(e, nil, func(x *EstimatePermission) bool {
		return x.EstimateID == estimateID && x.Key == key && x.Value == value
	})
}

func (e EstimatePermissions) EstimateIDs() []uuid.UUID {
	return lo.Map(e, func(xx *EstimatePermission, _ int) uuid.UUID {
		return xx.EstimateID
	})
}

func (e EstimatePermissions) EstimateIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(e)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(e, func(x *EstimatePermission, _ int) {
		ret = append(ret, x.EstimateID.String())
	})
	return ret
}

func (e EstimatePermissions) Keys() []string {
	return lo.Map(e, func(xx *EstimatePermission, _ int) string {
		return xx.Key
	})
}

func (e EstimatePermissions) KeyStrings(includeNil bool) []string {
	ret := make([]string, 0, len(e)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(e, func(x *EstimatePermission, _ int) {
		ret = append(ret, x.Key)
	})
	return ret
}

func (e EstimatePermissions) Values() []string {
	return lo.Map(e, func(xx *EstimatePermission, _ int) string {
		return xx.Value
	})
}

func (e EstimatePermissions) ValueStrings(includeNil bool) []string {
	ret := make([]string, 0, len(e)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(e, func(x *EstimatePermission, _ int) {
		ret = append(ret, x.Value)
	})
	return ret
}

func (e EstimatePermissions) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(e)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(e, func(x *EstimatePermission, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (e EstimatePermissions) ToPKs() []*PK {
	return lo.Map(e, func(x *EstimatePermission, _ int) *PK {
		return x.ToPK()
	})
}

func (e EstimatePermissions) GetByEstimateID(estimateID uuid.UUID) EstimatePermissions {
	return lo.Filter(e, func(xx *EstimatePermission, _ int) bool {
		return xx.EstimateID == estimateID
	})
}

func (e EstimatePermissions) GetByEstimateIDs(estimateIDs ...uuid.UUID) EstimatePermissions {
	return lo.Filter(e, func(xx *EstimatePermission, _ int) bool {
		return lo.Contains(estimateIDs, xx.EstimateID)
	})
}

func (e EstimatePermissions) GetByKey(key string) EstimatePermissions {
	return lo.Filter(e, func(xx *EstimatePermission, _ int) bool {
		return xx.Key == key
	})
}

func (e EstimatePermissions) GetByKeys(keys ...string) EstimatePermissions {
	return lo.Filter(e, func(xx *EstimatePermission, _ int) bool {
		return lo.Contains(keys, xx.Key)
	})
}

func (e EstimatePermissions) GetByValue(value string) EstimatePermissions {
	return lo.Filter(e, func(xx *EstimatePermission, _ int) bool {
		return xx.Value == value
	})
}

func (e EstimatePermissions) GetByValues(values ...string) EstimatePermissions {
	return lo.Filter(e, func(xx *EstimatePermission, _ int) bool {
		return lo.Contains(values, xx.Value)
	})
}

func (e EstimatePermissions) ToMaps() []util.ValueMap {
	return lo.Map(e, func(xx *EstimatePermission, _ int) util.ValueMap {
		return xx.ToMap()
	})
}

func (e EstimatePermissions) ToOrderedMaps() util.OrderedMaps[any] {
	return lo.Map(e, func(x *EstimatePermission, _ int) *util.OrderedMap[any] {
		return x.ToOrderedMap()
	})
}

func (e EstimatePermissions) ToCSV() ([]string, [][]string) {
	return EstimatePermissionFieldDescs.Keys(), lo.Map(e, func(x *EstimatePermission, _ int) []string {
		return x.Strings()
	})
}

func (e EstimatePermissions) Random() *EstimatePermission {
	return util.RandomElement(e)
}

func (e EstimatePermissions) Clone() EstimatePermissions {
	return lo.Map(e, func(xx *EstimatePermission, _ int) *EstimatePermission {
		return xx.Clone()
	})
}
