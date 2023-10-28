// Package sprint - Content managed by Project Forge, see [projectforge.md] for details.
package sprint

import (
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

type Sprints []*Sprint

func (s Sprints) Get(id uuid.UUID) *Sprint {
	return lo.FindOrElse(s, nil, func(x *Sprint) bool {
		return x.ID == id
	})
}

func (s Sprints) IDs() []uuid.UUID {
	return lo.Map(s, func(xx *Sprint, _ int) uuid.UUID {
		return xx.ID
	})
}

func (s Sprints) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(s, func(x *Sprint, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (s Sprints) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(s)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(s, func(x *Sprint, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (s Sprints) GetByID(id uuid.UUID) Sprints {
	return lo.Filter(s, func(xx *Sprint, _ int) bool {
		return xx.ID == id
	})
}

func (s Sprints) GetByIDs(ids ...uuid.UUID) Sprints {
	return lo.Filter(s, func(xx *Sprint, _ int) bool {
		return lo.Contains(ids, xx.ID)
	})
}

func (s Sprints) Slugs() []string {
	return lo.Map(s, func(xx *Sprint, _ int) string {
		return xx.Slug
	})
}

func (s Sprints) GetBySlug(slug string) Sprints {
	return lo.Filter(s, func(xx *Sprint, _ int) bool {
		return xx.Slug == slug
	})
}

func (s Sprints) GetBySlugs(slugs ...string) Sprints {
	return lo.Filter(s, func(xx *Sprint, _ int) bool {
		return lo.Contains(slugs, xx.Slug)
	})
}

func (s Sprints) Statuses() []enum.SessionStatus {
	return lo.Map(s, func(xx *Sprint, _ int) enum.SessionStatus {
		return xx.Status
	})
}

func (s Sprints) GetByStatus(status enum.SessionStatus) Sprints {
	return lo.Filter(s, func(xx *Sprint, _ int) bool {
		return xx.Status == status
	})
}

func (s Sprints) GetByStatuses(statuses ...enum.SessionStatus) Sprints {
	return lo.Filter(s, func(xx *Sprint, _ int) bool {
		return lo.Contains(statuses, xx.Status)
	})
}

func (s Sprints) TeamIDs() []*uuid.UUID {
	return lo.Map(s, func(xx *Sprint, _ int) *uuid.UUID {
		return xx.TeamID
	})
}

func (s Sprints) GetByTeamID(teamID *uuid.UUID) Sprints {
	return lo.Filter(s, func(xx *Sprint, _ int) bool {
		return xx.TeamID == teamID
	})
}

func (s Sprints) GetByTeamIDs(teamIDs ...*uuid.UUID) Sprints {
	return lo.Filter(s, func(xx *Sprint, _ int) bool {
		return lo.Contains(teamIDs, xx.TeamID)
	})
}

func (s Sprints) Random() *Sprint {
	if len(s) == 0 {
		return nil
	}
	return s[util.RandomInt(len(s))]
}

func (s Sprints) Clone() Sprints {
	return slices.Clone(s)
}
