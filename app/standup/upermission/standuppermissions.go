// Content managed by Project Forge, see [projectforge.md] for details.
package upermission

import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type StandupPermissions []*StandupPermission

func (s StandupPermissions) Get(standupID uuid.UUID, k string, v string) *StandupPermission {
	for _, x := range s {
		if x.StandupID == standupID && x.K == k && x.V == v {
			return x
		}
	}
	return nil
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

func (s StandupPermissions) KStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range s {
		ret = append(ret, x.K)
	}
	return ret
}

func (s StandupPermissions) VStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range s {
		ret = append(ret, x.V)
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
