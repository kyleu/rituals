// Content managed by Project Forge, see [projectforge.md] for details.
package epermission

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

type PK struct {
	EstimateID uuid.UUID `json:"estimateID"`
	Key        string    `json:"key"`
	Value      string    `json:"value"`
}

type EstimatePermission struct {
	EstimateID uuid.UUID `json:"estimateID"`
	Key        string    `json:"key"`
	Value      string    `json:"value"`
	Access     string    `json:"access"`
	Created    time.Time `json:"created"`
}

func New(estimateID uuid.UUID, key string, value string) *EstimatePermission {
	return &EstimatePermission{EstimateID: estimateID, Key: key, Value: value}
}

func Random() *EstimatePermission {
	return &EstimatePermission{
		EstimateID: util.UUID(),
		Key:        util.RandomString(12),
		Value:      util.RandomString(12),
		Access:     util.RandomString(12),
		Created:    time.Now(),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*EstimatePermission, error) {
	ret := &EstimatePermission{}
	var err error
	if setPK {
		retEstimateID, e := m.ParseUUID("estimateID", true, true)
		if e != nil {
			return nil, e
		}
		if retEstimateID != nil {
			ret.EstimateID = *retEstimateID
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

func (e *EstimatePermission) Clone() *EstimatePermission {
	return &EstimatePermission{
		EstimateID: e.EstimateID,
		Key:        e.Key,
		Value:      e.Value,
		Access:     e.Access,
		Created:    e.Created,
	}
}

func (e *EstimatePermission) String() string {
	return fmt.Sprintf("%s::%s::%s", e.EstimateID.String(), e.Key, e.Value)
}

func (e *EstimatePermission) TitleString() string {
	return e.String()
}

func (e *EstimatePermission) WebPath() string {
	return "/admin/db/estimate/permission/" + e.EstimateID.String() + "/" + e.Key + "/" + e.Value
}

func (e *EstimatePermission) Diff(ex *EstimatePermission) util.Diffs {
	var diffs util.Diffs
	if e.EstimateID != ex.EstimateID {
		diffs = append(diffs, util.NewDiff("estimateID", e.EstimateID.String(), ex.EstimateID.String()))
	}
	if e.Key != ex.Key {
		diffs = append(diffs, util.NewDiff("key", e.Key, ex.Key))
	}
	if e.Value != ex.Value {
		diffs = append(diffs, util.NewDiff("value", e.Value, ex.Value))
	}
	if e.Access != ex.Access {
		diffs = append(diffs, util.NewDiff("access", e.Access, ex.Access))
	}
	if e.Created != ex.Created {
		diffs = append(diffs, util.NewDiff("created", e.Created.String(), ex.Created.String()))
	}
	return diffs
}

func (e *EstimatePermission) ToData() []any {
	return []any{e.EstimateID, e.Key, e.Value, e.Access, e.Created}
}
