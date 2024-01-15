// Package action - Content managed by Project Forge, see [projectforge.md] for details.
package action

import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

type Action struct {
	ID      uuid.UUID         `json:"id,omitempty"`
	Svc     enum.ModelService `json:"svc,omitempty"`
	ModelID uuid.UUID         `json:"modelID,omitempty"`
	UserID  uuid.UUID         `json:"userID,omitempty"`
	Act     string            `json:"act,omitempty"`
	Content util.ValueMap     `json:"content,omitempty"`
	Note    string            `json:"note,omitempty"`
	Created time.Time         `json:"created,omitempty"`
}

func New(id uuid.UUID) *Action {
	return &Action{ID: id}
}

func (a *Action) Clone() *Action {
	return &Action{a.ID, a.Svc, a.ModelID, a.UserID, a.Act, a.Content.Clone(), a.Note, a.Created}
}

func (a *Action) String() string {
	return a.ID.String()
}

func (a *Action) TitleString() string {
	return a.String()
}

func Random() *Action {
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

func (a *Action) WebPath() string {
	return "/admin/db/action/" + a.ID.String()
}

func (a *Action) ToData() []any {
	return []any{a.ID, a.Svc, a.ModelID, a.UserID, a.Act, a.Content, a.Note, a.Created}
}

var FieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Description: "", Type: "uuid"},
	{Key: "svc", Title: "Svc", Description: "", Type: "enum(model_service)"},
	{Key: "modelID", Title: "Model ID", Description: "", Type: "uuid"},
	{Key: "userID", Title: "User ID", Description: "", Type: "uuid"},
	{Key: "act", Title: "Act", Description: "", Type: "string"},
	{Key: "content", Title: "Content", Description: "", Type: "map[string]any"},
	{Key: "note", Title: "Note", Description: "", Type: "string"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
}
