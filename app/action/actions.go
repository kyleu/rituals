// Package action - Content managed by Project Forge, see [projectforge.md] for details.
package action

import (
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Actions []*Action

func (a Actions) Get(id uuid.UUID) *Action {
	return lo.FindOrElse(a, nil, func(x *Action) bool {
		return x.ID == id
	})
}

func (a Actions) IDs() []uuid.UUID {
	return lo.Map(a, func(xx *Action, _ int) uuid.UUID {
		return xx.ID
	})
}

func (a Actions) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(a)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(a, func(x *Action, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (a Actions) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(a)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(a, func(x *Action, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (a Actions) GetByID(id uuid.UUID) Actions {
	return lo.Filter(a, func(xx *Action, _ int) bool {
		return xx.ID == id
	})
}

func (a Actions) GetByIDs(ids ...uuid.UUID) Actions {
	return lo.Filter(a, func(xx *Action, _ int) bool {
		return lo.Contains(ids, xx.ID)
	})
}

func (a Actions) UserIDs() []uuid.UUID {
	return lo.Map(a, func(xx *Action, _ int) uuid.UUID {
		return xx.UserID
	})
}

func (a Actions) GetByUserID(userID uuid.UUID) Actions {
	return lo.Filter(a, func(xx *Action, _ int) bool {
		return xx.UserID == userID
	})
}

func (a Actions) GetByUserIDs(userIDs ...uuid.UUID) Actions {
	return lo.Filter(a, func(xx *Action, _ int) bool {
		return lo.Contains(userIDs, xx.UserID)
	})
}

func (a Actions) Clone() Actions {
	return slices.Clone(a)
}
