package comment

import (
	"net/url"
	"path"
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
	return path.Join(paths...)
}

var _ svc.Model = (*Comment)(nil)

type Comment struct {
	ID      uuid.UUID         `json:"id,omitempty"`
	Svc     enum.ModelService `json:"svc,omitempty"`
	ModelID uuid.UUID         `json:"modelID,omitempty"`
	UserID  uuid.UUID         `json:"userID,omitempty"`
	Content string            `json:"content,omitempty"`
	HTML    string            `json:"html,omitempty"`
	Created time.Time         `json:"created,omitempty"`
}

func NewComment(id uuid.UUID) *Comment {
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
	return path.Join(append(paths, url.QueryEscape(c.ID.String()))...)
}

func (c *Comment) Breadcrumb(extra ...string) string {
	return c.TitleString() + "||" + c.WebPath(extra...) + "**comments"
}

func (c *Comment) ToData() []any {
	return []any{c.ID, c.Svc, c.ModelID, c.UserID, c.Content, c.HTML, c.Created}
}

var CommentFieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Description: "", Type: "uuid"},
	{Key: "svc", Title: "Svc", Description: "", Type: "enum(model_service)"},
	{Key: "modelID", Title: "Model ID", Description: "", Type: "uuid"},
	{Key: "userID", Title: "User ID", Description: "", Type: "uuid"},
	{Key: "content", Title: "Content", Description: "", Type: "string"},
	{Key: "html", Title: "HTML", Description: "", Type: "string"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
}
