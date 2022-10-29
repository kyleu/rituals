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
