// Content managed by Project Forge, see [projectforge.md] for details.
package rmember

import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type RetroMembers []*RetroMember

func (r RetroMembers) Get(retroID uuid.UUID, userID uuid.UUID) *RetroMember {
	for _, x := range r {
		if x.RetroID == retroID && x.UserID == userID {
			return x
		}
	}
	return nil
}

func (r RetroMembers) RetroIDs() []uuid.UUID {
	ret := make([]uuid.UUID, 0, len(r)+1)
	for _, x := range r {
		ret = append(ret, x.RetroID)
	}
	return ret
}

func (r RetroMembers) RetroIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(r)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range r {
		ret = append(ret, x.RetroID.String())
	}
	return ret
}

func (r RetroMembers) UserIDs() []uuid.UUID {
	ret := make([]uuid.UUID, 0, len(r)+1)
	for _, x := range r {
		ret = append(ret, x.UserID)
	}
	return ret
}

func (r RetroMembers) UserIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(r)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range r {
		ret = append(ret, x.UserID.String())
	}
	return ret
}

func (r RetroMembers) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(r)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range r {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (r RetroMembers) Clone() RetroMembers {
	return slices.Clone(r)
}
