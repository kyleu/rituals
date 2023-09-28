// Package retro - Content managed by Project Forge, see [projectforge.md] for details.
package retro

import (
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Retros []*Retro

func (r Retros) Get(id uuid.UUID) *Retro {
	return lo.FindOrElse(r, nil, func(x *Retro) bool {
		return x.ID == id
	})
}

func (r Retros) GetByIDs(ids ...uuid.UUID) Retros {
	return lo.Filter(r, func(x *Retro, _ int) bool {
		return lo.Contains(ids, x.ID)
	})
}

func (r Retros) IDs() []uuid.UUID {
	return lo.Map(r, func(x *Retro, _ int) uuid.UUID {
		return x.ID
	})
}

func (r Retros) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(r)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(r, func(x *Retro, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (r Retros) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(r)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(r, func(x *Retro, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (r Retros) Clone() Retros {
	return slices.Clone(r)
}
