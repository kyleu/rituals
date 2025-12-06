package feedback

import (
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/lib/svc"
	"github.com/kyleu/rituals/app/util"
)

const DefaultRoute = "/admin/db/retro/feedback"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(paths...)
}

var _ svc.Model = (*Feedback)(nil)

type Feedback struct {
	ID       uuid.UUID  `json:"id,omitzero"`
	RetroID  uuid.UUID  `json:"retroID,omitzero"`
	Idx      int        `json:"idx,omitzero"`
	UserID   uuid.UUID  `json:"userID,omitzero"`
	Category string     `json:"category,omitzero"`
	Content  string     `json:"content,omitzero"`
	HTML     string     `json:"html,omitzero"`
	Created  time.Time  `json:"created,omitzero"`
	Updated  *time.Time `json:"updated,omitzero"`
}

func NewFeedback(id uuid.UUID) *Feedback {
	return &Feedback{ID: id}
}

func (f *Feedback) Clone() *Feedback {
	if f == nil {
		return nil
	}
	return &Feedback{
		ID: f.ID, RetroID: f.RetroID, Idx: f.Idx, UserID: f.UserID, Category: f.Category, Content: f.Content, HTML: f.HTML,
		Created: f.Created, Updated: f.Updated,
	}
}

func (f *Feedback) String() string {
	return f.ID.String()
}

func (f *Feedback) TitleString() string {
	return f.String()
}

func RandomFeedback() *Feedback {
	return &Feedback{
		ID:       util.UUID(),
		RetroID:  util.UUID(),
		Idx:      util.RandomInt(10000),
		UserID:   util.UUID(),
		Category: util.RandomString(12),
		Content:  util.RandomString(12),
		HTML:     "<h3>" + util.RandomString(6) + "</h3>",
		Created:  util.TimeCurrent(),
		Updated:  util.TimeCurrentP(),
	}
}

//nolint:lll
func (f *Feedback) Strings() []string {
	return []string{f.ID.String(), f.RetroID.String(), fmt.Sprint(f.Idx), f.UserID.String(), f.Category, f.Content, f.HTML, util.TimeToFull(&f.Created), util.TimeToFull(f.Updated)}
}

func (f *Feedback) ToCSV() ([]string, [][]string) {
	return FeedbackFieldDescs.Keys(), [][]string{f.Strings()}
}

func (f *Feedback) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(append(paths, url.QueryEscape(f.ID.String()))...)
}

func (f *Feedback) Breadcrumb(extra ...string) string {
	return f.TitleString() + "||" + f.WebPath(extra...) + "**comment"
}

func (f *Feedback) ToData() []any {
	return []any{f.ID, f.RetroID, f.Idx, f.UserID, f.Category, f.Content, f.HTML, f.Created, f.Updated}
}

var FeedbackFieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Type: "uuid"},
	{Key: "retroID", Title: "Retro ID", Type: "uuid"},
	{Key: "idx", Title: "Idx", Type: "int"},
	{Key: "userID", Title: "User ID", Type: "uuid"},
	{Key: "category", Title: "Category", Type: "string"},
	{Key: "content", Title: "Content", Type: "string"},
	{Key: "html", Title: "HTML", Type: "string"},
	{Key: "created", Title: "Created", Type: "timestamp"},
	{Key: "updated", Title: "Updated", Type: "timestamp"},
}
