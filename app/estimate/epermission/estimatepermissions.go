// Package epermission - Content managed by Project Forge, see [projectforge.md] for details.
package epermission

import (
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type EstimatePermissions []*EstimatePermission

func (e EstimatePermissions) Get(estimateID uuid.UUID, key string, value string) *EstimatePermission {
	return lo.FindOrElse(e, nil, func(x *EstimatePermission) bool {
		return x.EstimateID == estimateID && x.Key == key && x.Value == value
	})
}

func (e EstimatePermissions) ToPKs() []*PK {
	return lo.Map(e, func(x *EstimatePermission, _ int) *PK {
		return x.ToPK()
	})
}

func (e EstimatePermissions) GetByEstimateIDs(estimateIDs ...uuid.UUID) EstimatePermissions {
	return lo.Filter(e, func(xx *EstimatePermission, _ int) bool {
		return lo.Contains(estimateIDs, xx.EstimateID)
	})
}

func (e EstimatePermissions) GetByEstimateID(estimateID uuid.UUID) EstimatePermissions {
	return lo.Filter(e, func(xx *EstimatePermission, _ int) bool {
		return xx.EstimateID == estimateID
	})
}

func (e EstimatePermissions) GetByKeys(keys ...string) EstimatePermissions {
	return lo.Filter(e, func(xx *EstimatePermission, _ int) bool {
		return lo.Contains(keys, xx.Key)
	})
}

func (e EstimatePermissions) GetByKey(key string) EstimatePermissions {
	return lo.Filter(e, func(xx *EstimatePermission, _ int) bool {
		return xx.Key == key
	})
}

func (e EstimatePermissions) GetByValues(values ...string) EstimatePermissions {
	return lo.Filter(e, func(xx *EstimatePermission, _ int) bool {
		return lo.Contains(values, xx.Value)
	})
}

func (e EstimatePermissions) GetByValue(value string) EstimatePermissions {
	return lo.Filter(e, func(xx *EstimatePermission, _ int) bool {
		return xx.Value == value
	})
}

func (e EstimatePermissions) EstimateIDs() []uuid.UUID {
	return lo.Map(e, func(x *EstimatePermission, _ int) uuid.UUID {
		return x.EstimateID
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
	return lo.Map(e, func(x *EstimatePermission, _ int) string {
		return x.Key
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
	return lo.Map(e, func(x *EstimatePermission, _ int) string {
		return x.Value
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

func (e EstimatePermissions) Clone() EstimatePermissions {
	return slices.Clone(e)
}
