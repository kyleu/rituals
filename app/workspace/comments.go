package workspace

import (
	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/comment"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

func sendComment(
	svc enum.ModelService, modelID uuid.UUID, c *comment.Comment, self *uuid.UUID, team *uuid.UUID, sprint *uuid.UUID,
	send action.SendFn, logger util.Logger, connIDs ...uuid.UUID,
) error {
	err := send(svc, modelID, action.ActComment, c, self, logger, connIDs...)
	if err != nil {
		return err
	}
	if team != nil {
		err = send(enum.ModelServiceTeam, *team, action.ActComment, c, self, logger, connIDs...)
		if err != nil {
			return err
		}
	}
	if sprint != nil {
		err = send(enum.ModelServiceSprint, *sprint, action.ActComment, c, self, logger, connIDs...)
		if err != nil {
			return err
		}
	}
	return nil
}
