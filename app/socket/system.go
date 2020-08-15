package socket

import (
	"encoding/json"
	"github.com/kyleu/npn/npnconnection"

	"github.com/kyleu/npn/npncore"

	"github.com/gofrs/uuid"

	"emperror.dev/errors"
	"github.com/kyleu/rituals.dev/app/session"
	"github.com/kyleu/rituals.dev/app/util"
)

func onSystemMessage(s *npnconnection.Service, conn *npnconnection.Connection, cmd string, param json.RawMessage) error {
	userID := conn.Profile.UserID
	if conn.Profile.UserID != userID {
		return errors.New("received name change for wrong user [" + userID.String() + "]")
	}

	var err error

	switch cmd {
	case ClientCmdPing:
		err = s.WriteMessage(conn.ID, npnconnection.NewMessage(util.SvcSystem.Key, ServerCmdPong, param))
	case ClientCmdAddComment:
		acp := addCommentParams{}
		_ = npncore.FromJSON(param, &acp)
		err = onAddComment(s, *conn.Channel, userID, acp)
	case ClientCmdUpdateComment:
		ucp := updateCommentParams{}
		_ = npncore.FromJSON(param, &ucp)
		err = onUpdateComment(s, *conn.Channel, userID, ucp)
	case ClientCmdRemoveComment:
		var u uuid.UUID
		_ = npncore.FromJSON(param, &u)
		err = onRemoveComment(s, *conn.Channel, userID, u)
	case ClientCmdUpdateProfile:
		snp := &saveProfileParams{}
		_ = npncore.FromJSON(param, snp)
		err = saveProfile(s, conn, userID, snp)
	case ClientCmdGetActions:
		err = sendActions(s, conn)
	case ClientCmdGetTeams:
		err = sendTeams(s, conn, userID)
	case ClientCmdGetSprints:
		err = sendSprints(s, conn, userID)
	case ClientCmdSetActive:
		err = setActive(s, conn, session.StatusActive, userID)
	default:
		err = errors.New("unhandled system command [" + cmd + "]")
	}
	return errors.Wrap(err, "error handling system message")
}

func sendActions(s *npnconnection.Service, conn *npnconnection.Connection) error {
	if conn.ModelID == nil {
		return errors.New("no active model for connection [" + conn.ID.String() + "]")
	}
	actions := actions(s).GetBySvcModel(conn.Svc, *conn.ModelID, nil)
	return s.WriteMessage(conn.ID, npnconnection.NewMessage(util.SvcSystem.Key, ServerCmdActions, actions))
}

func setActive(s *npnconnection.Service, conn *npnconnection.Connection, status session.Status, userID uuid.UUID) error {
	return dataFor(s, conn.Svc).Members.UpdateStatus(conn.Channel.ID, status.Key, userID)
}

func dataFor(s *npnconnection.Service, svc string) *session.DataServices {
	switch svc {
	case util.SvcTeam.Key:
		return teams(s).Data
	case util.SvcSprint.Key:
		return sprints(s).Data
	case util.SvcEstimate.Key:
		return estimates(s).Data
	case util.SvcStandup.Key:
		return standups(s).Data
	case util.SvcRetro.Key:
		return retros(s).Data
	default:
		return nil
	}
}
