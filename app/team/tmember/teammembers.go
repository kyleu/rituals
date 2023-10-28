// Package tmember - Content managed by Project Forge, see [projectforge.md] for details.
package tmember

import (
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

type TeamMembers []*TeamMember

func (t TeamMembers) Get(teamID uuid.UUID, userID uuid.UUID) *TeamMember {
	return lo.FindOrElse(t, nil, func(x *TeamMember) bool {
		return x.TeamID == teamID && x.UserID == userID
	})
}

func (t TeamMembers) TeamIDs() []uuid.UUID {
	return lo.Map(t, func(xx *TeamMember, _ int) uuid.UUID {
		return xx.TeamID
	})
}

func (t TeamMembers) TeamIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(t)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(t, func(x *TeamMember, _ int) {
		ret = append(ret, x.TeamID.String())
	})
	return ret
}

func (t TeamMembers) UserIDs() []uuid.UUID {
	return lo.Map(t, func(xx *TeamMember, _ int) uuid.UUID {
		return xx.UserID
	})
}

func (t TeamMembers) UserIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(t)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(t, func(x *TeamMember, _ int) {
		ret = append(ret, x.UserID.String())
	})
	return ret
}

func (t TeamMembers) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(t)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(t, func(x *TeamMember, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (t TeamMembers) ToPKs() []*PK {
	return lo.Map(t, func(x *TeamMember, _ int) *PK {
		return x.ToPK()
	})
}

func (t TeamMembers) GetByTeamID(teamID uuid.UUID) TeamMembers {
	return lo.Filter(t, func(xx *TeamMember, _ int) bool {
		return xx.TeamID == teamID
	})
}

func (t TeamMembers) GetByTeamIDs(teamIDs ...uuid.UUID) TeamMembers {
	return lo.Filter(t, func(xx *TeamMember, _ int) bool {
		return lo.Contains(teamIDs, xx.TeamID)
	})
}

func (t TeamMembers) GetByUserID(userID uuid.UUID) TeamMembers {
	return lo.Filter(t, func(xx *TeamMember, _ int) bool {
		return xx.UserID == userID
	})
}

func (t TeamMembers) GetByUserIDs(userIDs ...uuid.UUID) TeamMembers {
	return lo.Filter(t, func(xx *TeamMember, _ int) bool {
		return lo.Contains(userIDs, xx.UserID)
	})
}

func (t TeamMembers) Random() *TeamMember {
	if len(t) == 0 {
		return nil
	}
	return t[util.RandomInt(len(t))]
}

func (t TeamMembers) Clone() TeamMembers {
	return slices.Clone(t)
}
