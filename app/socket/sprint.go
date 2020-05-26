package socket

import (
	"fmt"
	"time"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/util"
)

func onSprintMessage(s *Service, conn *connection, userID uuid.UUID, cmd string, param interface{}) error {
	var err error

	switch cmd {
	case ClientCmdConnect:
		err = onSprintConnect(s, conn, userID, param.(string))
	case ClientCmdUpdateSession:
		err = onSprintSessionSave(s, *conn.Channel, userID, param.(map[string]interface{}))
	case ClientCmdRemoveMember:
		err = onRemoveMember(s, s.sprints.Members, *conn.Channel, userID, param.(string))
	default:
		err = errors.New("unhandled sprint command [" + cmd + "]")
	}
	return errors.Wrap(err, "error handling sprint message")
}

func onSprintSessionSave(s *Service, ch channel, userID uuid.UUID, param map[string]interface{}) error {
	title := util.ServiceTitle(util.SvcSprint, param[util.KeyTitle].(string))

	curr, err := s.sprints.GetByID(ch.ID)
	if err != nil {
		return errors.Wrap(err, "error loading sprint session ["+ch.ID.String()+"] for update")
	}

	teamID := getUUIDPointer(param, util.WithID(util.SvcTeam.Key))
	var startDate *time.Time
	var endDate *time.Time

	startDateP, sok := param["startDate"]
	if sok {
		d, e := util.FromYMD(startDateP.(string))
		if e == nil {
			startDate = d
		}
	}
	endDateP, eok := param["endDate"]
	if eok {
		d, e := util.FromYMD(endDateP.(string))
		if e == nil {
			endDate = d
		}
	}

	s.logger.Debug(fmt.Sprintf("saving sprint session [%s] in team [%s]", title, teamID))

	teamChanged := differentPointerValues(curr.TeamID, teamID)

	err = s.sprints.UpdateSession(ch.ID, title, teamID, startDate, endDate, userID)
	if err != nil {
		return errors.Wrap(err, "error updating sprint session")
	}

	err = sendSprintSessionUpdate(s, ch)
	if err != nil {
		return errors.Wrap(err, "error sending sprint session")
	}

	if teamChanged {
		tm := s.teams.GetByIDPointer(teamID)
		err = sendTeamUpdate(s, ch, curr.TeamID, tm)
		if err != nil {
			return errors.Wrap(err, "error sending team for updated sprint session")
		}
	}

	err = s.updatePerms(ch, userID, s.sprints.Permissions, param)
	if err != nil {
		return errors.Wrap(err, "error updating permissions")
	}

	return nil
}

func sendSprintUpdate(s *Service, ch channel, curr *uuid.UUID, spr *sprint.Session) error {
	err := s.WriteChannel(ch, &Message{Svc: ch.Svc, Cmd: ServerCmdSprintUpdate, Param: spr})
	if err != nil {
		return errors.Wrap(err, "error writing sprint update message")
	}
	err = s.SendContentUpdate(util.SvcSprint.Key, curr)
	if err != nil {
		return err
	}
	if spr != nil {
		err = s.SendContentUpdate(util.SvcSprint.Key, &spr.ID)
	}
	return err
}

func sendSprintSessionUpdate(s *Service, ch channel) error {
	sess, err := s.sprints.GetByID(ch.ID)
	if err != nil {
		return errors.Wrap(err, "error finding sprint session ["+ch.ID.String()+"]")
	}
	if sess == nil {
		return errors.Wrap(err, "cannot load sprint session ["+ch.ID.String()+"]")
	}
	msg := Message{Svc: util.SvcSprint.Key, Cmd: ServerCmdSessionUpdate, Param: sess}
	err = s.WriteChannel(ch, &msg)
	return errors.Wrap(err, "error sending sprint session")
}

func getSprintOpt(s *Service, sprintID *uuid.UUID) *sprint.Session {
	if sprintID == nil {
		return nil
	}
	spr, err := s.sprints.GetByID(*sprintID)
	if err != nil {
		s.logger.Warn(fmt.Sprintf("error getting associated sprint [%v]: %+v", sprintID, err))
	}
	return spr
}
