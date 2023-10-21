// Package standup - Content managed by Project Forge, see [projectforge.md] for details.
package standup

import (
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/enum"
)

type Standups []*Standup

func (s Standups) Get(id uuid.UUID) *Standup {
	return lo.FindOrElse(s, nil, func(x *Standup) bool {
		return x.ID == id
	})
}

func (s Standups) GetByIDs(ids ...uuid.UUID) Standups {
	return lo.Filter(s, func(xx *Standup, _ int) bool {
		return lo.Contains(ids, xx.ID)
	})
}

func (s Standups) GetByID(id uuid.UUID) Standups {
	return lo.Filter(s, func(xx *Standup, _ int) bool {
		return xx.ID == id
	})
}

func (s Standups) GetBySlugs(slugs ...string) Standups {
	return lo.Filter(s, func(xx *Standup, _ int) bool {
		return lo.Contains(slugs, xx.Slug)
	})
}

func (s Standups) GetBySlug(slug string) Standups {
	return lo.Filter(s, func(xx *Standup, _ int) bool {
		return xx.Slug == slug
	})
}

func (s Standups) GetByStatuses(statuses ...enum.SessionStatus) Standups {
	return lo.Filter(s, func(xx *Standup, _ int) bool {
		return lo.Contains(statuses, xx.Status)
	})
}

func (s Standups) GetByStatus(status enum.SessionStatus) Standups {
	return lo.Filter(s, func(xx *Standup, _ int) bool {
		return xx.Status == status
	})
}

func (s Standups) GetByTeamIDs(teamIDs ...*uuid.UUID) Standups {
	return lo.Filter(s, func(xx *Standup, _ int) bool {
		return lo.Contains(teamIDs, xx.TeamID)
	})
}

func (s Standups) GetByTeamID(teamID *uuid.UUID) Standups {
	return lo.Filter(s, func(xx *Standup, _ int) bool {
		return xx.TeamID == teamID
	})
}

func (s Standups) GetBySprintIDs(sprintIDs ...*uuid.UUID) Standups {
	return lo.Filter(s, func(xx *Standup, _ int) bool {
		return lo.Contains(sprintIDs, xx.SprintID)
	})
}

func (s Standups) GetBySprintID(sprintID *uuid.UUID) Standups {
	return lo.Filter(s, func(xx *Standup, _ int) bool {
		return xx.SprintID == sprintID
	})
}

func (s Standups) IDs() []uuid.UUID {
	return lo.Map(s, func(x *Standup, _ int) uuid.UUID {
		return x.ID
	})
}

func (s Standups) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(s, func(x *Standup, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (s Standups) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(s)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(s, func(x *Standup, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (s Standups) Clone() Standups {
	return slices.Clone(s)
}
