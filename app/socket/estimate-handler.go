package socket

import (
	"emperror.dev/errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/database/query"
	"github.com/kyleu/rituals.dev/app/model/estimate"
	"github.com/kyleu/rituals.dev/app/util"
)

func onEstimateSessionSave(s *Service, ch Channel, userID uuid.UUID, param estimateSessionSaveParams) error {
	title := util.ServiceTitle(util.SvcEstimate, param.Title)

	choices := query.StringToArray(param.Choices)
	if len(choices) == 0 {
		choices = estimate.DefaultChoices
	}

	sprintID := util.GetUUIDFromString(param.SprintID)
	teamID := util.GetUUIDFromString(param.TeamID)

	msg := "saving estimate session [%s] with choices [%s], team [%s] and sprint [%s]"
	s.Logger.Debug(fmt.Sprintf(msg, title, util.OxfordComma(choices, "and"), teamID, sprintID))

	curr := s.estimates.GetByID(ch.ID)

	teamChanged := differentPointerValues(curr.TeamID, teamID)
	sprintChanged := differentPointerValues(curr.SprintID, sprintID)

	err := s.estimates.UpdateSession(ch.ID, title, choices, teamID, sprintID, userID)
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

	err = s.updatePerms(ch, userID, teamID, sprintID, s.estimates.Data.Permissions, param.Permissions)
	if err != nil {
		return errors.Wrap(err, "error updating permissions")
	}

	return nil
}

func sendEstimateSessionUpdate(s *Service, ch Channel) error {
	est := s.estimates.GetByID(ch.ID)
	if est == nil {
		return errors.New("cannot load estimate session ["+ch.ID.String()+"]")
	}

	err := s.WriteChannel(ch, NewMessage(util.SvcEstimate, ServerCmdSessionUpdate, est))
	return errors.Wrap(err, "error sending estimate session")
}
