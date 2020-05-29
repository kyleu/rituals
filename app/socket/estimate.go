package socket

import (
	"encoding/json"
	"fmt"
	"github.com/kyleu/rituals.dev/app/query"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/estimate"
	"github.com/kyleu/rituals.dev/app/util"
)

func onEstimateMessage(s *Service, conn *connection, cmd string, param json.RawMessage) error {
	var err error
	userID := conn.Profile.UserID
	switch cmd {
	case ClientCmdConnect:
		var u uuid.UUID
		util.FromJSON(param, &u, s.logger)
		err = onEstimateConnect(s, conn, u)
	case ClientCmdUpdateSession:
		ess := estimateSessionSaveParams{}
		util.FromJSON(param, &ess, s.logger)
		err = onEstimateSessionSave(s, *conn.Channel, userID, ess)
	case ClientCmdRemoveMember:
		var u uuid.UUID
		util.FromJSON(param, &u, s.logger)
		err = onRemoveMember(s, s.estimates.Members, *conn.Channel, userID, u)
	case ClientCmdAddStory:
		asp := addStoryParams{}
		util.FromJSON(param, &asp, s.logger)
		err = onAddStory(s, *conn.Channel, userID, asp)
	case ClientCmdUpdateStory:
		usp := updateStoryParams{}
		util.FromJSON(param, &usp, s.logger)
		err = onUpdateStory(s, *conn.Channel, userID, usp)
	case ClientCmdRemoveStory:
		var u uuid.UUID
		util.FromJSON(param, &u, s.logger)
		err = onRemoveStory(s, *conn.Channel, userID, u)
	case ClientCmdSetStoryStatus:
		sssp := setStoryStatusParams{}
		util.FromJSON(param, &sssp, s.logger)
		err = onSetStoryStatus(s, *conn.Channel, userID, sssp)
	case ClientCmdSubmitVote:
		svp := submitVoteParams{}
		util.FromJSON(param, &svp, s.logger)
		err = onSubmitVote(s, *conn.Channel, userID, svp)
	default:
		err = errors.New("unhandled estimate command [" + cmd + "]")
	}
	return errors.Wrap(err, "error handling estimate message")
}

func onEstimateSessionSave(s *Service, ch channel, userID uuid.UUID, param estimateSessionSaveParams) error {
	title := util.ServiceTitle(util.SvcEstimate, param.Title)

	choices := query.StringToArray(param.Choices)
	if len(choices) == 0 {
		choices = estimate.DefaultChoices
	}

	sprintID := util.GetUUIDFromString(param.SprintID)
	teamID := util.GetUUIDFromString(param.TeamID)

	msg := "saving estimate session [%s] with choices [%s], team [%s] and sprint [%s]"
	s.logger.Debug(fmt.Sprintf(msg, title, util.OxfordComma(choices, "and"), teamID, sprintID))

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

	err = s.updatePerms(ch, userID, s.estimates.Permissions, param.Permissions)
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

	err = s.WriteChannel(ch, NewMessage(util.SvcEstimate, ServerCmdSessionUpdate, est))
	return errors.Wrap(err, "error sending estimate session")
}
