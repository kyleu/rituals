// Content managed by Project Forge, see [projectforge.md] for details.
package story

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type Stories []*Story

func (s Stories) Get(id uuid.UUID) *Story {
	return lo.FindOrElse(s, nil, func(x *Story) bool {
		return x.ID == id
	})
}

func (s Stories) GetByIDs(ids ...uuid.UUID) Stories {
	return lo.Filter(s, func(x *Story, _ int) bool {
		return lo.Contains(ids, x.ID)
	})
}

func (s Stories) IDs() []uuid.UUID {
	return lo.Map(s, func(x *Story, _ int) uuid.UUID {
		return x.ID
	})
}

func (s Stories) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(s, func(x *Story, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (s Stories) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(s)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(s, func(x *Story, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (s Stories) Clone() Stories {
	return slices.Clone(s)
}
