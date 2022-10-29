// Content managed by Project Forge, see [projectforge.md] for details.
package spermission

import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type SprintPermissions []*SprintPermission

func (s SprintPermissions) Get(sprintID uuid.UUID, k string, v string) *SprintPermission {
	for _, x := range s {
		if x.SprintID == sprintID && x.K == k && x.V == v {
			return x
		}
	}
	return nil
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

func (s SprintPermissions) KStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range s {
		ret = append(ret, x.K)
	}
	return ret
}

func (s SprintPermissions) VStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range s {
		ret = append(ret, x.V)
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
