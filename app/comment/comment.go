// Package comment - Content managed by Project Forge, see [projectforge.md] for details.
package comment

import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

type Comment struct {
	ID      uuid.UUID         `json:"id"`
	Svc     enum.ModelService `json:"svc"`
	ModelID uuid.UUID         `json:"modelID"`
	UserID  uuid.UUID         `json:"userID"`
	Content string            `json:"content"`
	HTML    string            `json:"html"`
	Created time.Time         `json:"created"`
}

func New(id uuid.UUID) *Comment {
	return &Comment{ID: id}
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

func FromMap(m util.ValueMap, setPK bool) (*Comment, error) {
	ret := &Comment{}
	var err error
	if setPK {
		retID, e := m.ParseUUID("id", true, true)
		if e != nil {
			return nil, e
		}
		if retID != nil {
			ret.ID = *retID
		}
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
	}
	retSvc, err := m.ParseString("svc", true, true)
	if err != nil {
		return nil, err
	}
	ret.Svc = enum.AllModelServices.Get(retSvc, nil)
	retModelID, e := m.ParseUUID("modelID", true, true)
	if e != nil {
		return nil, e
	}
	if retModelID != nil {
		ret.ModelID = *retModelID
	}
	retUserID, e := m.ParseUUID("userID", true, true)
	if e != nil {
		return nil, e
	}
	if retUserID != nil {
		ret.UserID = *retUserID
	}
	ret.Content, err = m.ParseString("content", true, true)
	if err != nil {
		return nil, err
	}
	ret.HTML, err = m.ParseString("html", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
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

func (c *Comment) WebPath() string {
	return "/admin/db/comment/" + c.ID.String()
}

func (c *Comment) Diff(cx *Comment) util.Diffs {
	var diffs util.Diffs
	if c.ID != cx.ID {
		diffs = append(diffs, util.NewDiff("id", c.ID.String(), cx.ID.String()))
	}
	if c.Svc != cx.Svc {
		diffs = append(diffs, util.NewDiff("svc", c.Svc.Key, cx.Svc.Key))
	}
	if c.ModelID != cx.ModelID {
		diffs = append(diffs, util.NewDiff("modelID", c.ModelID.String(), cx.ModelID.String()))
	}
	if c.UserID != cx.UserID {
		diffs = append(diffs, util.NewDiff("userID", c.UserID.String(), cx.UserID.String()))
	}
	if c.Content != cx.Content {
		diffs = append(diffs, util.NewDiff("content", c.Content, cx.Content))
	}
	if c.HTML != cx.HTML {
		diffs = append(diffs, util.NewDiff("html", c.HTML, cx.HTML))
	}
	if c.Created != cx.Created {
		diffs = append(diffs, util.NewDiff("created", c.Created.String(), cx.Created.String()))
	}
	return diffs
}

func (c *Comment) ToData() []any {
	return []any{c.ID, c.Svc, c.ModelID, c.UserID, c.Content, c.HTML, c.Created}
}
