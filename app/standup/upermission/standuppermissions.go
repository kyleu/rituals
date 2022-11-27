// Content managed by Project Forge, see [projectforge.md] for details.
package upermission

import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type StandupPermissions []*StandupPermission

func (s StandupPermissions) Get(standupID uuid.UUID, key string, value string) *StandupPermission {
	for _, x := range s {
		if x.StandupID == standupID && x.Key == key && x.Value == value {
			return x
		}
	}
	return nil
}

func (s StandupPermissions) GetByStandupIDs(standupIDs ...uuid.UUID) StandupPermissions {
	var ret StandupPermissions
	for _, x := range s {
		if slices.Contains(standupIDs, x.StandupID) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (s StandupPermissions) StandupIDs() []uuid.UUID {
	ret := make([]uuid.UUID, 0, len(s)+1)
	for _, x := range s {
		ret = append(ret, x.StandupID)
	}
	return ret
}

func (s StandupPermissions) StandupIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range s {
		ret = append(ret, x.StandupID.String())
	}
	return ret
}

func (s StandupPermissions) GetByKeys(keys ...string) StandupPermissions {
	var ret StandupPermissions
	for _, x := range s {
		if slices.Contains(keys, x.Key) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (s StandupPermissions) Keys() []string {
	ret := make([]string, 0, len(s)+1)
	for _, x := range s {
		ret = append(ret, x.Key)
	}
	return ret
}

func (s StandupPermissions) KeyStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range s {
		ret = append(ret, x.Key)
	}
	return ret
}

func (s StandupPermissions) GetByValues(values ...string) StandupPermissions {
	var ret StandupPermissions
	for _, x := range s {
		if slices.Contains(values, x.Value) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (s StandupPermissions) Values() []string {
	ret := make([]string, 0, len(s)+1)
	for _, x := range s {
		ret = append(ret, x.Value)
	}
	return ret
}

func (s StandupPermissions) ValueStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range s {
		ret = append(ret, x.Value)
	}
	return ret
}

func (s StandupPermissions) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(s)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range s {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (s StandupPermissions) Clone() StandupPermissions {
	return slices.Clone(s)
}
