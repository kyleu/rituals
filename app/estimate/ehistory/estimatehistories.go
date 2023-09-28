// Package ehistory - Content managed by Project Forge, see [projectforge.md] for details.
package ehistory

import (
	"slices"

	"github.com/samber/lo"
)

type EstimateHistories []*EstimateHistory

func (e EstimateHistories) Get(slug string) *EstimateHistory {
	return lo.FindOrElse(e, nil, func(x *EstimateHistory) bool {
		return x.Slug == slug
	})
}

func (e EstimateHistories) GetBySlugs(slugs ...string) EstimateHistories {
	return lo.Filter(e, func(x *EstimateHistory, _ int) bool {
		return lo.Contains(slugs, x.Slug)
	})
}

func (e EstimateHistories) Slugs() []string {
	return lo.Map(e, func(x *EstimateHistory, _ int) string {
		return x.Slug
	})
}

func (e EstimateHistories) SlugStrings(includeNil bool) []string {
	ret := make([]string, 0, len(e)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(e, func(x *EstimateHistory, _ int) {
		ret = append(ret, x.Slug)
	})
	return ret
}

func (e EstimateHistories) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(e)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(e, func(x *EstimateHistory, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (e EstimateHistories) Clone() EstimateHistories {
	return slices.Clone(e)
}
