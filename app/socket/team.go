package socket

import (
	"encoding/json"
	"fmt"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/team"
	"github.com/kyleu/rituals.dev/app/util"
)

func onTeamMessage(s *Service, conn *connection, cmd string, param json.RawMessage) error {
	var err error
	userID := conn.Profile.UserID
	switch cmd {
	case ClientCmdConnect:
		var u uuid.UUID
		util.FromJSON(param, &u, s.logger)
		err = onTeamConnect(s, conn, u)
	case ClientCmdUpdateSession:
		tss := teamSessionSaveParams{}
		util.FromJSON(param, &tss, s.logger)
		err = onTeamSessionSave(s, *conn.Channel, userID, tss)
	case ClientCmdRemoveMember:
		var u uuid.UUID
		util.FromJSON(param, &u, s.logger)
		err = onRemoveMember(s, s.teams.Members, *conn.Channel, userID, u)
	default:
		err = errors.New("unhandled team command [" + cmd + "]")
	}
	return errors.Wrap(err, "error handling team message")
}

func onTeamSessionSave(s *Service, ch channel, userID uuid.UUID, param teamSessionSaveParams) error {
	title := util.ServiceTitle(util.SvcTeam, param.Title)
	s.logger.Debug(fmt.Sprintf("saving team session [%s]", title))

	err := s.teams.UpdateSession(ch.ID, title, userID)
	if err != nil {
		return errors.Wrap(err, "error updating team session")
	}

	err = sendTeamSessionUpdate(s, ch)
	if err != nil {
		return errors.Wrap(err, "error sending team session")
	}

	err = s.updatePerms(ch, userID, s.teams.Permissions, param.Permissions)
	if err != nil {
		return errors.Wrap(err, "error updating permissions")
	}

	return nil
}

func sendTeamUpdate(s *Service, ch channel, curr *uuid.UUID, tm *team.Session) error {
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

func sendTeamSessionUpdate(s *Service, ch channel) error {
	sess, err := s.teams.GetByID(ch.ID)
	if err != nil {
		return errors.Wrap(err, "error finding team session ["+ch.ID.String()+"]")
	}
	if sess == nil {
		return errors.Wrap(err, "cannot load team session ["+ch.ID.String()+"]")
	}
	err = s.WriteChannel(ch, NewMessage(util.SvcTeam, ServerCmdSessionUpdate, sess))
	return errors.Wrap(err, "error sending team session")
}

func getTeamOpt(s *Service, teamID *uuid.UUID) *team.Session {
	if teamID == nil {
		return nil
	}
	tm, err := s.teams.GetByID(*teamID)
	if err != nil {
		s.logger.Warn(fmt.Sprintf("error getting associated team [%v]: %+v", teamID, err))
	}
	return tm
}
