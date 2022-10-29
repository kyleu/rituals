// Content managed by Project Forge, see [projectforge.md] for details.
package tmember

import (
	"github.com/google/uuid"
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
