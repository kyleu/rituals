package emember

import (
	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

type EstimateMembers []*EstimateMember

func (e EstimateMembers) Get(estimateID uuid.UUID, userID uuid.UUID) *EstimateMember {
	return lo.FindOrElse(e, nil, func(x *EstimateMember) bool {
		return x.EstimateID == estimateID && x.UserID == userID
	})
}

func (e EstimateMembers) EstimateIDs() []uuid.UUID {
	return lo.Map(e, func(xx *EstimateMember, _ int) uuid.UUID {
		return xx.EstimateID
	})
}

func (e EstimateMembers) EstimateIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(e)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(e, func(x *EstimateMember, _ int) {
		ret = append(ret, x.EstimateID.String())
	})
	return ret
}

func (e EstimateMembers) UserIDs() []uuid.UUID {
	return lo.Map(e, func(xx *EstimateMember, _ int) uuid.UUID {
		return xx.UserID
	})
}

func (e EstimateMembers) UserIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(e)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(e, func(x *EstimateMember, _ int) {
		ret = append(ret, x.UserID.String())
	})
	return ret
}

func (e EstimateMembers) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(e)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(e, func(x *EstimateMember, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (e EstimateMembers) ToPKs() []*PK {
	return lo.Map(e, func(x *EstimateMember, _ int) *PK {
		return x.ToPK()
	})
}

func (e EstimateMembers) GetByEstimateID(estimateID uuid.UUID) EstimateMembers {
	return lo.Filter(e, func(xx *EstimateMember, _ int) bool {
		return xx.EstimateID == estimateID
	})
}

func (e EstimateMembers) GetByEstimateIDs(estimateIDs ...uuid.UUID) EstimateMembers {
	return lo.Filter(e, func(xx *EstimateMember, _ int) bool {
		return lo.Contains(estimateIDs, xx.EstimateID)
	})
}

func (e EstimateMembers) GetByUserID(userID uuid.UUID) EstimateMembers {
	return lo.Filter(e, func(xx *EstimateMember, _ int) bool {
		return xx.UserID == userID
	})
}

func (e EstimateMembers) GetByUserIDs(userIDs ...uuid.UUID) EstimateMembers {
	return lo.Filter(e, func(xx *EstimateMember, _ int) bool {
		return lo.Contains(userIDs, xx.UserID)
	})
}

func (e EstimateMembers) ToMap() map[*PK]*EstimateMember {
	return lo.SliceToMap(e, func(xx *EstimateMember) (*PK, *EstimateMember) {
		return xx.ToPK(), xx
	})
}

func (e EstimateMembers) ToMaps() []util.ValueMap {
	return lo.Map(e, func(xx *EstimateMember, _ int) util.ValueMap {
		return xx.ToMap()
	})
}

func (e EstimateMembers) ToOrderedMaps() util.OrderedMaps[any] {
	return lo.Map(e, func(x *EstimateMember, _ int) *util.OrderedMap[any] {
		return x.ToOrderedMap()
	})
}

func (e EstimateMembers) ToCSV() ([]string, [][]string) {
	return EstimateMemberFieldDescs.Keys(), lo.Map(e, func(x *EstimateMember, _ int) []string {
		return x.Strings()
	})
}

func (e EstimateMembers) Random() *EstimateMember {
	return util.RandomElement(e)
}

func (e EstimateMembers) Clone() EstimateMembers {
	return lo.Map(e, func(xx *EstimateMember, _ int) *EstimateMember {
		return xx.Clone()
	})
}
