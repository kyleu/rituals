// Content managed by Project Forge, see [projectforge.md] for details.
package sprint

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/slices"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

type Sprint struct {
	ID      uuid.UUID          `json:"id"`
	Slug    string             `json:"slug"`
	Title   string             `json:"title"`
	Status  enum.SessionStatus `json:"status"`
	TeamID  *uuid.UUID         `json:"teamID,omitempty"`
	Owner   uuid.UUID          `json:"owner"`
	Created time.Time          `json:"created"`
	Updated *time.Time         `json:"updated,omitempty"`
}

func New(id uuid.UUID) *Sprint {
	return &Sprint{ID: id}
}

func Random() *Sprint {
	return &Sprint{
		ID:      util.UUID(),
		Slug:    util.RandomString(12),
		Title:   util.RandomString(12),
		Status:  enum.SessionStatus(util.RandomString(12)),
		TeamID:  util.UUIDP(),
		Owner:   util.UUID(),
		Created: time.Now(),
		Updated: util.NowPointer(),
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
	retStatus, err := m.ParseString("status", true, true)
	if err != nil {
		return nil, err
	}
	ret.Status = enum.SessionStatus(retStatus)
	ret.TeamID, err = m.ParseUUID("teamID", true, true)
	if err != nil {
		return nil, err
	}
	retOwner, e := m.ParseUUID("owner", true, true)
	if e != nil {
		return nil, e
	}
	if retOwner != nil {
		ret.Owner = *retOwner
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (s *Sprint) Clone() *Sprint {
	return &Sprint{
		ID:      s.ID,
		Slug:    s.Slug,
		Title:   s.Title,
		Status:  s.Status,
		TeamID:  s.TeamID,
		Owner:   s.Owner,
		Created: s.Created,
		Updated: s.Updated,
	}
}

func (s *Sprint) String() string {
	return s.ID.String()
}

func (s *Sprint) TitleString() string {
	return s.Title
}

func (s *Sprint) WebPath() string {
	return "/admin/db/sprint" + "/" + s.ID.String()
}

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
	if s.Status != sx.Status {
		diffs = append(diffs, util.NewDiff("status", string(s.Status), string(sx.Status)))
	}
	if (s.TeamID == nil && sx.TeamID != nil) || (s.TeamID != nil && sx.TeamID == nil) || (s.TeamID != nil && sx.TeamID != nil && *s.TeamID != *sx.TeamID) {
		diffs = append(diffs, util.NewDiff("teamID", fmt.Sprint(s.TeamID), fmt.Sprint(sx.TeamID))) //nolint:gocritic // it's nullable
	}
	if s.Owner != sx.Owner {
		diffs = append(diffs, util.NewDiff("owner", s.Owner.String(), sx.Owner.String()))
	}
	if s.Created != sx.Created {
		diffs = append(diffs, util.NewDiff("created", s.Created.String(), sx.Created.String()))
	}
	return diffs
}

func (s *Sprint) ToData() []any {
	return []any{s.ID, s.Slug, s.Title, s.Status, s.TeamID, s.Owner, s.Created, s.Updated}
}

type Sprints []*Sprint

func (s Sprints) Get(id uuid.UUID) *Sprint {
	for _, x := range s {
		if x.ID == id {
			return x
		}
	}
	return nil
}

func (s Sprints) Clone() Sprints {
	return slices.Clone(s)
}
