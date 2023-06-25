// Content managed by Project Forge, see [projectforge.md] for details.
package rmember

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type RetroMembers []*RetroMember

func (r RetroMembers) Get(retroID uuid.UUID, userID uuid.UUID) *RetroMember {
	return lo.FindOrElse(r, nil, func(x *RetroMember) bool {
		return x.RetroID == retroID && x.UserID == userID
	})
}

func (r RetroMembers) GetByRetroIDs(retroIDs ...uuid.UUID) RetroMembers {
	return lo.Filter(r, func(x *RetroMember, _ int) bool {
		return lo.Contains(retroIDs, x.RetroID)
	})
}

func (r RetroMembers) RetroIDs() []uuid.UUID {
	return lo.Map(r, func(x *RetroMember, _ int) uuid.UUID {
		return x.RetroID
	})
}

func (r RetroMembers) RetroIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(r)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(r, func(x *RetroMember, _ int) {
		ret = append(ret, x.RetroID.String())
	})
	return ret
}

func (r RetroMembers) GetByUserIDs(userIDs ...uuid.UUID) RetroMembers {
	return lo.Filter(r, func(x *RetroMember, _ int) bool {
		return lo.Contains(userIDs, x.UserID)
	})
}

func (r RetroMembers) UserIDs() []uuid.UUID {
	return lo.Map(r, func(x *RetroMember, _ int) uuid.UUID {
		return x.UserID
	})
}

func (r RetroMembers) UserIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(r)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(r, func(x *RetroMember, _ int) {
		ret = append(ret, x.UserID.String())
	})
	return ret
}

func (r RetroMembers) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(r)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(r, func(x *RetroMember, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (r RetroMembers) Clone() RetroMembers {
	return slices.Clone(r)
}
