// Content managed by Project Forge, see [projectforge.md] for details.
package tpermission

import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type TeamPermissions []*TeamPermission

func (t TeamPermissions) Get(teamID uuid.UUID, k string, v string) *TeamPermission {
	for _, x := range t {
		if x.TeamID == teamID && x.K == k && x.V == v {
			return x
		}
	}
	return nil
}

func (t TeamPermissions) TeamIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(t)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range t {
		ret = append(ret, x.TeamID.String())
	}
	return ret
}

func (t TeamPermissions) KStrings(includeNil bool) []string {
	ret := make([]string, 0, len(t)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range t {
		ret = append(ret, x.K)
	}
	return ret
}

func (t TeamPermissions) VStrings(includeNil bool) []string {
	ret := make([]string, 0, len(t)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range t {
		ret = append(ret, x.V)
	}
	return ret
}

func (t TeamPermissions) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(t)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range t {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (t TeamPermissions) Clone() TeamPermissions {
	return slices.Clone(t)
}
