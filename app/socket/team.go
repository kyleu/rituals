package socket

import (
	"fmt"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/estimate"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/retro"
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/standup"
	"github.com/kyleu/rituals.dev/app/team"
	"github.com/kyleu/rituals.dev/app/util"
)

type TeamSessionJoined struct {
	Profile   *util.Profile       `json:"profile"`
	Session   *team.Session       `json:"session"`
	Members   []*member.Entry     `json:"members"`
	Online    []uuid.UUID         `json:"online"`
	Sprints   []*sprint.Session   `json:"sprints"`
	Estimates []*estimate.Session `json:"estimates"`
	Standups  []*standup.Session  `json:"standups"`
	Retros    []*retro.Session    `json:"retros"`
}

func onTeamMessage(s *Service, conn *connection, userID uuid.UUID, cmd string, param interface{}) error {
	var err error
	switch cmd {
	case ClientCmdConnect:
		err = onTeamConnect(s, conn, userID, param.(string))
	case ClientCmdUpdateSession:
		err = onTeamSessionSave(s, *conn.Channel, userID, param.(map[string]interface{}))
	default:
		err = errors.New("unhandled team command [" + cmd + "]")
	}
	return errors.WithStack(errors.Wrap(err, "error handling team message"))
}

func onTeamSessionSave(s *Service, ch channel, userID uuid.UUID, param map[string]interface{}) error {
	title := util.ServiceTitle(param["title"].(string))
	s.logger.Debug(fmt.Sprintf("saving team session [%s]", title))

	err := s.teams.UpdateSession(ch.ID, title, userID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error updating team session"))
	}

	err = sendTeamSessionUpdate(s, ch)
	return errors.WithStack(errors.Wrap(err, "error sending team session"))
}

func sendTeamUpdate(s *Service, ch channel, curr *uuid.UUID, tm *team.Session) error {
	err := s.WriteChannel(ch, &Message{Svc: ch.Svc, Cmd: ServerCmdTeamUpdate, Param: tm})
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error writing team update message"))
	}

	err = s.SendContentUpdate(util.SvcTeam.Key, curr)
	if err != nil {
		return err
	}
	if tm != nil {
		err = s.SendContentUpdate(util.SvcTeam.Key, &tm.ID)
	}
	return err
}

func sendTeamSessionUpdate(s *Service, ch channel) error {
	sess, err := s.teams.GetByID(ch.ID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error finding team session ["+ch.ID.String()+"]"))
	}
	if sess == nil {
		return errors.WithStack(errors.Wrap(err, "cannot load team session ["+ch.ID.String()+"]"))
	}
	msg := Message{Svc: util.SvcTeam.Key, Cmd: ServerCmdSessionUpdate, Param: sess}
	err = s.WriteChannel(ch, &msg)
	return errors.WithStack(errors.Wrap(err, "error sending team session"))
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
