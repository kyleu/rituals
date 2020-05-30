package socket

import (
	"emperror.dev/errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/model/team"
	"github.com/kyleu/rituals.dev/app/util"
)

func onTeamSessionSave(s *Service, ch Channel, userID uuid.UUID, param teamSessionSaveParams) error {
	title := util.ServiceTitle(util.SvcTeam, param.Title)
	s.Logger.Debug(fmt.Sprintf("saving team session [%s]", title))

	err := s.teams.UpdateSession(ch.ID, title, userID)
	if err != nil {
		return errors.Wrap(err, "error updating team session")
	}

	err = sendTeamSessionUpdate(s, ch)
	if err != nil {
		return errors.Wrap(err, "error sending team session")
	}

	err = s.updatePerms(ch, userID, nil, nil, s.teams.Data.Permissions, param.Permissions)
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
		return errors.New("cannot load team session ["+ch.ID.String()+"]")
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
