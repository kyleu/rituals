package ehistory

import (
	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

type EstimateHistories []*EstimateHistory

func (e EstimateHistories) Get(slug string) *EstimateHistory {
	return lo.FindOrElse(e, nil, func(x *EstimateHistory) bool {
		return x.Slug == slug
	})
}

func (e EstimateHistories) Slugs() []string {
	return lo.Map(e, func(xx *EstimateHistory, _ int) string {
		return xx.Slug
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

func (e EstimateHistories) GetBySlug(slug string) EstimateHistories {
	return lo.Filter(e, func(xx *EstimateHistory, _ int) bool {
		return xx.Slug == slug
	})
}

func (e EstimateHistories) GetBySlugs(slugs ...string) EstimateHistories {
	return lo.Filter(e, func(xx *EstimateHistory, _ int) bool {
		return lo.Contains(slugs, xx.Slug)
	})
}

func (e EstimateHistories) EstimateIDs() []uuid.UUID {
	return lo.Map(e, func(xx *EstimateHistory, _ int) uuid.UUID {
		return xx.EstimateID
	})
}

func (e EstimateHistories) GetByEstimateID(estimateID uuid.UUID) EstimateHistories {
	return lo.Filter(e, func(xx *EstimateHistory, _ int) bool {
		return xx.EstimateID == estimateID
	})
}

func (e EstimateHistories) GetByEstimateIDs(estimateIDs ...uuid.UUID) EstimateHistories {
	return lo.Filter(e, func(xx *EstimateHistory, _ int) bool {
		return lo.Contains(estimateIDs, xx.EstimateID)
	})
}

func (e EstimateHistories) ToCSV() ([]string, [][]string) {
	return EstimateHistoryFieldDescs.Keys(), lo.Map(e, func(x *EstimateHistory, _ int) []string {
		return x.Strings()
	})
}

func (e EstimateHistories) Random() *EstimateHistory {
	return util.RandomElement(e)
}

func (e EstimateHistories) Clone() EstimateHistories {
	return lo.Map(e, func(xx *EstimateHistory, _ int) *EstimateHistory {
		return xx.Clone()
	})
}
