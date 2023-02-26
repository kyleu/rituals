// Content managed by Project Forge, see [projectforge.md] for details.
package sprint

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

type Sprint struct {
	ID        uuid.UUID          `json:"id"`
	Slug      string             `json:"slug"`
	Title     string             `json:"title"`
	Icon      string             `json:"icon"`
	Status    enum.SessionStatus `json:"status"`
	TeamID    *uuid.UUID         `json:"teamID,omitempty"`
	StartDate *time.Time         `json:"startDate,omitempty"`
	EndDate   *time.Time         `json:"endDate,omitempty"`
	Created   time.Time          `json:"created"`
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
		Status:    enum.SessionStatus(util.RandomString(12)),
		TeamID:    util.UUIDP(),
		StartDate: util.NowPointer(),
		EndDate:   util.NowPointer(),
		Created:   time.Now(),
		Updated:   util.NowPointer(),
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
	ret.Status = enum.SessionStatus(retStatus)
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
	return &Sprint{
		ID:        s.ID,
		Slug:      s.Slug,
		Title:     s.Title,
		Icon:      s.Icon,
		Status:    s.Status,
		TeamID:    s.TeamID,
		StartDate: s.StartDate,
		EndDate:   s.EndDate,
		Created:   s.Created,
		Updated:   s.Updated,
	}
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

//nolint:gocognit
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
		diffs = append(diffs, util.NewDiff("status", string(s.Status), string(sx.Status)))
	}
	if (s.TeamID == nil && sx.TeamID != nil) || (s.TeamID != nil && sx.TeamID == nil) || (s.TeamID != nil && sx.TeamID != nil && *s.TeamID != *sx.TeamID) {
		diffs = append(diffs, util.NewDiff("teamID", fmt.Sprint(s.TeamID), fmt.Sprint(sx.TeamID))) //nolint:gocritic // it's nullable
	}
	if (s.StartDate == nil && sx.StartDate != nil) || (s.StartDate != nil && sx.StartDate == nil) || (s.StartDate != nil && sx.StartDate != nil && *s.StartDate != *sx.StartDate) { //nolint:lll
		diffs = append(diffs, util.NewDiff("startDate", fmt.Sprint(s.StartDate), fmt.Sprint(sx.StartDate))) //nolint:gocritic // it's nullable
	}
	if (s.EndDate == nil && sx.EndDate != nil) || (s.EndDate != nil && sx.EndDate == nil) || (s.EndDate != nil && sx.EndDate != nil && *s.EndDate != *sx.EndDate) { //nolint:lll
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
