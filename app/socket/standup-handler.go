package socket

import (
	"fmt"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/util"
)

func onStandupSessionSave(s *Service, ch Channel, userID uuid.UUID, param standupSessionSaveParams) error {
	dataSvc := s.standups
	title := util.ServiceTitle(util.SvcStandup, param.Title)

	sprintID := util.GetUUIDFromString(param.SprintID)
	teamID := util.GetUUIDFromString(param.TeamID)

	curr := dataSvc.GetByID(ch.ID)
	if curr == nil {
		return errors.New("no standup available with id [" + ch.ID.String() + "]")
	}

	sr := s.checkPerms(userID, curr.TeamID, curr.SprintID, ch.Svc, ch.ID)
	if sr != nil {
		return sr
	}

	teamChanged := differentPointerValues(curr.TeamID, teamID)
	sprintChanged := differentPointerValues(curr.SprintID, sprintID)

	msg := "saving standup session [%s] with sprint [%s] and team [%s]"
	s.Logger.Debug(fmt.Sprintf(msg, title, sprintID, teamID))

	err := dataSvc.UpdateSession(ch.ID, title, teamID, sprintID, userID)
	if err != nil {
		return errors.Wrap(err, "error updating standup session")
	}

	if title != curr.Title {
		slug, err := dataSvc.Data.History.UpdateSlug(curr.ID, curr.Slug, curr.Title, title, userID)
		if err != nil {
			return errors.Wrap(err, "error updating standup slug from ["+curr.Slug+"] to ["+slug+"]")
		}
	}

	err = sendStandupSessionUpdate(s, ch)
	if err != nil {
		return errors.Wrap(err, "error sending standup session")
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
			return errors.Wrap(err, "error sending sprint for updated standup session")
		}
	}

	err = s.updatePerms(ch, userID, teamID, sprintID, dataSvc.Data.Permissions, param.Permissions)
	if err != nil {
		return errors.Wrap(err, "error updating permissions")
	}

	return nil
}

func sendStandupSessionUpdate(s *Service, ch Channel) error {
	sess := s.standups.GetByID(ch.ID)
	if sess == nil {
		return errors.New("cannot load standup session [" + ch.ID.String() + "]")
	}
	err := s.WriteChannel(ch, NewMessage(util.SvcStandup, ServerCmdSessionUpdate, sess))
	return errors.Wrap(err, "error sending standup session")
}
