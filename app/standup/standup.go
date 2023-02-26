// Content managed by Project Forge, see [projectforge.md] for details.
package standup

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

type Standup struct {
	ID       uuid.UUID          `json:"id"`
	Slug     string             `json:"slug"`
	Title    string             `json:"title"`
	Icon     string             `json:"icon"`
	Status   enum.SessionStatus `json:"status"`
	TeamID   *uuid.UUID         `json:"teamID,omitempty"`
	SprintID *uuid.UUID         `json:"sprintID,omitempty"`
	Created  time.Time          `json:"created"`
	Updated  *time.Time         `json:"updated,omitempty"`
}

func New(id uuid.UUID) *Standup {
	return &Standup{ID: id}
}

func Random() *Standup {
	return &Standup{
		ID:       util.UUID(),
		Slug:     util.RandomString(12),
		Title:    util.RandomString(12),
		Icon:     util.RandomString(12),
		Status:   enum.SessionStatus(util.RandomString(12)),
		TeamID:   util.UUIDP(),
		SprintID: util.UUIDP(),
		Created:  time.Now(),
		Updated:  util.NowPointer(),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*Standup, error) {
	ret := &Standup{}
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
	ret.SprintID, err = m.ParseUUID("sprintID", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (s *Standup) Clone() *Standup {
	return &Standup{
		ID:       s.ID,
		Slug:     s.Slug,
		Title:    s.Title,
		Icon:     s.Icon,
		Status:   s.Status,
		TeamID:   s.TeamID,
		SprintID: s.SprintID,
		Created:  s.Created,
		Updated:  s.Updated,
	}
}

func (s *Standup) String() string {
	return s.ID.String()
}

func (s *Standup) TitleString() string {
	return s.Title
}

func (s *Standup) WebPath() string {
	return "/admin/db/standup/" + s.ID.String()
}

func (s *Standup) Diff(sx *Standup) util.Diffs {
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
	if (s.SprintID == nil && sx.SprintID != nil) || (s.SprintID != nil && sx.SprintID == nil) || (s.SprintID != nil && sx.SprintID != nil && *s.SprintID != *sx.SprintID) { //nolint:lll
		diffs = append(diffs, util.NewDiff("sprintID", fmt.Sprint(s.SprintID), fmt.Sprint(sx.SprintID))) //nolint:gocritic // it's nullable
	}
	if s.Created != sx.Created {
		diffs = append(diffs, util.NewDiff("created", s.Created.String(), sx.Created.String()))
	}
	return diffs
}

func (s *Standup) ToData() []any {
	return []any{s.ID, s.Slug, s.Title, s.Icon, s.Status, s.TeamID, s.SprintID, s.Created, s.Updated}
}
