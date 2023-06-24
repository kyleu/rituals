// Content managed by Project Forge, see [projectforge.md] for details.
package tmember

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type TeamMembers []*TeamMember

func (t TeamMembers) Get(teamID uuid.UUID, userID uuid.UUID) *TeamMember {
	for _, x := range t {
		if x.TeamID == teamID && x.UserID == userID {
			return x
		}
	}
	return nil
}

func (t TeamMembers) GetByTeamIDs(teamIDs ...uuid.UUID) TeamMembers {
	var ret TeamMembers
	for _, x := range t {
		if lo.Contains(teamIDs, x.TeamID) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (t TeamMembers) TeamIDs() []uuid.UUID {
	ret := make([]uuid.UUID, 0, len(t)+1)
	for _, x := range t {
		ret = append(ret, x.TeamID)
	}
	return ret
}

func (t TeamMembers) TeamIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(t)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range t {
		ret = append(ret, x.TeamID.String())
	}
	return ret
}

func (t TeamMembers) GetByUserIDs(userIDs ...uuid.UUID) TeamMembers {
	var ret TeamMembers
	for _, x := range t {
		if lo.Contains(userIDs, x.UserID) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (t TeamMembers) UserIDs() []uuid.UUID {
	ret := make([]uuid.UUID, 0, len(t)+1)
	for _, x := range t {
		ret = append(ret, x.UserID)
	}
	return ret
}

func (t TeamMembers) UserIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(t)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range t {
		ret = append(ret, x.UserID.String())
	}
	return ret
}

func (t TeamMembers) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(t)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range t {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (t TeamMembers) Clone() TeamMembers {
	return slices.Clone(t)
}
