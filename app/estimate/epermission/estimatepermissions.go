// Content managed by Project Forge, see [projectforge.md] for details.
package epermission

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type EstimatePermissions []*EstimatePermission

func (e EstimatePermissions) Get(estimateID uuid.UUID, key string, value string) *EstimatePermission {
	return lo.FindOrElse(e, nil, func(x *EstimatePermission) bool {
		return x.EstimateID == estimateID && x.Key == key && x.Value == value
	})
}

func (e EstimatePermissions) GetByEstimateIDs(estimateIDs ...uuid.UUID) EstimatePermissions {
	return lo.Filter(e, func(x *EstimatePermission, _ int) bool {
		return lo.Contains(estimateIDs, x.EstimateID)
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

func (e EstimatePermissions) GetByKeys(keys ...string) EstimatePermissions {
	return lo.Filter(e, func(x *EstimatePermission, _ int) bool {
		return lo.Contains(keys, x.Key)
	})
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

func (e EstimatePermissions) GetByValues(values ...string) EstimatePermissions {
	return lo.Filter(e, func(x *EstimatePermission, _ int) bool {
		return lo.Contains(values, x.Value)
	})
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
