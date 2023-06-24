// Content managed by Project Forge, see [projectforge.md] for details.
package spermission

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type SprintPermissions []*SprintPermission

func (s SprintPermissions) Get(sprintID uuid.UUID, key string, value string) *SprintPermission {
	for _, x := range s {
		if x.SprintID == sprintID && x.Key == key && x.Value == value {
			return x
		}
	}
	return nil
}

func (s SprintPermissions) GetBySprintIDs(sprintIDs ...uuid.UUID) SprintPermissions {
	var ret SprintPermissions
	for _, x := range s {
		if lo.Contains(sprintIDs, x.SprintID) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (s SprintPermissions) SprintIDs() []uuid.UUID {
	ret := make([]uuid.UUID, 0, len(s)+1)
	for _, x := range s {
		ret = append(ret, x.SprintID)
	}
	return ret
}

func (s SprintPermissions) SprintIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range s {
		ret = append(ret, x.SprintID.String())
	}
	return ret
}

func (s SprintPermissions) GetByKeys(keys ...string) SprintPermissions {
	var ret SprintPermissions
	for _, x := range s {
		if lo.Contains(keys, x.Key) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (s SprintPermissions) Keys() []string {
	ret := make([]string, 0, len(s)+1)
	for _, x := range s {
		ret = append(ret, x.Key)
	}
	return ret
}

func (s SprintPermissions) KeyStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range s {
		ret = append(ret, x.Key)
	}
	return ret
}

func (s SprintPermissions) GetByValues(values ...string) SprintPermissions {
	var ret SprintPermissions
	for _, x := range s {
		if lo.Contains(values, x.Value) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (s SprintPermissions) Values() []string {
	ret := make([]string, 0, len(s)+1)
	for _, x := range s {
		ret = append(ret, x.Value)
	}
	return ret
}

func (s SprintPermissions) ValueStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range s {
		ret = append(ret, x.Value)
	}
	return ret
}

func (s SprintPermissions) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(s)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range s {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (s SprintPermissions) Clone() SprintPermissions {
	return slices.Clone(s)
}
