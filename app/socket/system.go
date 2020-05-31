package socket

import (
	"emperror.dev/errors"
	"encoding/json"
	"github.com/kyleu/rituals.dev/app/model/session"
	"github.com/kyleu/rituals.dev/app/util"
)

func onSystemMessage(s *Service, conn *Connection, cmd string, param json.RawMessage) error {
	userID := conn.Profile.UserID
	if conn.Profile.UserID != userID {
		return errors.New("received name change for wrong user [" + userID.String() + "]")
	}

	var err error

	switch cmd {
	case ClientCmdPing:
		err = s.WriteMessage(conn.ID, NewMessage(util.SvcSystem, ServerCmdPong, param))
	case ClientCmdAddComment:
		acp := addCommentParams{}
		util.FromJSON(param, &acp, s.Logger)
		err = onAddComment(s, *conn.Channel, userID, acp)
	case ClientCmdUpdateComment:
		ucp := updateCommentParams{}
		util.FromJSON(param, &ucp, s.Logger)
		err = onUpdateComment(s, *conn.Channel, userID, ucp)
	case ClientCmdUpdateProfile:
		snp := &saveProfileParams{}
		util.FromJSON(param, snp, s.Logger)
		err = saveProfile(s, conn, userID, snp)
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

func sendActions(s *Service, conn *Connection) error {
	if conn.ModelID == nil {
		return errors.New("no active model for connection [" + conn.ID.String() + "]")
	}
	actions := s.actions.GetBySvcModel(conn.Svc, *conn.ModelID, nil)
	return s.WriteMessage(conn.ID, NewMessage(util.SvcSystem, ServerCmdActions, actions))
}

func dataFor(s *Service, svc util.Service) *session.DataServices {
	switch svc {
	case util.SvcTeam:
		return s.teams.Data
	case util.SvcSprint:
		return s.sprints.Data
	case util.SvcEstimate:
		return s.estimates.Data
	case util.SvcStandup:
		return s.standups.Data
	case util.SvcRetro:
		return s.retros.Data
	default:
		return nil
	}
}
