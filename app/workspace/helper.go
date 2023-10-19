package workspace

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/comment"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

func commentFromForm(frm util.ValueMap, userID uuid.UUID) (*comment.Comment, string, error) {
	svcStr := frm.GetStringOpt("svc")
	if svcStr == "" {
		return nil, "", errors.New("must provide [svc]")
	}
	svc := enum.AllModelServices.Get(svcStr, nil)
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

func sendTeamSprintUpdates(t enum.ModelService, teamID *uuid.UUID, sprintID *uuid.UUID, model any, userID *uuid.UUID, svc *Service, logger util.Logger) error {
	msg := map[string]any{"type": t, "model": model}
	if teamID != nil {
		err := svc.send(enum.ModelServiceTeam, *teamID, action.ActChildUpdate, msg, userID, logger)
		if err != nil {
			return err
		}
	}
	if sprintID != nil {
		msg := map[string]any{"type": t, "model": model}
		err := svc.send(enum.ModelServiceSprint, *sprintID, action.ActChildUpdate, msg, userID, logger)
		if err != nil {
			return err
		}
	}
	return nil
}
