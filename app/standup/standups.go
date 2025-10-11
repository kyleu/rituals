package standup

import (
	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

type Standups []*Standup

func (s Standups) Get(id uuid.UUID) *Standup {
	return lo.FindOrElse(s, nil, func(x *Standup) bool {
		return x.ID == id
	})
}

func (s Standups) IDs() []uuid.UUID {
	return lo.Map(s, func(xx *Standup, _ int) uuid.UUID {
		return xx.ID
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

func (s Standups) GetByID(id uuid.UUID) Standups {
	return lo.Filter(s, func(xx *Standup, _ int) bool {
		return xx.ID == id
	})
}

func (s Standups) GetByIDs(ids ...uuid.UUID) Standups {
	return lo.Filter(s, func(xx *Standup, _ int) bool {
		return lo.Contains(ids, xx.ID)
	})
}

func (s Standups) Slugs() []string {
	return lo.Map(s, func(xx *Standup, _ int) string {
		return xx.Slug
	})
}

func (s Standups) GetBySlug(slug string) Standups {
	return lo.Filter(s, func(xx *Standup, _ int) bool {
		return xx.Slug == slug
	})
}

func (s Standups) GetBySlugs(slugs ...string) Standups {
	return lo.Filter(s, func(xx *Standup, _ int) bool {
		return lo.Contains(slugs, xx.Slug)
	})
}

func (s Standups) Statuses() []enum.SessionStatus {
	return lo.Map(s, func(xx *Standup, _ int) enum.SessionStatus {
		return xx.Status
	})
}

func (s Standups) GetByStatus(status enum.SessionStatus) Standups {
	return lo.Filter(s, func(xx *Standup, _ int) bool {
		return xx.Status == status
	})
}

func (s Standups) GetByStatuses(statuses ...enum.SessionStatus) Standups {
	return lo.Filter(s, func(xx *Standup, _ int) bool {
		return lo.Contains(statuses, xx.Status)
	})
}

func (s Standups) TeamIDs() []*uuid.UUID {
	return lo.Map(s, func(xx *Standup, _ int) *uuid.UUID {
		return xx.TeamID
	})
}

func (s Standups) GetByTeamID(teamID *uuid.UUID) Standups {
	return lo.Filter(s, func(xx *Standup, _ int) bool {
		return xx.TeamID == teamID
	})
}

func (s Standups) GetByTeamIDs(teamIDs ...*uuid.UUID) Standups {
	return lo.Filter(s, func(xx *Standup, _ int) bool {
		return lo.Contains(teamIDs, xx.TeamID)
	})
}

func (s Standups) SprintIDs() []*uuid.UUID {
	return lo.Map(s, func(xx *Standup, _ int) *uuid.UUID {
		return xx.SprintID
	})
}

func (s Standups) GetBySprintID(sprintID *uuid.UUID) Standups {
	return lo.Filter(s, func(xx *Standup, _ int) bool {
		return xx.SprintID == sprintID
	})
}

func (s Standups) GetBySprintIDs(sprintIDs ...*uuid.UUID) Standups {
	return lo.Filter(s, func(xx *Standup, _ int) bool {
		return lo.Contains(sprintIDs, xx.SprintID)
	})
}

func (s Standups) ToMaps() []util.ValueMap {
	return lo.Map(s, func(xx *Standup, _ int) util.ValueMap {
		return xx.ToMap()
	})
}

func (s Standups) ToOrderedMaps() util.OrderedMaps[any] {
	return lo.Map(s, func(x *Standup, _ int) *util.OrderedMap[any] {
		return x.ToOrderedMap()
	})
}

func (s Standups) ToCSV() ([]string, [][]string) {
	return StandupFieldDescs.Keys(), lo.Map(s, func(x *Standup, _ int) []string {
		return x.Strings()
	})
}

func (s Standups) Random() *Standup {
	return util.RandomElement(s)
}

func (s Standups) Clone() Standups {
	return lo.Map(s, func(xx *Standup, _ int) *Standup {
		return xx.Clone()
	})
}
