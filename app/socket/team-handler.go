package socket

import (
	"fmt"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/model/team"
	"github.com/kyleu/rituals.dev/app/util"
)

func onTeamSessionSave(s *Service, ch Channel, userID uuid.UUID, param teamSessionSaveParams) error {
	dataSvc := s.teams
	title := util.ServiceTitle(util.SvcTeam, param.Title)

	curr := dataSvc.GetByID(ch.ID)
	if curr == nil {
		return errors.New("no team available with id [" + ch.ID.String() + "]")
	}

	sr := s.checkPerms(userID, nil, nil, ch.Svc, ch.ID)
	if sr != nil {
		return sr
	}

	msg := "saving team session [%s]"
	s.Logger.Debug(fmt.Sprintf(msg, title))

	err := dataSvc.UpdateSession(ch.ID, title, userID)
	if err != nil {
		return errors.Wrap(err, "error updating team session")
	}

	if title != curr.Title {
		slug, err := dataSvc.Data.History.UpdateSlug(curr.ID, curr.Slug, curr.Title, title, userID)
		if err != nil {
			return errors.Wrap(err, "error updating team slug from ["+curr.Slug+"] to ["+slug+"]")
		}
	}

	err = sendTeamSessionUpdate(s, ch)
	if err != nil {
		return errors.Wrap(err, "error sending team session")
	}

	err = s.updatePerms(ch, userID, nil, nil, dataSvc.Data.Permissions, param.Permissions)
	if err != nil {
		return errors.Wrap(err, "error updating permissions")
	}

	return nil
}

func sendTeamUpdate(s *Service, ch Channel, curr *uuid.UUID, tm *team.Session) error {
	err := s.WriteChannel(ch, NewMessage(ch.Svc, ServerCmdTeamUpdate, tm))
	if err != nil {
		return errors.Wrap(err, "error writing team update message")
	}

	err = s.SendContentUpdate(util.SvcTeam, curr)
	if err != nil {
		return err
	}
	if tm != nil {
		err = s.SendContentUpdate(util.SvcTeam, &tm.ID)
	}
	return err
}

func sendTeamSessionUpdate(s *Service, ch Channel) error {
	sess := s.teams.GetByID(ch.ID)
	if sess == nil {
		return errors.New("cannot load team session [" + ch.ID.String() + "]")
	}
	err := s.WriteChannel(ch, NewMessage(util.SvcTeam, ServerCmdSessionUpdate, sess))
	return errors.Wrap(err, "error sending team session")
}

func getTeamOpt(s *Service, teamID *uuid.UUID) *team.Session {
	if teamID == nil {
		return nil
	}
	return s.teams.GetByID(*teamID)
}
