package socket

import (
	"emperror.dev/errors"
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/util"
)

func onStandupMessage(s *Service, conn *connection, cmd string, param json.RawMessage) error {
	var err error
	userID := conn.Profile.UserID

	switch cmd {
	case ClientCmdConnect:
		var u uuid.UUID
		util.FromJSON(param, &u, s.logger)
		err = onStandupConnect(s, conn, u)
	case ClientCmdUpdateSession:
		sss := standupSessionSaveParams{}
		util.FromJSON(param, &sss, s.logger)
		err = onStandupSessionSave(s, *conn.Channel, userID, sss)
	case ClientCmdRemoveMember:
		var u uuid.UUID
		util.FromJSON(param, &u, s.logger)
		err = onRemoveMember(s, s.standups.Members, *conn.Channel, userID, u)
	case ClientCmdAddReport:
		arp := addReportParams{}
		util.FromJSON(param, &arp, s.logger)
		err = onAddReport(s, *conn.Channel, userID, arp)
	case ClientCmdUpdateReport:
		erp := editReportParams{}
		util.FromJSON(param, &erp, s.logger)
		err = onEditReport(s, *conn.Channel, userID, erp)
	case ClientCmdRemoveReport:
		var u uuid.UUID
		util.FromJSON(param, &u, s.logger)
		err = onRemoveReport(s, *conn.Channel, userID, u)
	default:
		err = errors.New("unhandled standup command [" + cmd + "]")
	}
	return errors.Wrap(err, "error handling standup message")
}

func onStandupSessionSave(s *Service, ch channel, userID uuid.UUID, param standupSessionSaveParams) error {
	title := util.ServiceTitle(util.SvcStandup, param.Title)

	sprintID := util.GetUUIDFromString(param.SprintID)
	teamID := util.GetUUIDFromString(param.TeamID)

	msg := "saving standup session [%s] with sprint [%s] and team [%s]"
	s.logger.Debug(fmt.Sprintf(msg, title, sprintID, teamID))

	curr, err := s.standups.GetByID(ch.ID)
	if err != nil {
		return errors.Wrap(err, "error loading standup session ["+ch.ID.String()+"] for update")
	}

	teamChanged := differentPointerValues(curr.TeamID, teamID)
	sprintChanged := differentPointerValues(curr.SprintID, sprintID)

	err = s.standups.UpdateSession(ch.ID, title, teamID, sprintID, userID)
	if err != nil {
		return errors.Wrap(err, "error updating standup session")
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

	err = s.updatePerms(ch, userID, s.standups.Permissions, param.Permissions)
	if err != nil {
		return errors.Wrap(err, "error updating permissions")
	}

	return nil
}

func sendStandupSessionUpdate(s *Service, ch channel) error {
	sess, err := s.standups.GetByID(ch.ID)
	if err != nil {
		return errors.Wrap(err, "error finding standup session ["+ch.ID.String()+"]")
	}
	if sess == nil {
		return errors.Wrap(err, "cannot load standup session ["+ch.ID.String()+"]")
	}
	err = s.WriteChannel(ch, NewMessage(util.SvcStandup, ServerCmdSessionUpdate, sess))
	return errors.Wrap(err, "error sending standup session")
}
