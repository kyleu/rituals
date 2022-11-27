// Content managed by Project Forge, see [projectforge.md] for details.
package umember

import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type StandupMembers []*StandupMember

func (s StandupMembers) Get(standupID uuid.UUID, userID uuid.UUID) *StandupMember {
	for _, x := range s {
		if x.StandupID == standupID && x.UserID == userID {
			return x
		}
	}
	return nil
}

func (s StandupMembers) GetByStandupIDs(standupIDs ...uuid.UUID) StandupMembers {
	var ret StandupMembers
	for _, x := range s {
		if slices.Contains(standupIDs, x.StandupID) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (s StandupMembers) StandupIDs() []uuid.UUID {
	ret := make([]uuid.UUID, 0, len(s)+1)
	for _, x := range s {
		ret = append(ret, x.StandupID)
	}
	return ret
}

func (s StandupMembers) StandupIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range s {
		ret = append(ret, x.StandupID.String())
	}
	return ret
}

func (s StandupMembers) GetByUserIDs(userIDs ...uuid.UUID) StandupMembers {
	var ret StandupMembers
	for _, x := range s {
		if slices.Contains(userIDs, x.UserID) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (s StandupMembers) UserIDs() []uuid.UUID {
	ret := make([]uuid.UUID, 0, len(s)+1)
	for _, x := range s {
		ret = append(ret, x.UserID)
	}
	return ret
}

func (s StandupMembers) UserIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range s {
		ret = append(ret, x.UserID.String())
	}
	return ret
}

func (s StandupMembers) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(s)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range s {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (s StandupMembers) Clone() StandupMembers {
	return slices.Clone(s)
}
