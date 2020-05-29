package socket

import (
	"encoding/json"
	"fmt"
	"time"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/util"
)

func onSprintMessage(s *Service, conn *connection, cmd string, param json.RawMessage) error {
	var err error
	userID := conn.Profile.UserID

	switch cmd {
	case ClientCmdConnect:
		var u uuid.UUID
		util.FromJSON(param, &u, s.logger)
		err = onSprintConnect(s, conn, u)
	case ClientCmdUpdateSession:
		sss := sprintSessionSaveParams{}
		util.FromJSON(param, &sss, s.logger)
		err = onSprintSessionSave(s, *conn.Channel, userID, sss)
	case ClientCmdRemoveMember:
		var u uuid.UUID
		util.FromJSON(param, &u, s.logger)
		err = onRemoveMember(s, s.sprints.Members, *conn.Channel, userID, u)
	default:
		err = errors.New("unhandled sprint command [" + cmd + "]")
	}
	return errors.Wrap(err, "error handling sprint message")
}

func onSprintSessionSave(s *Service, ch channel, userID uuid.UUID, param sprintSessionSaveParams) error {
	title := util.ServiceTitle(util.SvcSprint, param.Title)

	curr, err := s.sprints.GetByID(ch.ID)
	if err != nil {
		return errors.Wrap(err, "error loading sprint session ["+ch.ID.String()+"] for update")
	}

	teamID := util.GetUUIDFromString(param.TeamID)
	var startDate *time.Time
	var endDate *time.Time

	if len(param.StartDate) > 0 {
		d, e := util.FromYMD(param.StartDate)
		if e == nil {
			startDate = d
		}
	}
	if len(param.EndDate) > 0 {
		d, e := util.FromYMD(param.EndDate)
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

	err = s.updatePerms(ch, userID, s.sprints.Permissions, param.Permissions)
	if err != nil {
		return errors.Wrap(err, "error updating permissions")
	}

	return nil
}

func sendSprintUpdate(s *Service, ch channel, curr *uuid.UUID, spr *sprint.Session) error {
	err := s.WriteChannel(ch, NewMessage(ch.Svc, ServerCmdSprintUpdate, spr))
	if err != nil {
		return errors.Wrap(err, "error writing sprint update message")
	}
	err = s.SendContentUpdate(util.SvcSprint, curr)
	if err != nil {
		return err
	}
	if spr != nil {
		err = s.SendContentUpdate(util.SvcSprint, &spr.ID)
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
	err = s.WriteChannel(ch, NewMessage(util.SvcSprint, ServerCmdSessionUpdate, sess))
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
