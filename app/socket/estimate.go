package socket

import (
	"fmt"
	"strings"

	"github.com/kyleu/rituals.dev/app/query"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/estimate"
	"github.com/kyleu/rituals.dev/app/util"
)

func onEstimateMessage(s *Service, conn *connection, userID uuid.UUID, cmd string, param interface{}) error {
	var err error

	switch cmd {
	case ClientCmdConnect:
		err = onEstimateConnect(s, conn, userID, param.(string))
	case ClientCmdUpdateSession:
		err = onEstimateSessionSave(s, *conn.Channel, userID, param.(map[string]interface{}))
	case ClientCmdRemoveMember:
		err = onRemoveMember(s, s.estimates.Members, *conn.Channel, userID, param.(string))
	case ClientCmdAddStory:
		err = onAddStory(s, *conn.Channel, userID, param.(map[string]interface{}))
	case ClientCmdUpdateStory:
		err = onUpdateStory(s, *conn.Channel, userID, param.(map[string]interface{}))
	case ClientCmdRemoveStory:
		err = onRemoveStory(s, *conn.Channel, userID, param.(string))
	case ClientCmdSetStoryStatus:
		err = onSetStoryStatus(s, *conn.Channel, userID, param.(map[string]interface{}))
	case ClientCmdSubmitVote:
		err = onSubmitVote(s, *conn.Channel, userID, param.(map[string]interface{}))
	default:
		err = errors.New("unhandled estimate command [" + cmd + "]")
	}
	return errors.Wrap(err, "error handling estimate message")
}

func onEstimateSessionSave(s *Service, ch channel, userID uuid.UUID, param map[string]interface{}) error {
	titleString, ok := param[util.KeyTitle].(string)
	if !ok {
		return errors.New("cannot read choices as string")
	}
	title := util.ServiceTitle(util.SvcEstimate, titleString)

	choicesString, ok := param["choices"].(string)
	if !ok {
		return errors.New(fmt.Sprintf("cannot parse [%v] as string", param["choices"]))
	}
	choices := query.StringToArray(choicesString)
	if len(choices) == 0 {
		choices = estimate.DefaultChoices
	}

	sprintID := getUUIDPointer(param, util.WithID(util.SvcSprint.Key))
	teamID := getUUIDPointer(param, util.WithID(util.SvcTeam.Key))

	msg := "saving estimate session [%s] with choices [%s], team [%s] and sprint [%s]"
	s.logger.Debug(fmt.Sprintf(msg, title, strings.Join(choices, ", "), teamID, sprintID))

	curr, err := s.estimates.GetByID(ch.ID)
	if err != nil {
		return errors.Wrap(err, "error loading estimate session ["+ch.ID.String()+"] for update")
	}

	teamChanged := differentPointerValues(curr.TeamID, teamID)
	sprintChanged := differentPointerValues(curr.SprintID, sprintID)

	err = s.estimates.UpdateSession(ch.ID, title, choices, teamID, sprintID, userID)
	if err != nil {
		return errors.Wrap(err, "error updating estimate session")
	}

	err = sendEstimateSessionUpdate(s, ch)
	if err != nil {
		return errors.Wrap(err, "error sending estimate session")
	}

	if teamChanged {
		tm := s.teams.GetByIDPointer(teamID)
		err = sendTeamUpdate(s, ch, curr.TeamID, tm)
		if err != nil {
			return errors.Wrap(err, "error sending team for updated estimate session")
		}
	}

	if sprintChanged {
		spr := s.sprints.GetByIDPointer(sprintID)
		err = sendSprintUpdate(s, ch, curr.SprintID, spr)
		if err != nil {
			return errors.Wrap(err, "error sending sprint for updated estimate session")
		}
	}

	err = s.updatePerms(ch, userID, s.estimates.Permissions, param)
	if err != nil {
		return errors.Wrap(err, "error updating permissions")
	}

	return nil
}

func sendEstimateSessionUpdate(s *Service, ch channel) error {
	est, err := s.estimates.GetByID(ch.ID)
	if err != nil {
		return errors.Wrap(err, "error finding estimate session ["+ch.ID.String()+"]")
	}
	if est == nil {
		return errors.Wrap(err, "cannot load estimate session ["+ch.ID.String()+"]")
	}

	msg := Message{Svc: util.SvcEstimate.Key, Cmd: ServerCmdSessionUpdate, Param: est}
	err = s.WriteChannel(ch, &msg)
	return errors.Wrap(err, "error sending estimate session")
}
