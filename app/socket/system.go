package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/util"
)

func onSystemMessage(s *Service, conn *connection, userID uuid.UUID, cmd string, param interface{}) error {
	if conn.Profile.UserID != userID {
		return errors.WithStack(errors.New("received name change for wrong user [" + userID.String() + "]"))
	}
	var err error

	switch cmd {
	case ClientCmdPing:
		msg := Message{Svc: util.SvcSystem.Key, Cmd: ServerCmdPong, Param: param}
		err = s.WriteMessage(conn.ID, &msg)
	case ClientCmdUpdateProfile:
		err = saveName(s, conn, userID, param.(map[string]interface{}))
	case ClientCmdGetActions:
		err = sendActions(s, conn)
	case ClientCmdGetTeams:
		err = sendTeams(s, conn, userID)
	case ClientCmdGetSprints:
		err = sendSprints(s, conn, userID)
	default:
		err = errors.New("unhandled system command [" + cmd + "]")
	}
	return errors.WithStack(errors.Wrap(err, "error handling system message"))
}

func saveName(s *Service, conn *connection, userID uuid.UUID, o map[string]interface{}) error {
	name := o["name"].(string)
	choice := o["choice"].(string)
	if choice == "global" {
		err := s.UpdateName(userID, name)
		if err != nil {
			return err
		}
	}
	memberSvc, err := memberSvcFor(s, conn.Channel.Svc)
	if err != nil {
		return err
	}

	current, err := memberSvc.Get(conn.Channel.ID, userID)
	if err != nil {
		return err
	}

	if current.Name != name {
		current, err = memberSvc.UpdateName(conn.Channel.ID, userID, name)
		if err != nil {
			return err
		}
	}

	if conn.Channel == nil {
		return errors.New("no channel registered for [" + conn.ID.String() + "]")
	}
	return s.sendMemberUpdate(*conn.Channel, current)
}

func sendActions(s *Service, conn *connection) error {
	if conn.ModelID == nil {
		return errors.New("no active model for connection [" + conn.ID.String() + "]")
	}
	actions, err := s.actions.GetBySvcModel(conn.Svc, *conn.ModelID, nil)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "cannot get actions"))
	}
	actionsMsg := Message{Svc: util.SvcSystem.Key, Cmd: ServerCmdActions, Param: actions}
	return s.WriteMessage(conn.ID, &actionsMsg)
}

func sendTeams(s *Service, conn *connection, userID uuid.UUID) error {
	teams, err := s.teams.GetByMember(userID, nil)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "cannot get teams"))
	}
	actionsMsg := Message{Svc: util.SvcSystem.Key, Cmd: ServerCmdTeams, Param: teams}
	return s.WriteMessage(conn.ID, &actionsMsg)
}

func sendSprints(s *Service, conn *connection, userID uuid.UUID) error {
	sprints, err := s.sprints.GetByMember(userID, nil)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "cannot get sprints"))
	}
	actionsMsg := Message{Svc: util.SvcSystem.Key, Cmd: ServerCmdSprints, Param: sprints}
	return s.WriteMessage(conn.ID, &actionsMsg)
}

func memberSvcFor(s *Service, svc string) (*member.Service, error) {
	var ret *member.Service

	switch svc {
	case util.SvcTeam.Key:
		ret = s.teams.Members
	case util.SvcSprint.Key:
		ret = s.sprints.Members
	case util.SvcEstimate.Key:
		ret = s.estimates.Members
	case util.SvcStandup.Key:
		ret = s.standups.Members
	case util.SvcRetro.Key:
		ret = s.retros.Members
	default:
		return nil, errors.New("invalid service [" + svc + "]")
	}
	return ret, nil
}

func getUUIDPointer(m map[string]interface{}, key string) *uuid.UUID {
	retOut, ok := m[key]
	if !ok {
		return nil
	}
	return util.GetUUIDFromString(retOut.(string))
}
