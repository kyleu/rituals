package action

import (
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

const DefaultRoute = "/admin/db/action"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(paths...)
}

var _ svc.Model = (*Action)(nil)

type Action struct {
	ID      uuid.UUID         `json:"id,omitzero"`
	Svc     enum.ModelService `json:"svc,omitzero"`
	ModelID uuid.UUID         `json:"modelID,omitzero"`
	UserID  uuid.UUID         `json:"userID,omitzero"`
	Act     string            `json:"act,omitzero"`
	Content util.ValueMap     `json:"content,omitempty"`
	Note    string            `json:"note,omitzero"`
	Created time.Time         `json:"created,omitzero"`
}

func NewAction(id uuid.UUID) *Action {
	return &Action{ID: id}
}

func (a *Action) Clone() *Action {
	if a == nil {
		return nil
	}
	return &Action{
		ID: a.ID, Svc: a.Svc, ModelID: a.ModelID, UserID: a.UserID, Act: a.Act, Content: a.Content.Clone(), Note: a.Note,
		Created: a.Created,
	}
}

func (a *Action) String() string {
	return a.ID.String()
}

func (a *Action) TitleString() string {
	return a.String()
}

func RandomAction() *Action {
	return &Action{
		ID:      util.UUID(),
		Svc:     enum.AllModelServices.Random(),
		ModelID: util.UUID(),
		UserID:  util.UUID(),
		Act:     util.RandomString(12),
		Content: util.RandomValueMap(4),
		Note:    util.RandomString(12),
		Created: util.TimeCurrent(),
	}
}

//nolint:lll
func (a *Action) Strings() []string {
	return []string{a.ID.String(), a.Svc.String(), a.ModelID.String(), a.UserID.String(), a.Act, util.ToJSONCompact(a.Content), a.Note, util.TimeToFull(&a.Created)}
}

func (a *Action) ToCSV() ([]string, [][]string) {
	return ActionFieldDescs.Keys(), [][]string{a.Strings()}
}

func (a *Action) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(append(paths, url.QueryEscape(a.ID.String()))...)
}

func (a *Action) Breadcrumb(extra ...string) string {
	return a.TitleString() + "||" + a.WebPath(extra...) + "**action"
}

func (a *Action) ToData() []any {
	return []any{a.ID, a.Svc, a.ModelID, a.UserID, a.Act, a.Content, a.Note, a.Created}
}

var ActionFieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Type: "uuid"},
	{Key: "svc", Title: "Svc", Type: "enum(model_service)"},
	{Key: "modelID", Title: "Model ID", Type: "uuid"},
	{Key: "userID", Title: "User ID", Type: "uuid"},
	{Key: "act", Title: "Act", Type: "string"},
	{Key: "content", Title: "Content", Type: "map"},
	{Key: "note", Title: "Note", Type: "string"},
	{Key: "created", Title: "Created", Type: "timestamp"},
}
