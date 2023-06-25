// Content managed by Project Forge, see [projectforge.md] for details.
package standup

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type Standups []*Standup

func (s Standups) Get(id uuid.UUID) *Standup {
	return lo.FindOrElse(s, nil, func(x *Standup) bool {
		return x.ID == id
	})
}

func (s Standups) GetByIDs(ids ...uuid.UUID) Standups {
	return lo.Filter(s, func(x *Standup, _ int) bool {
		return lo.Contains(ids, x.ID)
	})
}

func (s Standups) IDs() []uuid.UUID {
	return lo.Map(s, func(x *Standup, _ int) uuid.UUID {
		return x.ID
	})
}

func (s Standups) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(s, func(x *Standup, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (s Standups) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(s)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(s, func(x *Standup, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (s Standups) Clone() Standups {
	return slices.Clone(s)
}
