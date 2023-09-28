// Package estimate - Content managed by Project Forge, see [projectforge.md] for details.
package estimate

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

type Estimate struct {
	ID       uuid.UUID          `json:"id"`
	Slug     string             `json:"slug"`
	Title    string             `json:"title"`
	Icon     string             `json:"icon"`
	Status   enum.SessionStatus `json:"status"`
	TeamID   *uuid.UUID         `json:"teamID,omitempty"`
	SprintID *uuid.UUID         `json:"sprintID,omitempty"`
	Choices  []string           `json:"choices"`
	Created  time.Time          `json:"created"`
	Updated  *time.Time         `json:"updated,omitempty"`
}

func New(id uuid.UUID) *Estimate {
	return &Estimate{ID: id}
}

func Random() *Estimate {
	return &Estimate{
		ID:       util.UUID(),
		Slug:     util.RandomString(12),
		Title:    util.RandomString(12),
		Icon:     util.RandomString(12),
		Status:   enum.SessionStatus(util.RandomString(12)),
		TeamID:   util.UUIDP(),
		SprintID: util.UUIDP(),
		Choices:  nil,
		Created:  util.TimeCurrent(),
		Updated:  util.TimeCurrentP(),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*Estimate, error) {
	ret := &Estimate{}
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
	ret.Choices, err = m.ParseArrayString("choices", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (e *Estimate) Clone() *Estimate {
	return &Estimate{e.ID, e.Slug, e.Title, e.Icon, e.Status, e.TeamID, e.SprintID, e.Choices, e.Created, e.Updated}
}

func (e *Estimate) String() string {
	return e.ID.String()
}

func (e *Estimate) TitleString() string {
	return e.Title
}

func (e *Estimate) WebPath() string {
	return "/admin/db/estimate/" + e.ID.String()
}

//nolint:lll,gocognit
func (e *Estimate) Diff(ex *Estimate) util.Diffs {
	var diffs util.Diffs
	if e.ID != ex.ID {
		diffs = append(diffs, util.NewDiff("id", e.ID.String(), ex.ID.String()))
	}
	if e.Slug != ex.Slug {
		diffs = append(diffs, util.NewDiff("slug", e.Slug, ex.Slug))
	}
	if e.Title != ex.Title {
		diffs = append(diffs, util.NewDiff("title", e.Title, ex.Title))
	}
	if e.Icon != ex.Icon {
		diffs = append(diffs, util.NewDiff("icon", e.Icon, ex.Icon))
	}
	if e.Status != ex.Status {
		diffs = append(diffs, util.NewDiff("status", string(e.Status), string(ex.Status)))
	}
	if (e.TeamID == nil && ex.TeamID != nil) || (e.TeamID != nil && ex.TeamID == nil) || (e.TeamID != nil && ex.TeamID != nil && *e.TeamID != *ex.TeamID) {
		diffs = append(diffs, util.NewDiff("teamID", fmt.Sprint(e.TeamID), fmt.Sprint(ex.TeamID))) //nolint:gocritic // it's nullable
	}
	if (e.SprintID == nil && ex.SprintID != nil) || (e.SprintID != nil && ex.SprintID == nil) || (e.SprintID != nil && ex.SprintID != nil && *e.SprintID != *ex.SprintID) {
		diffs = append(diffs, util.NewDiff("sprintID", fmt.Sprint(e.SprintID), fmt.Sprint(ex.SprintID))) //nolint:gocritic // it's nullable
	}
	diffs = append(diffs, util.DiffObjects(e.Choices, ex.Choices, "choices")...)
	if e.Created != ex.Created {
		diffs = append(diffs, util.NewDiff("created", e.Created.String(), ex.Created.String()))
	}
	return diffs
}

func (e *Estimate) ToData() []any {
	return []any{e.ID, e.Slug, e.Title, e.Icon, e.Status, e.TeamID, e.SprintID, e.Choices, e.Created, e.Updated}
}
