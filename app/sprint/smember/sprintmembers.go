package smember

import (
	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

type SprintMembers []*SprintMember

func (s SprintMembers) Get(sprintID uuid.UUID, userID uuid.UUID) *SprintMember {
	return lo.FindOrElse(s, nil, func(x *SprintMember) bool {
		return x.SprintID == sprintID && x.UserID == userID
	})
}

func (s SprintMembers) SprintIDs() []uuid.UUID {
	return lo.Map(s, func(xx *SprintMember, _ int) uuid.UUID {
		return xx.SprintID
	})
}

func (s SprintMembers) SprintIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(s, func(x *SprintMember, _ int) {
		ret = append(ret, x.SprintID.String())
	})
	return ret
}

func (s SprintMembers) UserIDs() []uuid.UUID {
	return lo.Map(s, func(xx *SprintMember, _ int) uuid.UUID {
		return xx.UserID
	})
}

func (s SprintMembers) UserIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(s, func(x *SprintMember, _ int) {
		ret = append(ret, x.UserID.String())
	})
	return ret
}

func (s SprintMembers) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(s)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(s, func(x *SprintMember, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (s SprintMembers) ToPKs() []*PK {
	return lo.Map(s, func(x *SprintMember, _ int) *PK {
		return x.ToPK()
	})
}

func (s SprintMembers) GetBySprintID(sprintID uuid.UUID) SprintMembers {
	return lo.Filter(s, func(xx *SprintMember, _ int) bool {
		return xx.SprintID == sprintID
	})
}

func (s SprintMembers) GetBySprintIDs(sprintIDs ...uuid.UUID) SprintMembers {
	return lo.Filter(s, func(xx *SprintMember, _ int) bool {
		return lo.Contains(sprintIDs, xx.SprintID)
	})
}

func (s SprintMembers) GetByUserID(userID uuid.UUID) SprintMembers {
	return lo.Filter(s, func(xx *SprintMember, _ int) bool {
		return xx.UserID == userID
	})
}

func (s SprintMembers) GetByUserIDs(userIDs ...uuid.UUID) SprintMembers {
	return lo.Filter(s, func(xx *SprintMember, _ int) bool {
		return lo.Contains(userIDs, xx.UserID)
	})
}

func (s SprintMembers) ToCSV() ([]string, [][]string) {
	return SprintMemberFieldDescs.Keys(), lo.Map(s, func(x *SprintMember, _ int) []string {
		return x.Strings()
	})
}

func (s SprintMembers) Random() *SprintMember {
	return util.RandomElement(s)
}

func (s SprintMembers) Clone() SprintMembers {
	return lo.Map(s, func(xx *SprintMember, _ int) *SprintMember {
		return xx.Clone()
	})
}
