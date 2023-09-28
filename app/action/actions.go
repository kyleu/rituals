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

func (a Actions) GetByIDs(ids ...uuid.UUID) Actions {
	return lo.Filter(a, func(x *Action, _ int) bool {
		return lo.Contains(ids, x.ID)
	})
}

func (a Actions) IDs() []uuid.UUID {
	return lo.Map(a, func(x *Action, _ int) uuid.UUID {
		return x.ID
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

func (a Actions) Clone() Actions {
	return slices.Clone(a)
}
