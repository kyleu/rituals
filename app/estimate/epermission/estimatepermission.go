package epermission

import (
	"fmt"
	"net/url"
	"path"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

const DefaultRoute = "/admin/db/estimate/permission"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return path.Join(paths...)
}

var _ svc.Model = (*EstimatePermission)(nil)

type PK struct {
	EstimateID uuid.UUID `json:"estimateID,omitempty"`
	Key        string    `json:"key,omitempty"`
	Value      string    `json:"value,omitempty"`
}

func (p *PK) String() string {
	return fmt.Sprintf("%v • %s • %s", p.EstimateID, p.Key, p.Value)
}

type EstimatePermission struct {
	EstimateID uuid.UUID `json:"estimateID,omitempty"`
	Key        string    `json:"key,omitempty"`
	Value      string    `json:"value,omitempty"`
	Access     string    `json:"access,omitempty"`
	Created    time.Time `json:"created,omitempty"`
}

func NewEstimatePermission(estimateID uuid.UUID, key string, value string) *EstimatePermission {
	return &EstimatePermission{EstimateID: estimateID, Key: key, Value: value}
}

func (e *EstimatePermission) Clone() *EstimatePermission {
	return &EstimatePermission{e.EstimateID, e.Key, e.Value, e.Access, e.Created}
}

func (e *EstimatePermission) String() string {
	return fmt.Sprintf("%s • %s • %s", e.EstimateID.String(), e.Key, e.Value)
}

func (e *EstimatePermission) TitleString() string {
	return e.String()
}

func (e *EstimatePermission) ToPK() *PK {
	return &PK{
		EstimateID: e.EstimateID,
		Key:        e.Key,
		Value:      e.Value,
	}
}

func RandomEstimatePermission() *EstimatePermission {
	return &EstimatePermission{
		EstimateID: util.UUID(),
		Key:        util.RandomString(12),
		Value:      util.RandomString(12),
		Access:     util.RandomString(12),
		Created:    util.TimeCurrent(),
	}
}

func (e *EstimatePermission) Strings() []string {
	return []string{e.EstimateID.String(), e.Key, e.Value, e.Access, util.TimeToFull(&e.Created)}
}

func (e *EstimatePermission) ToCSV() ([]string, [][]string) {
	return EstimatePermissionFieldDescs.Keys(), [][]string{e.Strings()}
}

func (e *EstimatePermission) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return path.Join(append(paths, url.QueryEscape(e.EstimateID.String()), url.QueryEscape(e.Key), url.QueryEscape(e.Value))...)
}

func (e *EstimatePermission) Breadcrumb(extra ...string) string {
	return e.TitleString() + "||" + e.WebPath(extra...) + "**permission"
}

func (e *EstimatePermission) ToData() []any {
	return []any{e.EstimateID, e.Key, e.Value, e.Access, e.Created}
}

var EstimatePermissionFieldDescs = util.FieldDescs{
	{Key: "estimateID", Title: "Estimate ID", Description: "", Type: "uuid"},
	{Key: "key", Title: "Key", Description: "", Type: "string"},
	{Key: "value", Title: "Value", Description: "", Type: "string"},
	{Key: "access", Title: "Access", Description: "", Type: "string"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
}
