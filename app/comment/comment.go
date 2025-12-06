package comment

import (
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

const DefaultRoute = "/admin/db/comment"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(paths...)
}

var _ svc.Model = (*Comment)(nil)

type Comment struct {
	ID      uuid.UUID         `json:"id,omitzero"`
	Svc     enum.ModelService `json:"svc,omitzero"`
	ModelID uuid.UUID         `json:"modelID,omitzero"`
	UserID  uuid.UUID         `json:"userID,omitzero"`
	Content string            `json:"content,omitzero"`
	HTML    string            `json:"html,omitzero"`
	Created time.Time         `json:"created,omitzero"`
}

func NewComment(id uuid.UUID) *Comment {
	return &Comment{ID: id}
}

func (c *Comment) Clone() *Comment {
	if c == nil {
		return nil
	}
	return &Comment{
		ID: c.ID, Svc: c.Svc, ModelID: c.ModelID, UserID: c.UserID, Content: c.Content, HTML: c.HTML, Created: c.Created,
	}
}

func (c *Comment) String() string {
	return c.ID.String()
}

func (c *Comment) TitleString() string {
	return c.String()
}

func RandomComment() *Comment {
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

func (c *Comment) Strings() []string {
	return []string{c.ID.String(), c.Svc.String(), c.ModelID.String(), c.UserID.String(), c.Content, c.HTML, util.TimeToFull(&c.Created)}
}

func (c *Comment) ToCSV() ([]string, [][]string) {
	return CommentFieldDescs.Keys(), [][]string{c.Strings()}
}

func (c *Comment) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(append(paths, url.QueryEscape(c.ID.String()))...)
}

func (c *Comment) Breadcrumb(extra ...string) string {
	return c.TitleString() + "||" + c.WebPath(extra...) + "**comments"
}

func (c *Comment) ToData() []any {
	return []any{c.ID, c.Svc, c.ModelID, c.UserID, c.Content, c.HTML, c.Created}
}

var CommentFieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Type: "uuid"},
	{Key: "svc", Title: "Svc", Type: "enum(model_service)"},
	{Key: "modelID", Title: "Model ID", Type: "uuid"},
	{Key: "userID", Title: "User ID", Type: "uuid"},
	{Key: "content", Title: "Content", Type: "string"},
	{Key: "html", Title: "HTML", Type: "string"},
	{Key: "created", Title: "Created", Type: "timestamp"},
}
