// Content managed by Project Forge, see [projectforge.md] for details.
package tpermission

import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type TeamPermissions []*TeamPermission

func (t TeamPermissions) Get(teamID uuid.UUID, key string, value string) *TeamPermission {
	for _, x := range t {
		if x.TeamID == teamID && x.Key == key && x.Value == value {
			return x
		}
	}
	return nil
}

func (t TeamPermissions) TeamIDs() []uuid.UUID {
	ret := make([]uuid.UUID, 0, len(t)+1)
	for _, x := range t {
		ret = append(ret, x.TeamID)
	}
	return ret
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

func (t TeamPermissions) Keys() []string {
	ret := make([]string, 0, len(t)+1)
	for _, x := range t {
		ret = append(ret, x.Key)
	}
	return ret
}

func (t TeamPermissions) KeyStrings(includeNil bool) []string {
	ret := make([]string, 0, len(t)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range t {
		ret = append(ret, x.Key)
	}
	return ret
}

func (t TeamPermissions) Values() []string {
	ret := make([]string, 0, len(t)+1)
	for _, x := range t {
		ret = append(ret, x.Value)
	}
	return ret
}

func (t TeamPermissions) ValueStrings(includeNil bool) []string {
	ret := make([]string, 0, len(t)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range t {
		ret = append(ret, x.Value)
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
