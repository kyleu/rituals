// Content managed by Project Forge, see [projectforge.md] for details.
package rpermission

import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type RetroPermissions []*RetroPermission

func (r RetroPermissions) Get(retroID uuid.UUID, k string, v string) *RetroPermission {
	for _, x := range r {
		if x.RetroID == retroID && x.K == k && x.V == v {
			return x
		}
	}
	return nil
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

func (r RetroPermissions) KStrings(includeNil bool) []string {
	ret := make([]string, 0, len(r)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range r {
		ret = append(ret, x.K)
	}
	return ret
}

func (r RetroPermissions) VStrings(includeNil bool) []string {
	ret := make([]string, 0, len(r)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range r {
		ret = append(ret, x.V)
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
