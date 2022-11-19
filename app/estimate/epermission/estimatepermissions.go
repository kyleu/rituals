// Content managed by Project Forge, see [projectforge.md] for details.
package epermission

import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type EstimatePermissions []*EstimatePermission

func (e EstimatePermissions) Get(estimateID uuid.UUID, key string, value string) *EstimatePermission {
	for _, x := range e {
		if x.EstimateID == estimateID && x.Key == key && x.Value == value {
			return x
		}
	}
	return nil
}

func (e EstimatePermissions) EstimateIDs() []uuid.UUID {
	ret := make([]uuid.UUID, 0, len(e)+1)
	for _, x := range e {
		ret = append(ret, x.EstimateID)
	}
	return ret
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

func (e EstimatePermissions) Keys() []string {
	ret := make([]string, 0, len(e)+1)
	for _, x := range e {
		ret = append(ret, x.Key)
	}
	return ret
}

func (e EstimatePermissions) KeyStrings(includeNil bool) []string {
	ret := make([]string, 0, len(e)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range e {
		ret = append(ret, x.Key)
	}
	return ret
}

func (e EstimatePermissions) Values() []string {
	ret := make([]string, 0, len(e)+1)
	for _, x := range e {
		ret = append(ret, x.Value)
	}
	return ret
}

func (e EstimatePermissions) ValueStrings(includeNil bool) []string {
	ret := make([]string, 0, len(e)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range e {
		ret = append(ret, x.Value)
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
