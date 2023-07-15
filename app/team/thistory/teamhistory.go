// Content managed by Project Forge, see [projectforge.md] for details.
package thistory

import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

type TeamHistory struct {
	Slug     string    `json:"slug"`
	TeamID   uuid.UUID `json:"teamID"`
	TeamName string    `json:"teamName"`
	Created  time.Time `json:"created"`
}

func New(slug string) *TeamHistory {
	return &TeamHistory{Slug: slug}
}

func Random() *TeamHistory {
	return &TeamHistory{
		Slug:     util.RandomString(12),
		TeamID:   util.UUID(),
		TeamName: util.RandomString(12),
		Created:  util.TimeCurrent(),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*TeamHistory, error) {
	ret := &TeamHistory{}
	var err error
	if setPK {
		ret.Slug, err = m.ParseString("slug", true, true)
		if err != nil {
			return nil, err
		}
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
	}
	retTeamID, e := m.ParseUUID("teamID", true, true)
	if e != nil {
		return nil, e
	}
	if retTeamID != nil {
		ret.TeamID = *retTeamID
	}
	ret.TeamName, err = m.ParseString("teamName", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (t *TeamHistory) Clone() *TeamHistory {
	return &TeamHistory{t.Slug, t.TeamID, t.TeamName, t.Created}
}

func (t *TeamHistory) String() string {
	return t.Slug
}

func (t *TeamHistory) TitleString() string {
	return t.String()
}

func (t *TeamHistory) WebPath() string {
	return "/admin/db/team/history/" + t.Slug
}

func (t *TeamHistory) Diff(tx *TeamHistory) util.Diffs {
	var diffs util.Diffs
	if t.Slug != tx.Slug {
		diffs = append(diffs, util.NewDiff("slug", t.Slug, tx.Slug))
	}
	if t.TeamID != tx.TeamID {
		diffs = append(diffs, util.NewDiff("teamID", t.TeamID.String(), tx.TeamID.String()))
	}
	if t.TeamName != tx.TeamName {
		diffs = append(diffs, util.NewDiff("teamName", t.TeamName, tx.TeamName))
	}
	if t.Created != tx.Created {
		diffs = append(diffs, util.NewDiff("created", t.Created.String(), tx.Created.String()))
	}
	return diffs
}

func (t *TeamHistory) ToData() []any {
	return []any{t.Slug, t.TeamID, t.TeamName, t.Created}
}
