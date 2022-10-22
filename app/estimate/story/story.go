// Content managed by Project Forge, see [projectforge.md] for details.
package story

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/slices"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

type Story struct {
	ID         uuid.UUID          `json:"id"`
	EstimateID uuid.UUID          `json:"estimateID"`
	Idx        int                `json:"idx"`
	UserID     uuid.UUID          `json:"userID"`
	Title      string             `json:"title"`
	Status     enum.SessionStatus `json:"status"`
	FinalVote  string             `json:"finalVote"`
	Created    time.Time          `json:"created"`
	Updated    *time.Time         `json:"updated,omitempty"`
}

func New(id uuid.UUID) *Story {
	return &Story{ID: id}
}

func Random() *Story {
	return &Story{
		ID:         util.UUID(),
		EstimateID: util.UUID(),
		Idx:        util.RandomInt(10000),
		UserID:     util.UUID(),
		Title:      util.RandomString(12),
		Status:     enum.SessionStatus(util.RandomString(12)),
		FinalVote:  util.RandomString(12),
		Created:    time.Now(),
		Updated:    util.NowPointer(),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*Story, error) {
	ret := &Story{}
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
	retEstimateID, e := m.ParseUUID("estimateID", true, true)
	if e != nil {
		return nil, e
	}
	if retEstimateID != nil {
		ret.EstimateID = *retEstimateID
	}
	ret.Idx, err = m.ParseInt("idx", true, true)
	if err != nil {
		return nil, err
	}
	retUserID, e := m.ParseUUID("userID", true, true)
	if e != nil {
		return nil, e
	}
	if retUserID != nil {
		ret.UserID = *retUserID
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
	ret.FinalVote, err = m.ParseString("finalVote", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (s *Story) Clone() *Story {
	return &Story{
		ID:         s.ID,
		EstimateID: s.EstimateID,
		Idx:        s.Idx,
		UserID:     s.UserID,
		Title:      s.Title,
		Status:     s.Status,
		FinalVote:  s.FinalVote,
		Created:    s.Created,
		Updated:    s.Updated,
	}
}

func (s *Story) String() string {
	return s.ID.String()
}

func (s *Story) TitleString() string {
	return s.Title
}

func (s *Story) WebPath() string {
	return "/estimate/story" + "/" + s.ID.String()
}

func (s *Story) Diff(sx *Story) util.Diffs {
	var diffs util.Diffs
	if s.ID != sx.ID {
		diffs = append(diffs, util.NewDiff("id", s.ID.String(), sx.ID.String()))
	}
	if s.EstimateID != sx.EstimateID {
		diffs = append(diffs, util.NewDiff("estimateID", s.EstimateID.String(), sx.EstimateID.String()))
	}
	if s.Idx != sx.Idx {
		diffs = append(diffs, util.NewDiff("idx", fmt.Sprint(s.Idx), fmt.Sprint(sx.Idx)))
	}
	if s.UserID != sx.UserID {
		diffs = append(diffs, util.NewDiff("userID", s.UserID.String(), sx.UserID.String()))
	}
	if s.Title != sx.Title {
		diffs = append(diffs, util.NewDiff("title", s.Title, sx.Title))
	}
	if s.Status != sx.Status {
		diffs = append(diffs, util.NewDiff("status", string(s.Status), string(sx.Status)))
	}
	if s.FinalVote != sx.FinalVote {
		diffs = append(diffs, util.NewDiff("finalVote", s.FinalVote, sx.FinalVote))
	}
	if s.Created != sx.Created {
		diffs = append(diffs, util.NewDiff("created", s.Created.String(), sx.Created.String()))
	}
	return diffs
}

func (s *Story) ToData() []any {
	return []any{s.ID, s.EstimateID, s.Idx, s.UserID, s.Title, s.Status, s.FinalVote, s.Created, s.Updated}
}

type Stories []*Story

func (s Stories) Get(id uuid.UUID) *Story {
	for _, x := range s {
		if x.ID == id {
			return x
		}
	}
	return nil
}

func (s Stories) Clone() Stories {
	return slices.Clone(s)
}
