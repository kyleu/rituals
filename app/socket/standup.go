package socket

import (
	"fmt"

	"github.com/kyleu/rituals.dev/app/team"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/standup"
	"github.com/kyleu/rituals.dev/app/util"
)

type StandupSessionJoined struct {
	Profile *util.Profile     `json:"profile"`
	Session *standup.Session  `json:"session"`
	Team    *team.Session     `json:"team"`
	Sprint  *sprint.Session   `json:"sprint"`
	Members []*member.Entry   `json:"members"`
	Online  []uuid.UUID       `json:"online"`
	Reports []*standup.Report `json:"reports"`
}

func onStandupMessage(s *Service, conn *connection, userID uuid.UUID, cmd string, param interface{}) error {
	var err error
	switch cmd {
	case ClientCmdConnect:
		err = onStandupConnect(s, conn, userID, param.(string))
	case ClientCmdUpdateSession:
		err = onStandupSessionSave(s, *conn.Channel, userID, param.(map[string]interface{}))
	case ClientCmdAddReport:
		err = onAddReport(s, *conn.Channel, userID, param.(map[string]interface{}))
	case ClientCmdUpdateReport:
		err = onEditReport(s, *conn.Channel, userID, param.(map[string]interface{}))
	case ClientCmdRemoveReport:
		err = onRemoveReport(s, *conn.Channel, userID, param.(string))
	default:
		err = errors.New("unhandled standup command [" + cmd + "]")
	}
	return errors.WithStack(errors.Wrap(err, "error handling standup message"))
}

func onStandupSessionSave(s *Service, ch channel, userID uuid.UUID, param map[string]interface{}) error {
	title := util.ServiceTitle(param["title"].(string))
	s.logger.Debug(fmt.Sprintf("saving standup session [%s]", title))

	sprintID := getUUIDPointer(param, "sprintID")
	teamID := getUUIDPointer(param, "teamID")

	s.logger.Debug(fmt.Sprintf("saving standup session [%s] with sprint [%s]", title, sprintID))

	curr, err := s.standups.GetByID(ch.ID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error loading standup session ["+ch.ID.String()+"] for update"))
	}

	teamChanged := differentPointerValues(curr.TeamID, teamID)
	sprintChanged := differentPointerValues(curr.SprintID, sprintID)

	err = s.standups.UpdateSession(ch.ID, title, teamID, sprintID, userID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error updating standup session"))
	}

	err = sendStandupSessionUpdate(s, ch)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error sending standup session"))
	}

	if teamChanged {
		tm := s.teams.GetByIDPointer(teamID)
		err = sendTeamUpdate(s, ch, curr.TeamID, tm)
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "error sending team for updated retro session"))
		}
	}

	if sprintChanged {
		spr := s.sprints.GetByIDPointer(sprintID)
		err = sendSprintUpdate(s, ch, curr.SprintID, spr)
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "error sending sprint for updated standup session"))
		}
	}

	return nil
}

func sendStandupSessionUpdate(s *Service, ch channel) error {
	sess, err := s.standups.GetByID(ch.ID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error finding standup session ["+ch.ID.String()+"]"))
	}
	if sess == nil {
		return errors.WithStack(errors.Wrap(err, "cannot load standup session ["+ch.ID.String()+"]"))
	}
	msg := Message{Svc: util.SvcStandup.Key, Cmd: ServerCmdSessionUpdate, Param: sess}
	err = s.WriteChannel(ch, &msg)
	return errors.WithStack(errors.Wrap(err, "error sending standup session"))
}
