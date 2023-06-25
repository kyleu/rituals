// Content managed by Project Forge, see [projectforge.md] for details.
package rhistory

import (
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type RetroHistories []*RetroHistory

func (r RetroHistories) Get(slug string) *RetroHistory {
	return lo.FindOrElse(r, nil, func(x *RetroHistory) bool {
		return x.Slug == slug
	})
}

func (r RetroHistories) GetBySlugs(slugs ...string) RetroHistories {
	return lo.Filter(r, func(x *RetroHistory, _ int) bool {
		return lo.Contains(slugs, x.Slug)
	})
}

func (r RetroHistories) Slugs() []string {
	return lo.Map(r, func(x *RetroHistory, _ int) string {
		return x.Slug
	})
}

func (r RetroHistories) SlugStrings(includeNil bool) []string {
	ret := make([]string, 0, len(r)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(r, func(x *RetroHistory, _ int) {
		ret = append(ret, x.Slug)
	})
	return ret
}

func (r RetroHistories) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(r)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(r, func(x *RetroHistory, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (r RetroHistories) Clone() RetroHistories {
	return slices.Clone(r)
}
