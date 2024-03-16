// Package team - Content managed by Project Forge, see [projectforge.md] for details.
package team

import (
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

type Teams []*Team

func (t Teams) Get(id uuid.UUID) *Team {
	return lo.FindOrElse(t, nil, func(x *Team) bool {
		return x.ID == id
	})
}

func (t Teams) IDs() []uuid.UUID {
	return lo.Map(t, func(xx *Team, _ int) uuid.UUID {
		return xx.ID
	})
}

func (t Teams) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(t)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(t, func(x *Team, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (t Teams) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(t)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(t, func(x *Team, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (t Teams) GetByID(id uuid.UUID) Teams {
	return lo.Filter(t, func(xx *Team, _ int) bool {
		return xx.ID == id
	})
}

func (t Teams) GetByIDs(ids ...uuid.UUID) Teams {
	return lo.Filter(t, func(xx *Team, _ int) bool {
		return lo.Contains(ids, xx.ID)
	})
}

func (t Teams) Slugs() []string {
	return lo.Map(t, func(xx *Team, _ int) string {
		return xx.Slug
	})
}

func (t Teams) GetBySlug(slug string) Teams {
	return lo.Filter(t, func(xx *Team, _ int) bool {
		return xx.Slug == slug
	})
}

func (t Teams) GetBySlugs(slugs ...string) Teams {
	return lo.Filter(t, func(xx *Team, _ int) bool {
		return lo.Contains(slugs, xx.Slug)
	})
}

func (t Teams) Statuses() []enum.SessionStatus {
	return lo.Map(t, func(xx *Team, _ int) enum.SessionStatus {
		return xx.Status
	})
}

func (t Teams) GetByStatus(status enum.SessionStatus) Teams {
	return lo.Filter(t, func(xx *Team, _ int) bool {
		return xx.Status == status
	})
}

func (t Teams) GetByStatuses(statuses ...enum.SessionStatus) Teams {
	return lo.Filter(t, func(xx *Team, _ int) bool {
		return lo.Contains(statuses, xx.Status)
	})
}

func (t Teams) ToCSV() ([]string, [][]string) {
	return FieldDescs.Keys(), lo.Map(t, func(x *Team, _ int) []string {
		return x.Strings()
	})
}

func (t Teams) Random() *Team {
	if len(t) == 0 {
		return nil
	}
	return t[util.RandomInt(len(t))]
}

func (t Teams) Clone() Teams {
	return slices.Clone(t)
}
