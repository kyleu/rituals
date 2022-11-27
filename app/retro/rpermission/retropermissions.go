// Content managed by Project Forge, see [projectforge.md] for details.
package rpermission

import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type RetroPermissions []*RetroPermission

func (r RetroPermissions) Get(retroID uuid.UUID, key string, value string) *RetroPermission {
	for _, x := range r {
		if x.RetroID == retroID && x.Key == key && x.Value == value {
			return x
		}
	}
	return nil
}

func (r RetroPermissions) GetByRetroIDs(retroIDs ...uuid.UUID) RetroPermissions {
	var ret RetroPermissions
	for _, x := range r {
		if slices.Contains(retroIDs, x.RetroID) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (r RetroPermissions) RetroIDs() []uuid.UUID {
	ret := make([]uuid.UUID, 0, len(r)+1)
	for _, x := range r {
		ret = append(ret, x.RetroID)
	}
	return ret
}

func (r RetroPermissions) RetroIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(r)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range r {
		ret = append(ret, x.RetroID.String())
	}
	return ret
}

func (r RetroPermissions) GetByKeys(keys ...string) RetroPermissions {
	var ret RetroPermissions
	for _, x := range r {
		if slices.Contains(keys, x.Key) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (r RetroPermissions) Keys() []string {
	ret := make([]string, 0, len(r)+1)
	for _, x := range r {
		ret = append(ret, x.Key)
	}
	return ret
}

func (r RetroPermissions) KeyStrings(includeNil bool) []string {
	ret := make([]string, 0, len(r)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range r {
		ret = append(ret, x.Key)
	}
	return ret
}

func (r RetroPermissions) GetByValues(values ...string) RetroPermissions {
	var ret RetroPermissions
	for _, x := range r {
		if slices.Contains(values, x.Value) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (r RetroPermissions) Values() []string {
	ret := make([]string, 0, len(r)+1)
	for _, x := range r {
		ret = append(ret, x.Value)
	}
	return ret
}

func (r RetroPermissions) ValueStrings(includeNil bool) []string {
	ret := make([]string, 0, len(r)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range r {
		ret = append(ret, x.Value)
	}
	return ret
}

func (r RetroPermissions) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(r)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range r {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (r RetroPermissions) Clone() RetroPermissions {
	return slices.Clone(r)
}
