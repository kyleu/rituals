// Package sprint - Content managed by Project Forge, see [projectforge.md] for details.
package sprint

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

type Sprint struct {
	ID        uuid.UUID          `json:"id,omitempty"`
	Slug      string             `json:"slug,omitempty"`
	Title     string             `json:"title,omitempty"`
	Icon      string             `json:"icon,omitempty"`
	Status    enum.SessionStatus `json:"status,omitempty"`
	TeamID    *uuid.UUID         `json:"teamID,omitempty"`
	StartDate *time.Time         `json:"startDate,omitempty"`
	EndDate   *time.Time         `json:"endDate,omitempty"`
	Created   time.Time          `json:"created,omitempty"`
	Updated   *time.Time         `json:"updated,omitempty"`
}

func New(id uuid.UUID) *Sprint {
	return &Sprint{ID: id}
}

func Random() *Sprint {
	return &Sprint{
		ID:        util.UUID(),
		Slug:      util.RandomString(12),
		Title:     util.RandomString(12),
		Icon:      util.RandomString(12),
		Status:    enum.AllSessionStatuses.Random(),
		TeamID:    util.UUIDP(),
		StartDate: util.TimeCurrentP(),
		EndDate:   util.TimeCurrentP(),
		Created:   util.TimeCurrent(),
		Updated:   util.TimeCurrentP(),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*Sprint, error) {
	ret := &Sprint{}
	var err error
	if setPK {
		retID, e := m.ParseUUID("id", true, true)
		if e != nil {
			return nil, e
		}
		if retID != nil {
			ret.ID = *retID
		}
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
	}
	ret.Slug, err = m.ParseString("slug", true, true)
	if err != nil {
		return nil, err
	}
	ret.Title, err = m.ParseString("title", true, true)
	if err != nil {
		return nil, err
	}
	ret.Icon, err = m.ParseString("icon", true, true)
	if err != nil {
		return nil, err
	}
	retStatus, err := m.ParseString("status", true, true)
	if err != nil {
		return nil, err
	}
	ret.Status = enum.AllSessionStatuses.Get(retStatus, nil)
	ret.TeamID, err = m.ParseUUID("teamID", true, true)
	if err != nil {
		return nil, err
	}
	ret.StartDate, err = m.ParseTime("startDate", true, true)
	if err != nil {
		return nil, err
	}
	ret.EndDate, err = m.ParseTime("endDate", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (s *Sprint) Clone() *Sprint {
	return &Sprint{s.ID, s.Slug, s.Title, s.Icon, s.Status, s.TeamID, s.StartDate, s.EndDate, s.Created, s.Updated}
}

func (s *Sprint) String() string {
	return s.ID.String()
}

func (s *Sprint) TitleString() string {
	return s.Title
}

func (s *Sprint) WebPath() string {
	return "/admin/db/sprint/" + s.ID.String()
}

//nolint:lll,gocognit
func (s *Sprint) Diff(sx *Sprint) util.Diffs {
	var diffs util.Diffs
	if s.ID != sx.ID {
		diffs = append(diffs, util.NewDiff("id", s.ID.String(), sx.ID.String()))
	}
	if s.Slug != sx.Slug {
		diffs = append(diffs, util.NewDiff("slug", s.Slug, sx.Slug))
	}
	if s.Title != sx.Title {
		diffs = append(diffs, util.NewDiff("title", s.Title, sx.Title))
	}
	if s.Icon != sx.Icon {
		diffs = append(diffs, util.NewDiff("icon", s.Icon, sx.Icon))
	}
	if s.Status != sx.Status {
		diffs = append(diffs, util.NewDiff("status", s.Status.Key, sx.Status.Key))
	}
	if (s.TeamID == nil && sx.TeamID != nil) || (s.TeamID != nil && sx.TeamID == nil) || (s.TeamID != nil && sx.TeamID != nil && *s.TeamID != *sx.TeamID) {
		diffs = append(diffs, util.NewDiff("teamID", fmt.Sprint(s.TeamID), fmt.Sprint(sx.TeamID))) //nolint:gocritic // it's nullable
	}
	if (s.StartDate == nil && sx.StartDate != nil) || (s.StartDate != nil && sx.StartDate == nil) || (s.StartDate != nil && sx.StartDate != nil && *s.StartDate != *sx.StartDate) {
		diffs = append(diffs, util.NewDiff("startDate", fmt.Sprint(s.StartDate), fmt.Sprint(sx.StartDate))) //nolint:gocritic // it's nullable
	}
	if (s.EndDate == nil && sx.EndDate != nil) || (s.EndDate != nil && sx.EndDate == nil) || (s.EndDate != nil && sx.EndDate != nil && *s.EndDate != *sx.EndDate) {
		diffs = append(diffs, util.NewDiff("endDate", fmt.Sprint(s.EndDate), fmt.Sprint(sx.EndDate))) //nolint:gocritic // it's nullable
	}
	if s.Created != sx.Created {
		diffs = append(diffs, util.NewDiff("created", s.Created.String(), sx.Created.String()))
	}
	return diffs
}

func (s *Sprint) ToData() []any {
	return []any{s.ID, s.Slug, s.Title, s.Icon, s.Status, s.TeamID, s.StartDate, s.EndDate, s.Created, s.Updated}
}
