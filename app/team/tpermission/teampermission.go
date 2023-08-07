// Content managed by Project Forge, see [projectforge.md] for details.
package tpermission

import (
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

type PK struct {
	TeamID uuid.UUID `json:"teamID"`
	Key    string    `json:"key"`
	Value  string    `json:"value"`
}

type TeamPermission struct {
	TeamID  uuid.UUID `json:"teamID"`
	Key     string    `json:"key"`
	Value   string    `json:"value"`
	Access  string    `json:"access"`
	Created time.Time `json:"created"`
}

func New(teamID uuid.UUID, key string, value string) *TeamPermission {
	return &TeamPermission{TeamID: teamID, Key: key, Value: value}
}

func Random() *TeamPermission {
	return &TeamPermission{
		TeamID:  util.UUID(),
		Key:     util.RandomString(12),
		Value:   util.RandomString(12),
		Access:  util.RandomString(12),
		Created: util.TimeCurrent(),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*TeamPermission, error) {
	ret := &TeamPermission{}
	var err error
	if setPK {
		retTeamID, e := m.ParseUUID("teamID", true, true)
		if e != nil {
			return nil, e
		}
		if retTeamID != nil {
			ret.TeamID = *retTeamID
		}
		ret.Key, err = m.ParseString("key", true, true)
		if err != nil {
			return nil, err
		}
		ret.Value, err = m.ParseString("value", true, true)
		if err != nil {
			return nil, err
		}
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
	}
	ret.Access, err = m.ParseString("access", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (t *TeamPermission) Clone() *TeamPermission {
	return &TeamPermission{t.TeamID, t.Key, t.Value, t.Access, t.Created}
}

func (t *TeamPermission) String() string {
	return fmt.Sprintf("%s::%s::%s", t.TeamID.String(), t.Key, t.Value)
}

func (t *TeamPermission) TitleString() string {
	return t.String()
}

func (t *TeamPermission) ToPK() *PK {
	return &PK{
		TeamID: t.TeamID,
		Key:    t.Key,
		Value:  t.Value,
	}
}

func (t *TeamPermission) WebPath() string {
	return "/admin/db/team/permission/" + t.TeamID.String() + "/" + url.QueryEscape(t.Key) + "/" + url.QueryEscape(t.Value)
}

func (t *TeamPermission) Diff(tx *TeamPermission) util.Diffs {
	var diffs util.Diffs
	if t.TeamID != tx.TeamID {
		diffs = append(diffs, util.NewDiff("teamID", t.TeamID.String(), tx.TeamID.String()))
	}
	if t.Key != tx.Key {
		diffs = append(diffs, util.NewDiff("key", t.Key, tx.Key))
	}
	if t.Value != tx.Value {
		diffs = append(diffs, util.NewDiff("value", t.Value, tx.Value))
	}
	if t.Access != tx.Access {
		diffs = append(diffs, util.NewDiff("access", t.Access, tx.Access))
	}
	if t.Created != tx.Created {
		diffs = append(diffs, util.NewDiff("created", t.Created.String(), tx.Created.String()))
	}
	return diffs
}

func (t *TeamPermission) ToData() []any {
	return []any{t.TeamID, t.Key, t.Value, t.Access, t.Created}
}
