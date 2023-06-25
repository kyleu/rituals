// Content managed by Project Forge, see [projectforge.md] for details.
package smember

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type SprintMembers []*SprintMember

func (s SprintMembers) Get(sprintID uuid.UUID, userID uuid.UUID) *SprintMember {
	return lo.FindOrElse(s, nil, func(x *SprintMember) bool {
		return x.SprintID == sprintID && x.UserID == userID
	})
}

func (s SprintMembers) GetBySprintIDs(sprintIDs ...uuid.UUID) SprintMembers {
	return lo.Filter(s, func(x *SprintMember, _ int) bool {
		return lo.Contains(sprintIDs, x.SprintID)
	})
}

func (s SprintMembers) SprintIDs() []uuid.UUID {
	return lo.Map(s, func(x *SprintMember, _ int) uuid.UUID {
		return x.SprintID
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

func (s SprintMembers) GetByUserIDs(userIDs ...uuid.UUID) SprintMembers {
	return lo.Filter(s, func(x *SprintMember, _ int) bool {
		return lo.Contains(userIDs, x.UserID)
	})
}

func (s SprintMembers) UserIDs() []uuid.UUID {
	return lo.Map(s, func(x *SprintMember, _ int) uuid.UUID {
		return x.UserID
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

func (s SprintMembers) Clone() SprintMembers {
	return slices.Clone(s)
}
