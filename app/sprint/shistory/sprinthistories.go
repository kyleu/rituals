// Content managed by Project Forge, see [projectforge.md] for details.
package shistory

import (
	"slices"

	"github.com/samber/lo"
)

type SprintHistories []*SprintHistory

func (s SprintHistories) Get(slug string) *SprintHistory {
	return lo.FindOrElse(s, nil, func(x *SprintHistory) bool {
		return x.Slug == slug
	})
}

func (s SprintHistories) GetBySlugs(slugs ...string) SprintHistories {
	return lo.Filter(s, func(x *SprintHistory, _ int) bool {
		return lo.Contains(slugs, x.Slug)
	})
}

func (s SprintHistories) Slugs() []string {
	return lo.Map(s, func(x *SprintHistory, _ int) string {
		return x.Slug
	})
}

func (s SprintHistories) SlugStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(s, func(x *SprintHistory, _ int) {
		ret = append(ret, x.Slug)
	})
	return ret
}

func (s SprintHistories) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(s)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(s, func(x *SprintHistory, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (s SprintHistories) Clone() SprintHistories {
	return slices.Clone(s)
}
