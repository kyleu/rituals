package socket

import (
	"fmt"

	"github.com/kyleu/npn/npnconnection"

	"github.com/kyleu/npn/npncore"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/util"
)

func onStandupSessionSave(s *npnconnection.Service, ch npnconnection.Channel, userID uuid.UUID, param standupSessionSaveParams) error {
	dataSvc := standups(s)
	title := util.ServiceTitle(util.SvcStandup, param.Title)

	sprintID := npncore.GetUUIDFromString(param.SprintID)
	teamID := npncore.GetUUIDFromString(param.TeamID)

	curr := dataSvc.GetByID(ch.ID)
	if curr == nil {
		return errors.New("no standup available with id [" + ch.ID.String() + "]")
	}

	sr := checkPerms(s, userID, curr.TeamID, curr.SprintID, ch.Svc, ch.ID)
	if sr != nil {
		return sr
	}

	teamChanged := npnconnection.DifferentPointerValues(curr.TeamID, teamID)
	sprintChanged := npnconnection.DifferentPointerValues(curr.SprintID, sprintID)

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
		tm := teams(s).GetByIDPointer(teamID)
		err = sendTeamUpdate(s, ch, curr.TeamID, tm)
		if err != nil {
			return errors.Wrap(err, "error sending team for updated retro session")
		}
	}

	if sprintChanged {
		spr := sprints(s).GetByIDPointer(sprintID)
		err = sendSprintUpdate(s, ch, curr.SprintID, spr)
		if err != nil {
			return errors.Wrap(err, "error sending sprint for updated standup session")
		}
	}

	err = updatePerms(s, ch, userID, teamID, sprintID, dataSvc.Data.Permissions, param.Permissions)
	if err != nil {
		return errors.Wrap(err, "error updating permissions")
	}

	return nil
}

func sendStandupSessionUpdate(s *npnconnection.Service, ch npnconnection.Channel) error {
	sess := standups(s).GetByID(ch.ID)
	if sess == nil {
		return errors.New("cannot load standup session [" + ch.ID.String() + "]")
	}
	err := s.WriteChannel(ch, npnconnection.NewMessage(util.SvcStandup.Key, ServerCmdSessionUpdate, sess))
	return errors.Wrap(err, "error sending standup session")
}
