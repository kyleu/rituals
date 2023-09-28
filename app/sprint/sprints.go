// Package sprint - Content managed by Project Forge, see [projectforge.md] for details.
package sprint

import (
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Sprints []*Sprint

func (s Sprints) Get(id uuid.UUID) *Sprint {
	return lo.FindOrElse(s, nil, func(x *Sprint) bool {
		return x.ID == id
	})
}

func (s Sprints) GetByIDs(ids ...uuid.UUID) Sprints {
	return lo.Filter(s, func(x *Sprint, _ int) bool {
		return lo.Contains(ids, x.ID)
	})
}

func (s Sprints) IDs() []uuid.UUID {
	return lo.Map(s, func(x *Sprint, _ int) uuid.UUID {
		return x.ID
	})
}

func (s Sprints) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(s, func(x *Sprint, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (s Sprints) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(s)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(s, func(x *Sprint, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (s Sprints) Clone() Sprints {
	return slices.Clone(s)
}
