package retro

import (
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

type Retros []*Retro

func (r Retros) Get(id uuid.UUID) *Retro {
	return lo.FindOrElse(r, nil, func(x *Retro) bool {
		return x.ID == id
	})
}

func (r Retros) IDs() []uuid.UUID {
	return lo.Map(r, func(xx *Retro, _ int) uuid.UUID {
		return xx.ID
	})
}

func (r Retros) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(r)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(r, func(x *Retro, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (r Retros) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(r)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(r, func(x *Retro, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (r Retros) GetByID(id uuid.UUID) Retros {
	return lo.Filter(r, func(xx *Retro, _ int) bool {
		return xx.ID == id
	})
}

func (r Retros) GetByIDs(ids ...uuid.UUID) Retros {
	return lo.Filter(r, func(xx *Retro, _ int) bool {
		return lo.Contains(ids, xx.ID)
	})
}

func (r Retros) Slugs() []string {
	return lo.Map(r, func(xx *Retro, _ int) string {
		return xx.Slug
	})
}

func (r Retros) GetBySlug(slug string) Retros {
	return lo.Filter(r, func(xx *Retro, _ int) bool {
		return xx.Slug == slug
	})
}

func (r Retros) GetBySlugs(slugs ...string) Retros {
	return lo.Filter(r, func(xx *Retro, _ int) bool {
		return lo.Contains(slugs, xx.Slug)
	})
}

func (r Retros) Statuses() []enum.SessionStatus {
	return lo.Map(r, func(xx *Retro, _ int) enum.SessionStatus {
		return xx.Status
	})
}

func (r Retros) GetByStatus(status enum.SessionStatus) Retros {
	return lo.Filter(r, func(xx *Retro, _ int) bool {
		return xx.Status == status
	})
}

func (r Retros) GetByStatuses(statuses ...enum.SessionStatus) Retros {
	return lo.Filter(r, func(xx *Retro, _ int) bool {
		return lo.Contains(statuses, xx.Status)
	})
}

func (r Retros) TeamIDs() []*uuid.UUID {
	return lo.Map(r, func(xx *Retro, _ int) *uuid.UUID {
		return xx.TeamID
	})
}

func (r Retros) GetByTeamID(teamID *uuid.UUID) Retros {
	return lo.Filter(r, func(xx *Retro, _ int) bool {
		return xx.TeamID == teamID
	})
}

func (r Retros) GetByTeamIDs(teamIDs ...*uuid.UUID) Retros {
	return lo.Filter(r, func(xx *Retro, _ int) bool {
		return lo.Contains(teamIDs, xx.TeamID)
	})
}

func (r Retros) SprintIDs() []*uuid.UUID {
	return lo.Map(r, func(xx *Retro, _ int) *uuid.UUID {
		return xx.SprintID
	})
}

func (r Retros) GetBySprintID(sprintID *uuid.UUID) Retros {
	return lo.Filter(r, func(xx *Retro, _ int) bool {
		return xx.SprintID == sprintID
	})
}

func (r Retros) GetBySprintIDs(sprintIDs ...*uuid.UUID) Retros {
	return lo.Filter(r, func(xx *Retro, _ int) bool {
		return lo.Contains(sprintIDs, xx.SprintID)
	})
}

func (r Retros) ToCSV() ([]string, [][]string) {
	return RetroFieldDescs.Keys(), lo.Map(r, func(x *Retro, _ int) []string {
		return x.Strings()
	})
}

func (r Retros) Random() *Retro {
	return util.RandomElement(r)
}

func (r Retros) Clone() Retros {
	return slices.Clone(r)
}
