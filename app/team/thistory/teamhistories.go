// Package thistory - Content managed by Project Forge, see [projectforge.md] for details.
package thistory

import (
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

type TeamHistories []*TeamHistory

func (t TeamHistories) Get(slug string) *TeamHistory {
	return lo.FindOrElse(t, nil, func(x *TeamHistory) bool {
		return x.Slug == slug
	})
}

func (t TeamHistories) Slugs() []string {
	return lo.Map(t, func(xx *TeamHistory, _ int) string {
		return xx.Slug
	})
}

func (t TeamHistories) SlugStrings(includeNil bool) []string {
	ret := make([]string, 0, len(t)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(t, func(x *TeamHistory, _ int) {
		ret = append(ret, x.Slug)
	})
	return ret
}

func (t TeamHistories) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(t)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(t, func(x *TeamHistory, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (t TeamHistories) GetBySlug(slug string) TeamHistories {
	return lo.Filter(t, func(xx *TeamHistory, _ int) bool {
		return xx.Slug == slug
	})
}

func (t TeamHistories) GetBySlugs(slugs ...string) TeamHistories {
	return lo.Filter(t, func(xx *TeamHistory, _ int) bool {
		return lo.Contains(slugs, xx.Slug)
	})
}

func (t TeamHistories) TeamIDs() []uuid.UUID {
	return lo.Map(t, func(xx *TeamHistory, _ int) uuid.UUID {
		return xx.TeamID
	})
}

func (t TeamHistories) GetByTeamID(teamID uuid.UUID) TeamHistories {
	return lo.Filter(t, func(xx *TeamHistory, _ int) bool {
		return xx.TeamID == teamID
	})
}

func (t TeamHistories) GetByTeamIDs(teamIDs ...uuid.UUID) TeamHistories {
	return lo.Filter(t, func(xx *TeamHistory, _ int) bool {
		return lo.Contains(teamIDs, xx.TeamID)
	})
}

func (t TeamHistories) ToCSV() ([]string, [][]string) {
	return FieldDescs.Keys(), lo.Map(t, func(x *TeamHistory, _ int) []string {
		return x.Strings()
	})
}

func (t TeamHistories) Random() *TeamHistory {
	return util.RandomElement(t)
}

func (t TeamHistories) Clone() TeamHistories {
	return slices.Clone(t)
}
