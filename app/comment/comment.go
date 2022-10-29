// Content managed by Project Forge, see [projectforge.md] for details.
package comment

import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

type Comment struct {
	ID         uuid.UUID         `json:"id"`
	Svc        enum.ModelService `json:"svc"`
	ModelID    uuid.UUID         `json:"modelID"`
	TargetType string            `json:"targetType"`
	TargetID   uuid.UUID         `json:"targetID"`
	UserID     uuid.UUID         `json:"userID"`
	Content    string            `json:"content"`
	HTML       string            `json:"html"`
	Created    time.Time         `json:"created"`
}

func New(id uuid.UUID) *Comment {
	return &Comment{ID: id}
}

func Random() *Comment {
	return &Comment{
		ID:         util.UUID(),
		Svc:        enum.ModelService(util.RandomString(12)),
		ModelID:    util.UUID(),
		TargetType: util.RandomString(12),
		TargetID:   util.UUID(),
		UserID:     util.UUID(),
		Content:    util.RandomString(12),
		HTML:       util.RandomString(12),
		Created:    time.Now(),
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
	ret.Svc = enum.ModelService(retSvc)
	retModelID, e := m.ParseUUID("modelID", true, true)
	if e != nil {
		return nil, e
	}
	if retModelID != nil {
		ret.ModelID = *retModelID
	}
	ret.TargetType, err = m.ParseString("targetType", true, true)
	if err != nil {
		return nil, err
	}
	retTargetID, e := m.ParseUUID("targetID", true, true)
	if e != nil {
		return nil, e
	}
	if retTargetID != nil {
		ret.TargetID = *retTargetID
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
	return &Comment{
		ID:         c.ID,
		Svc:        c.Svc,
		ModelID:    c.ModelID,
		TargetType: c.TargetType,
		TargetID:   c.TargetID,
		UserID:     c.UserID,
		Content:    c.Content,
		HTML:       c.HTML,
		Created:    c.Created,
	}
}

func (c *Comment) String() string {
	return c.ID.String()
}

func (c *Comment) TitleString() string {
	return c.String()
}

func (c *Comment) WebPath() string {
	return "/admin/db/comment" + "/" + c.ID.String()
}

func (c *Comment) Diff(cx *Comment) util.Diffs {
	var diffs util.Diffs
	if c.ID != cx.ID {
		diffs = append(diffs, util.NewDiff("id", c.ID.String(), cx.ID.String()))
	}
	if c.Svc != cx.Svc {
		diffs = append(diffs, util.NewDiff("svc", string(c.Svc), string(cx.Svc)))
	}
	if c.ModelID != cx.ModelID {
		diffs = append(diffs, util.NewDiff("modelID", c.ModelID.String(), cx.ModelID.String()))
	}
	if c.TargetType != cx.TargetType {
		diffs = append(diffs, util.NewDiff("targetType", c.TargetType, cx.TargetType))
	}
	if c.TargetID != cx.TargetID {
		diffs = append(diffs, util.NewDiff("targetID", c.TargetID.String(), cx.TargetID.String()))
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
	return []any{c.ID, c.Svc, c.ModelID, c.TargetType, c.TargetID, c.UserID, c.Content, c.HTML, c.Created}
}
