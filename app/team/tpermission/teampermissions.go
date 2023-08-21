// Content managed by Project Forge, see [projectforge.md] for details.
package tpermission

import (
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type TeamPermissions []*TeamPermission

func (t TeamPermissions) Get(teamID uuid.UUID, key string, value string) *TeamPermission {
	return lo.FindOrElse(t, nil, func(x *TeamPermission) bool {
		return x.TeamID == teamID && x.Key == key && x.Value == value
	})
}

func (t TeamPermissions) ToPKs() []*PK {
	return lo.Map(t, func(x *TeamPermission, _ int) *PK {
		return x.ToPK()
	})
}

func (t TeamPermissions) GetByTeamIDs(teamIDs ...uuid.UUID) TeamPermissions {
	return lo.Filter(t, func(x *TeamPermission, _ int) bool {
		return lo.Contains(teamIDs, x.TeamID)
	})
}

func (t TeamPermissions) TeamIDs() []uuid.UUID {
	return lo.Map(t, func(x *TeamPermission, _ int) uuid.UUID {
		return x.TeamID
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

func (t TeamPermissions) GetByKeys(keys ...string) TeamPermissions {
	return lo.Filter(t, func(x *TeamPermission, _ int) bool {
		return lo.Contains(keys, x.Key)
	})
}

func (t TeamPermissions) Keys() []string {
	return lo.Map(t, func(x *TeamPermission, _ int) string {
		return x.Key
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

func (t TeamPermissions) GetByValues(values ...string) TeamPermissions {
	return lo.Filter(t, func(x *TeamPermission, _ int) bool {
		return lo.Contains(values, x.Value)
	})
}

func (t TeamPermissions) Values() []string {
	return lo.Map(t, func(x *TeamPermission, _ int) string {
		return x.Value
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

func (t TeamPermissions) Clone() TeamPermissions {
	return slices.Clone(t)
}
