// Content managed by Project Forge, see [projectforge.md] for details.
package emember

import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type EstimateMembers []*EstimateMember

func (e EstimateMembers) Get(estimateID uuid.UUID, userID uuid.UUID) *EstimateMember {
	for _, x := range e {
		if x.EstimateID == estimateID && x.UserID == userID {
			return x
		}
	}
	return nil
}

func (e EstimateMembers) GetByEstimateIDs(estimateIDs ...uuid.UUID) EstimateMembers {
	var ret EstimateMembers
	for _, x := range e {
		if slices.Contains(estimateIDs, x.EstimateID) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (e EstimateMembers) EstimateIDs() []uuid.UUID {
	ret := make([]uuid.UUID, 0, len(e)+1)
	for _, x := range e {
		ret = append(ret, x.EstimateID)
	}
	return ret
}

func (e EstimateMembers) EstimateIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(e)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range e {
		ret = append(ret, x.EstimateID.String())
	}
	return ret
}

func (e EstimateMembers) GetByUserIDs(userIDs ...uuid.UUID) EstimateMembers {
	var ret EstimateMembers
	for _, x := range e {
		if slices.Contains(userIDs, x.UserID) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (e EstimateMembers) UserIDs() []uuid.UUID {
	ret := make([]uuid.UUID, 0, len(e)+1)
	for _, x := range e {
		ret = append(ret, x.UserID)
	}
	return ret
}

func (e EstimateMembers) UserIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(e)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range e {
		ret = append(ret, x.UserID.String())
	}
	return ret
}

func (e EstimateMembers) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(e)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range e {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (e EstimateMembers) Clone() EstimateMembers {
	return slices.Clone(e)
}
