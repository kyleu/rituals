// Content managed by Project Forge, see [projectforge.md] for details.
package team

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type Teams []*Team

func (t Teams) Get(id uuid.UUID) *Team {
	return lo.FindOrElse(t, nil, func(x *Team) bool {
		return x.ID == id
	})
}

func (t Teams) GetByIDs(ids ...uuid.UUID) Teams {
	return lo.Filter(t, func(x *Team, _ int) bool {
		return lo.Contains(ids, x.ID)
	})
}

func (t Teams) IDs() []uuid.UUID {
	return lo.Map(t, func(x *Team, _ int) uuid.UUID {
		return x.ID
	})
}

func (t Teams) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(t)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(t, func(x *Team, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (t Teams) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(t)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(t, func(x *Team, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (t Teams) Clone() Teams {
	return slices.Clone(t)
}
