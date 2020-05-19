package socket

import (
	"emperror.dev/errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/sprint"
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
	case ClientCmdGetSprints:
		err = sendSprints(s, conn, userID)
	case ClientCmdSetSprint:
		err = setSprint(s, conn, userID, param)
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

func sendSprints(s *Service, conn *connection, userID uuid.UUID) error {
	sprints, err := s.sprints.GetByMember(userID, nil)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "cannot get sprints"))
	}
	actionsMsg := Message{Svc: util.SvcSystem.Key, Cmd: ServerCmdSprints, Param: sprints}
	return s.WriteMessage(conn.ID, &actionsMsg)
}

func setSprint(s *Service, conn *connection, userID uuid.UUID, param interface{}) error {
	if conn.ModelID == nil {
		return errors.WithStack(errors.New("no active model"))
	}

	var spr *sprint.Session

	if param != nil {
		sprintID, err := uuid.FromString(param.(string))
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "sprintID isn't a valid UUID"))
		}
		spr, err = s.sprints.AssignSprint(conn.Svc, conn.ModelID, userID, &sprintID)
		if err != nil {
			return errors.WithStack(errors.Wrap(err, fmt.Sprintf("cannot remove %v [%v] from sprint [%v]", conn.Svc, conn.ModelID, sprintID)))
		}

	} else {
		sprP, err := s.sprints.AssignSprint(conn.Svc, conn.ModelID, userID, nil)
		spr = sprP
		if err != nil {
			return errors.WithStack(errors.Wrap(err, fmt.Sprintf("cannot remove sprint from %v [%v]", conn.Svc, conn.ModelID)))
		}
	}
	return sendSprintUpdate(s, *conn.Channel, spr)
}

func sendSprintUpdate(s *Service, ch channel, spr *sprint.Session) error {
	err := s.WriteChannel(ch, &Message{Svc: ch.Svc, Cmd: ServerCmdSprintUpdate, Param: spr})
	return errors.WithStack(errors.Wrap(err, "error writing sprint update message"))
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
