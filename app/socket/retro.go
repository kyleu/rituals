package socket

import (
	"emperror.dev/errors"
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/query"
	"github.com/kyleu/rituals.dev/app/retro"
	"github.com/kyleu/rituals.dev/app/util"
)

func onRetroMessage(s *Service, conn *connection, cmd string, param json.RawMessage) error {
	var err error
	userID := conn.Profile.UserID

	switch cmd {
	case ClientCmdConnect:
		var u uuid.UUID
		util.FromJSON(param, &u, s.logger)
		err = onRetroConnect(s, conn, u)
	case ClientCmdUpdateSession:
		rss := retroSessionSaveParams{}
		util.FromJSON(param, &rss, s.logger)
		err = onRetroSessionSave(s, *conn.Channel, userID, rss)
	case ClientCmdRemoveMember:
		var u uuid.UUID
		util.FromJSON(param, &u, s.logger)
		err = onRemoveMember(s, s.retros.Members, *conn.Channel, userID, u)
	case ClientCmdAddFeedback:
		afp := addFeedbackParams{}
		util.FromJSON(param, &afp, s.logger)
		err = onAddFeedback(s, *conn.Channel, userID, afp)
	case ClientCmdUpdateFeedback:
		efp := editFeedbackParams{}
		util.FromJSON(param, &efp, s.logger)
		err = onEditFeedback(s, *conn.Channel, userID, efp)
	case ClientCmdRemoveFeedback:
		var u uuid.UUID
		util.FromJSON(param, &u, s.logger)
		err = onRemoveFeedback(s, *conn.Channel, userID, u)
	default:
		err = errors.New("unhandled retro command [" + cmd + "]")
	}
	return errors.Wrap(err, "error handling retro message")
}

func onRetroSessionSave(s *Service, ch channel, userID uuid.UUID, param retroSessionSaveParams) error {
	title := util.ServiceTitle(util.SvcRetro, param.Title)
	categories := query.StringToArray(param.Categories)
	if len(categories) == 0 {
		categories = retro.DefaultCategories
	}

	sprintID := util.GetUUIDFromString(param.SprintID)
	teamID := util.GetUUIDFromString(param.TeamID)

	msg := "saving retro session [%s] with categories [%s], sprint [%s] and team [%s]"
	s.logger.Debug(fmt.Sprintf(msg, title, util.OxfordComma(categories, "and"), sprintID, teamID))

	curr, err := s.retros.GetByID(ch.ID)
	if err != nil {
		return errors.Wrap(err, "error loading retro session ["+ch.ID.String()+"] for update")
	}

	teamChanged := differentPointerValues(curr.TeamID, teamID)
	sprintChanged := differentPointerValues(curr.SprintID, sprintID)

	err = s.retros.UpdateSession(ch.ID, title, categories, teamID, sprintID, userID)
	if err != nil {
		return errors.Wrap(err, "error updating retro session")
	}

	err = sendRetroSessionUpdate(s, ch)
	if err != nil {
		return errors.Wrap(err, "error sending retro session")
	}

	if teamChanged {
		tm := s.teams.GetByIDPointer(teamID)
		err = sendTeamUpdate(s, ch, curr.TeamID, tm)
		if err != nil {
			return errors.Wrap(err, "error sending team for updated retro session")
		}
	}

	if sprintChanged {
		spr := s.sprints.GetByIDPointer(sprintID)
		err = sendSprintUpdate(s, ch, curr.SprintID, spr)
		if err != nil {
			return errors.Wrap(err, "error sending sprint for updated retro session")
		}
	}

	err = s.updatePerms(ch, userID, s.retros.Permissions, param.Permissions)
	if err != nil {
		return errors.Wrap(err, "error updating permissions")
	}

	return nil
}

func sendRetroSessionUpdate(s *Service, ch channel) error {
	sess, err := s.retros.GetByID(ch.ID)
	if err != nil {
		return errors.Wrap(err, "error finding retro session ["+ch.ID.String()+"]")
	}
	if sess == nil {
		return errors.Wrap(err, "cannot load retro session ["+ch.ID.String()+"]")
	}

	err = s.WriteChannel(ch, NewMessage(util.SvcRetro, ServerCmdSessionUpdate, sess))
	return errors.Wrap(err, "error sending retro session")
}
