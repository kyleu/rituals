// Content managed by Project Forge, see [projectforge.md] for details.
package uhistory

import (
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type StandupHistories []*StandupHistory

func (s StandupHistories) Get(slug string) *StandupHistory {
	return lo.FindOrElse(s, nil, func(x *StandupHistory) bool {
		return x.Slug == slug
	})
}

func (s StandupHistories) GetBySlugs(slugs ...string) StandupHistories {
	return lo.Filter(s, func(x *StandupHistory, _ int) bool {
		return lo.Contains(slugs, x.Slug)
	})
}

func (s StandupHistories) Slugs() []string {
	return lo.Map(s, func(x *StandupHistory, _ int) string {
		return x.Slug
	})
}

func (s StandupHistories) SlugStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(s, func(x *StandupHistory, _ int) {
		ret = append(ret, x.Slug)
	})
	return ret
}

func (s StandupHistories) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(s)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(s, func(x *StandupHistory, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (s StandupHistories) Clone() StandupHistories {
	return slices.Clone(s)
}
