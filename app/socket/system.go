package socket

import (
	"emperror.dev/errors"
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/util"
)

func onSystemMessage(s *Service, conn *connection, cmd string, param json.RawMessage) error {
	userID := conn.Profile.UserID
	if conn.Profile.UserID != userID {
		return errors.New("received name change for wrong user [" + userID.String() + "]")
	}

	var err error

	switch cmd {
	case ClientCmdPing:
		err = s.WriteMessage(conn.ID, NewMessage(util.SvcSystem, ServerCmdPong, param))
	case ClientCmdUpdateProfile:
		snp := &ParamsSaveName{}
		util.FromJSON(param, snp, s.logger)
		err = saveName(s, conn, userID, snp)
	case ClientCmdGetActions:
		err = sendActions(s, conn)
	case ClientCmdGetTeams:
		err = sendTeams(s, conn, userID)
	case ClientCmdGetSprints:
		err = sendSprints(s, conn, userID)
	default:
		err = errors.New("unhandled system command [" + cmd + "]")
	}
	return errors.Wrap(err, "error handling system message")
}

type ParamsSaveName struct {
	Name   string `json:"name"`
	Choice string `json:"choice"`
}

func saveName(s *Service, conn *connection, userID uuid.UUID, p *ParamsSaveName) error {
	if p.Choice == "global" {
		err := s.UpdateName(userID, p.Name)
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

	if current.Name != p.Name {
		current, err = memberSvc.UpdateName(conn.Channel.ID, userID, p.Name)
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
	actions := s.actions.GetBySvcModel(conn.Svc, *conn.ModelID, nil)
	return s.WriteMessage(conn.ID, NewMessage(util.SvcSystem, ServerCmdActions, actions))
}

func sendTeams(s *Service, conn *connection, userID uuid.UUID) error {
	teams := s.teams.GetByMember(userID, nil)
	return s.WriteMessage(conn.ID, NewMessage(util.SvcSystem, ServerCmdTeams, teams))
}

func sendSprints(s *Service, conn *connection, userID uuid.UUID) error {
	sprints := s.sprints.GetByMember(userID, nil)
	return s.WriteMessage(conn.ID, NewMessage(util.SvcSystem, ServerCmdSprints, sprints))
}

func memberSvcFor(s *Service, svc util.Service) (*member.Service, error) {
	var ret *member.Service

	switch svc {
	case util.SvcTeam:
		ret = s.teams.Members
	case util.SvcSprint:
		ret = s.sprints.Members
	case util.SvcEstimate:
		ret = s.estimates.Members
	case util.SvcStandup:
		ret = s.standups.Members
	case util.SvcRetro:
		ret = s.retros.Members
	default:
		return nil, errors.New(util.IDErrorString(util.KeyService, svc.Key))
	}
	return ret, nil
}
