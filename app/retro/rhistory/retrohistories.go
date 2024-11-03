package rhistory

import (
	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

type RetroHistories []*RetroHistory

func (r RetroHistories) Get(slug string) *RetroHistory {
	return lo.FindOrElse(r, nil, func(x *RetroHistory) bool {
		return x.Slug == slug
	})
}

func (r RetroHistories) Slugs() []string {
	return lo.Map(r, func(xx *RetroHistory, _ int) string {
		return xx.Slug
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

func (r RetroHistories) GetBySlug(slug string) RetroHistories {
	return lo.Filter(r, func(xx *RetroHistory, _ int) bool {
		return xx.Slug == slug
	})
}

func (r RetroHistories) GetBySlugs(slugs ...string) RetroHistories {
	return lo.Filter(r, func(xx *RetroHistory, _ int) bool {
		return lo.Contains(slugs, xx.Slug)
	})
}

func (r RetroHistories) RetroIDs() []uuid.UUID {
	return lo.Map(r, func(xx *RetroHistory, _ int) uuid.UUID {
		return xx.RetroID
	})
}

func (r RetroHistories) GetByRetroID(retroID uuid.UUID) RetroHistories {
	return lo.Filter(r, func(xx *RetroHistory, _ int) bool {
		return xx.RetroID == retroID
	})
}

func (r RetroHistories) GetByRetroIDs(retroIDs ...uuid.UUID) RetroHistories {
	return lo.Filter(r, func(xx *RetroHistory, _ int) bool {
		return lo.Contains(retroIDs, xx.RetroID)
	})
}

func (r RetroHistories) ToCSV() ([]string, [][]string) {
	return RetroHistoryFieldDescs.Keys(), lo.Map(r, func(x *RetroHistory, _ int) []string {
		return x.Strings()
	})
}

func (r RetroHistories) Random() *RetroHistory {
	return util.RandomElement(r)
}

func (r RetroHistories) Clone() RetroHistories {
	return lo.Map(r, func(xx *RetroHistory, _ int) *RetroHistory {
		return xx.Clone()
	})
}
