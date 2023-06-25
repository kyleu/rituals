// Content managed by Project Forge, see [projectforge.md] for details.
package spermission

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type SprintPermissions []*SprintPermission

func (s SprintPermissions) Get(sprintID uuid.UUID, key string, value string) *SprintPermission {
	return lo.FindOrElse(s, nil, func(x *SprintPermission) bool {
		return x.SprintID == sprintID && x.Key == key && x.Value == value
	})
}

func (s SprintPermissions) GetBySprintIDs(sprintIDs ...uuid.UUID) SprintPermissions {
	return lo.Filter(s, func(x *SprintPermission, _ int) bool {
		return lo.Contains(sprintIDs, x.SprintID)
	})
}

func (s SprintPermissions) SprintIDs() []uuid.UUID {
	return lo.Map(s, func(x *SprintPermission, _ int) uuid.UUID {
		return x.SprintID
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

func (s SprintPermissions) GetByKeys(keys ...string) SprintPermissions {
	return lo.Filter(s, func(x *SprintPermission, _ int) bool {
		return lo.Contains(keys, x.Key)
	})
}

func (s SprintPermissions) Keys() []string {
	return lo.Map(s, func(x *SprintPermission, _ int) string {
		return x.Key
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

func (s SprintPermissions) GetByValues(values ...string) SprintPermissions {
	return lo.Filter(s, func(x *SprintPermission, _ int) bool {
		return lo.Contains(values, x.Value)
	})
}

func (s SprintPermissions) Values() []string {
	return lo.Map(s, func(x *SprintPermission, _ int) string {
		return x.Value
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

func (s SprintPermissions) Clone() SprintPermissions {
	return slices.Clone(s)
}
