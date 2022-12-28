package workspace

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/comment"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

func commentFromForm(frm util.ValueMap, userID uuid.UUID) (*comment.Comment, string, error) {
	svcStr := frm.GetStringOpt("svc")
	if svcStr == "" {
		return nil, "", errors.New("must provide [svc]")
	}
	svc := enum.ModelService(svcStr)
	modelID, _ := frm.GetUUID("modelID", false)
	if modelID == nil {
		return nil, "", errors.New("must provide [modelID]")
	}
	content := frm.GetStringOpt("content")
	if content == "" {
		return nil, "", errors.New("[content] may not be empty")
	}
	html := util.ToHTML(content, true)
	c := &comment.Comment{ID: util.UUID(), Svc: svc, ModelID: *modelID, UserID: userID, Content: content, HTML: html, Created: time.Now()}
	u := fmt.Sprintf("#modal-%s-%s-comments", c.Svc, c.ModelID.String())

	return c, u, nil
}
