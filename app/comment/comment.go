// Package comment - Content managed by Project Forge, see [projectforge.md] for details.
package comment

import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

type Comment struct {
	ID      uuid.UUID         `json:"id,omitempty"`
	Svc     enum.ModelService `json:"svc,omitempty"`
	ModelID uuid.UUID         `json:"modelID,omitempty"`
	UserID  uuid.UUID         `json:"userID,omitempty"`
	Content string            `json:"content,omitempty"`
	HTML    string            `json:"html,omitempty"`
	Created time.Time         `json:"created,omitempty"`
}

func New(id uuid.UUID) *Comment {
	return &Comment{ID: id}
}

func (c *Comment) Clone() *Comment {
	return &Comment{c.ID, c.Svc, c.ModelID, c.UserID, c.Content, c.HTML, c.Created}
}

func (c *Comment) String() string {
	return c.ID.String()
}

func (c *Comment) TitleString() string {
	return c.String()
}

func Random() *Comment {
	return &Comment{
		ID:      util.UUID(),
		Svc:     enum.AllModelServices.Random(),
		ModelID: util.UUID(),
		UserID:  util.UUID(),
		Content: util.RandomString(12),
		HTML:    "<h3>" + util.RandomString(6) + "</h3>",
		Created: util.TimeCurrent(),
	}
}

func (c *Comment) WebPath() string {
	return "/admin/db/comment/" + c.ID.String()
}

func (c *Comment) ToData() []any {
	return []any{c.ID, c.Svc, c.ModelID, c.UserID, c.Content, c.HTML, c.Created}
}
