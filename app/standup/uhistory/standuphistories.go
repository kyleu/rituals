package uhistory

import (
	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

type StandupHistories []*StandupHistory

func (s StandupHistories) Get(slug string) *StandupHistory {
	return lo.FindOrElse(s, nil, func(x *StandupHistory) bool {
		return x.Slug == slug
	})
}

func (s StandupHistories) Slugs() []string {
	return lo.Map(s, func(xx *StandupHistory, _ int) string {
		return xx.Slug
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

func (s StandupHistories) GetBySlug(slug string) StandupHistories {
	return lo.Filter(s, func(xx *StandupHistory, _ int) bool {
		return xx.Slug == slug
	})
}

func (s StandupHistories) GetBySlugs(slugs ...string) StandupHistories {
	return lo.Filter(s, func(xx *StandupHistory, _ int) bool {
		return lo.Contains(slugs, xx.Slug)
	})
}

func (s StandupHistories) StandupIDs() []uuid.UUID {
	return lo.Map(s, func(xx *StandupHistory, _ int) uuid.UUID {
		return xx.StandupID
	})
}

func (s StandupHistories) GetByStandupID(standupID uuid.UUID) StandupHistories {
	return lo.Filter(s, func(xx *StandupHistory, _ int) bool {
		return xx.StandupID == standupID
	})
}

func (s StandupHistories) GetByStandupIDs(standupIDs ...uuid.UUID) StandupHistories {
	return lo.Filter(s, func(xx *StandupHistory, _ int) bool {
		return lo.Contains(standupIDs, xx.StandupID)
	})
}

func (s StandupHistories) ToMaps() []util.ValueMap {
	return lo.Map(s, func(xx *StandupHistory, _ int) util.ValueMap {
		return xx.ToMap()
	})
}

func (s StandupHistories) ToOrderedMaps() util.OrderedMaps[any] {
	return lo.Map(s, func(x *StandupHistory, _ int) *util.OrderedMap[any] {
		return x.ToOrderedMap()
	})
}

func (s StandupHistories) ToCSV() ([]string, [][]string) {
	return StandupHistoryFieldDescs.Keys(), lo.Map(s, func(x *StandupHistory, _ int) []string {
		return x.Strings()
	})
}

func (s StandupHistories) Random() *StandupHistory {
	return util.RandomElement(s)
}

func (s StandupHistories) Clone() StandupHistories {
	return lo.Map(s, func(xx *StandupHistory, _ int) *StandupHistory {
		return xx.Clone()
	})
}
