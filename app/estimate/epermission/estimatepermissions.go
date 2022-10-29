// Content managed by Project Forge, see [projectforge.md] for details.
package epermission

import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type EstimatePermissions []*EstimatePermission

func (e EstimatePermissions) Get(estimateID uuid.UUID, k string, v string) *EstimatePermission {
	for _, x := range e {
		if x.EstimateID == estimateID && x.K == k && x.V == v {
			return x
		}
	}
	return nil
}

func (e EstimatePermissions) EstimateIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(e)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range e {
		ret = append(ret, x.EstimateID.String())
	}
	return ret
}

func (e EstimatePermissions) KStrings(includeNil bool) []string {
	ret := make([]string, 0, len(e)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range e {
		ret = append(ret, x.K)
	}
	return ret
}

func (e EstimatePermissions) VStrings(includeNil bool) []string {
	ret := make([]string, 0, len(e)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range e {
		ret = append(ret, x.V)
	}
	return ret
}

func (e EstimatePermissions) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(e)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range e {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (e EstimatePermissions) Clone() EstimatePermissions {
	return slices.Clone(e)
}
