// Content managed by Project Forge, see [projectforge.md] for details.
package report

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type Reports []*Report

func (r Reports) Get(id uuid.UUID) *Report {
	return lo.FindOrElse(r, nil, func(x *Report) bool {
		return x.ID == id
	})
}

func (r Reports) GetByIDs(ids ...uuid.UUID) Reports {
	return lo.Filter(r, func(x *Report, _ int) bool {
		return lo.Contains(ids, x.ID)
	})
}

func (r Reports) IDs() []uuid.UUID {
	return lo.Map(r, func(x *Report, _ int) uuid.UUID {
		return x.ID
	})
}

func (r Reports) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(r)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(r, func(x *Report, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (r Reports) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(r)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(r, func(x *Report, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (r Reports) Clone() Reports {
	return slices.Clone(r)
}
