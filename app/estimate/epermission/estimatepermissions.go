// Content managed by Project Forge, see [projectforge.md] for details.
package epermission

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
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

func (e EstimatePermissions) GetByEstimateIDs(estimateIDs ...uuid.UUID) EstimatePermissions {
	var ret EstimatePermissions
	for _, x := range e {
		if lo.Contains(estimateIDs, x.EstimateID) {
			ret = append(ret, x)
		}
	}
	return ret
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

func (e EstimatePermissions) GetByKeys(keys ...string) EstimatePermissions {
	var ret EstimatePermissions
	for _, x := range e {
		if lo.Contains(keys, x.Key) {
			ret = append(ret, x)
		}
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

func (e EstimatePermissions) GetByValues(values ...string) EstimatePermissions {
	var ret EstimatePermissions
	for _, x := range e {
		if lo.Contains(values, x.Value) {
			ret = append(ret, x)
		}
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
