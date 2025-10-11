package umember

import (
	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

type StandupMembers []*StandupMember

func (s StandupMembers) Get(standupID uuid.UUID, userID uuid.UUID) *StandupMember {
	return lo.FindOrElse(s, nil, func(x *StandupMember) bool {
		return x.StandupID == standupID && x.UserID == userID
	})
}

func (s StandupMembers) StandupIDs() []uuid.UUID {
	return lo.Map(s, func(xx *StandupMember, _ int) uuid.UUID {
		return xx.StandupID
	})
}

func (s StandupMembers) StandupIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(s, func(x *StandupMember, _ int) {
		ret = append(ret, x.StandupID.String())
	})
	return ret
}

func (s StandupMembers) UserIDs() []uuid.UUID {
	return lo.Map(s, func(xx *StandupMember, _ int) uuid.UUID {
		return xx.UserID
	})
}

func (s StandupMembers) UserIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(s, func(x *StandupMember, _ int) {
		ret = append(ret, x.UserID.String())
	})
	return ret
}

func (s StandupMembers) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(s)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(s, func(x *StandupMember, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (s StandupMembers) ToPKs() []*PK {
	return lo.Map(s, func(x *StandupMember, _ int) *PK {
		return x.ToPK()
	})
}

func (s StandupMembers) GetByStandupID(standupID uuid.UUID) StandupMembers {
	return lo.Filter(s, func(xx *StandupMember, _ int) bool {
		return xx.StandupID == standupID
	})
}

func (s StandupMembers) GetByStandupIDs(standupIDs ...uuid.UUID) StandupMembers {
	return lo.Filter(s, func(xx *StandupMember, _ int) bool {
		return lo.Contains(standupIDs, xx.StandupID)
	})
}

func (s StandupMembers) GetByUserID(userID uuid.UUID) StandupMembers {
	return lo.Filter(s, func(xx *StandupMember, _ int) bool {
		return xx.UserID == userID
	})
}

func (s StandupMembers) GetByUserIDs(userIDs ...uuid.UUID) StandupMembers {
	return lo.Filter(s, func(xx *StandupMember, _ int) bool {
		return lo.Contains(userIDs, xx.UserID)
	})
}

func (s StandupMembers) ToMaps() []util.ValueMap {
	return lo.Map(s, func(xx *StandupMember, _ int) util.ValueMap {
		return xx.ToMap()
	})
}

func (s StandupMembers) ToOrderedMaps() util.OrderedMaps[any] {
	return lo.Map(s, func(x *StandupMember, _ int) *util.OrderedMap[any] {
		return x.ToOrderedMap()
	})
}

func (s StandupMembers) ToCSV() ([]string, [][]string) {
	return StandupMemberFieldDescs.Keys(), lo.Map(s, func(x *StandupMember, _ int) []string {
		return x.Strings()
	})
}

func (s StandupMembers) Random() *StandupMember {
	return util.RandomElement(s)
}

func (s StandupMembers) Clone() StandupMembers {
	return lo.Map(s, func(xx *StandupMember, _ int) *StandupMember {
		return xx.Clone()
	})
}
