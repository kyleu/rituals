// Content managed by Project Forge, see [projectforge.md] for details.
package smember

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type SprintMembers []*SprintMember

func (s SprintMembers) Get(sprintID uuid.UUID, userID uuid.UUID) *SprintMember {
	for _, x := range s {
		if x.SprintID == sprintID && x.UserID == userID {
			return x
		}
	}
	return nil
}

func (s SprintMembers) GetBySprintIDs(sprintIDs ...uuid.UUID) SprintMembers {
	var ret SprintMembers
	for _, x := range s {
		if lo.Contains(sprintIDs, x.SprintID) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (s SprintMembers) SprintIDs() []uuid.UUID {
	ret := make([]uuid.UUID, 0, len(s)+1)
	for _, x := range s {
		ret = append(ret, x.SprintID)
	}
	return ret
}

func (s SprintMembers) SprintIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range s {
		ret = append(ret, x.SprintID.String())
	}
	return ret
}

func (s SprintMembers) GetByUserIDs(userIDs ...uuid.UUID) SprintMembers {
	var ret SprintMembers
	for _, x := range s {
		if lo.Contains(userIDs, x.UserID) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (s SprintMembers) UserIDs() []uuid.UUID {
	ret := make([]uuid.UUID, 0, len(s)+1)
	for _, x := range s {
		ret = append(ret, x.UserID)
	}
	return ret
}

func (s SprintMembers) UserIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range s {
		ret = append(ret, x.UserID.String())
	}
	return ret
}

func (s SprintMembers) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(s)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range s {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (s SprintMembers) Clone() SprintMembers {
	return slices.Clone(s)
}
