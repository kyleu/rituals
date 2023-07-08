// Content managed by Project Forge, see [projectforge.md] for details.
package retro

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

type Retro struct {
	ID         uuid.UUID          `json:"id"`
	Slug       string             `json:"slug"`
	Title      string             `json:"title"`
	Icon       string             `json:"icon"`
	Status     enum.SessionStatus `json:"status"`
	TeamID     *uuid.UUID         `json:"teamID,omitempty"`
	SprintID   *uuid.UUID         `json:"sprintID,omitempty"`
	Categories []string           `json:"categories"`
	Created    time.Time          `json:"created"`
	Updated    *time.Time         `json:"updated,omitempty"`
}

func New(id uuid.UUID) *Retro {
	return &Retro{ID: id}
}

func Random() *Retro {
	return &Retro{
		ID:         util.UUID(),
		Slug:       util.RandomString(12),
		Title:      util.RandomString(12),
		Icon:       util.RandomString(12),
		Status:     enum.SessionStatus(util.RandomString(12)),
		TeamID:     util.UUIDP(),
		SprintID:   util.UUIDP(),
		Categories: nil,
		Created:    util.TimeCurrent(),
		Updated:    util.TimeCurrentP(),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*Retro, error) {
	ret := &Retro{}
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
	ret.Categories, err = m.ParseArrayString("categories", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (r *Retro) Clone() *Retro {
	return &Retro{
		ID:         r.ID,
		Slug:       r.Slug,
		Title:      r.Title,
		Icon:       r.Icon,
		Status:     r.Status,
		TeamID:     r.TeamID,
		SprintID:   r.SprintID,
		Categories: r.Categories,
		Created:    r.Created,
		Updated:    r.Updated,
	}
}

func (r *Retro) String() string {
	return r.ID.String()
}

func (r *Retro) TitleString() string {
	return r.Title
}

func (r *Retro) WebPath() string {
	return "/admin/db/retro/" + r.ID.String()
}

func (r *Retro) Diff(rx *Retro) util.Diffs {
	var diffs util.Diffs
	if r.ID != rx.ID {
		diffs = append(diffs, util.NewDiff("id", r.ID.String(), rx.ID.String()))
	}
	if r.Slug != rx.Slug {
		diffs = append(diffs, util.NewDiff("slug", r.Slug, rx.Slug))
	}
	if r.Title != rx.Title {
		diffs = append(diffs, util.NewDiff("title", r.Title, rx.Title))
	}
	if r.Icon != rx.Icon {
		diffs = append(diffs, util.NewDiff("icon", r.Icon, rx.Icon))
	}
	if r.Status != rx.Status {
		diffs = append(diffs, util.NewDiff("status", string(r.Status), string(rx.Status)))
	}
	if (r.TeamID == nil && rx.TeamID != nil) || (r.TeamID != nil && rx.TeamID == nil) || (r.TeamID != nil && rx.TeamID != nil && *r.TeamID != *rx.TeamID) {
		diffs = append(diffs, util.NewDiff("teamID", fmt.Sprint(r.TeamID), fmt.Sprint(rx.TeamID))) //nolint:gocritic // it's nullable
	}
	if (r.SprintID == nil && rx.SprintID != nil) || (r.SprintID != nil && rx.SprintID == nil) || (r.SprintID != nil && rx.SprintID != nil && *r.SprintID != *rx.SprintID) { //nolint:lll
		diffs = append(diffs, util.NewDiff("sprintID", fmt.Sprint(r.SprintID), fmt.Sprint(rx.SprintID))) //nolint:gocritic // it's nullable
	}
	diffs = append(diffs, util.DiffObjects(r.Categories, rx.Categories, "categories")...)
	if r.Created != rx.Created {
		diffs = append(diffs, util.NewDiff("created", r.Created.String(), rx.Created.String()))
	}
	return diffs
}

func (r *Retro) ToData() []any {
	return []any{r.ID, r.Slug, r.Title, r.Icon, r.Status, r.TeamID, r.SprintID, r.Categories, r.Created, r.Updated}
}
