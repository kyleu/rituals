package shistory

import (
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

type SprintHistories []*SprintHistory

func (s SprintHistories) Get(slug string) *SprintHistory {
	return lo.FindOrElse(s, nil, func(x *SprintHistory) bool {
		return x.Slug == slug
	})
}

func (s SprintHistories) Slugs() []string {
	return lo.Map(s, func(xx *SprintHistory, _ int) string {
		return xx.Slug
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

func (s SprintHistories) GetBySlug(slug string) SprintHistories {
	return lo.Filter(s, func(xx *SprintHistory, _ int) bool {
		return xx.Slug == slug
	})
}

func (s SprintHistories) GetBySlugs(slugs ...string) SprintHistories {
	return lo.Filter(s, func(xx *SprintHistory, _ int) bool {
		return lo.Contains(slugs, xx.Slug)
	})
}

func (s SprintHistories) SprintIDs() []uuid.UUID {
	return lo.Map(s, func(xx *SprintHistory, _ int) uuid.UUID {
		return xx.SprintID
	})
}

func (s SprintHistories) GetBySprintID(sprintID uuid.UUID) SprintHistories {
	return lo.Filter(s, func(xx *SprintHistory, _ int) bool {
		return xx.SprintID == sprintID
	})
}

func (s SprintHistories) GetBySprintIDs(sprintIDs ...uuid.UUID) SprintHistories {
	return lo.Filter(s, func(xx *SprintHistory, _ int) bool {
		return lo.Contains(sprintIDs, xx.SprintID)
	})
}

func (s SprintHistories) ToCSV() ([]string, [][]string) {
	return FieldDescs.Keys(), lo.Map(s, func(x *SprintHistory, _ int) []string {
		return x.Strings()
	})
}

func (s SprintHistories) Random() *SprintHistory {
	return util.RandomElement(s)
}

func (s SprintHistories) Clone() SprintHistories {
	return slices.Clone(s)
}
