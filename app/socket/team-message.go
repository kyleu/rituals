package socket

import (
	"emperror.dev/errors"
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/model/permission"
	"github.com/kyleu/rituals.dev/app/util"
)

type teamSessionSaveParams struct {
	Title       string                 `json:"title"`
	Permissions permission.Permissions `json:"permissions"`
}

func onTeamMessage(s *Service, conn *Connection, cmd string, param json.RawMessage) error {
	var err error
	userID := conn.Profile.UserID
	switch cmd {
	case ClientCmdConnect:
		var u uuid.UUID
		util.FromJSON(param, &u, s.Logger)
		err = onTeamConnect(s, conn, u)
	case ClientCmdUpdateSession:
		tss := teamSessionSaveParams{}
		util.FromJSON(param, &tss, s.Logger)
		err = onTeamSessionSave(s, *conn.Channel, userID, tss)
	case ClientCmdRemoveMember:
		var u uuid.UUID
		util.FromJSON(param, &u, s.Logger)
		err = onRemoveMember(s, s.teams.Data.Members, *conn.Channel, userID, u)
	default:
		err = errors.New("unhandled team command [" + cmd + "]")
	}
	return errors.Wrap(err, "error handling team message")
}

func sendTeams(s *Service, conn *Connection, userID uuid.UUID) error {
	teams := s.teams.GetByMember(userID, nil)
	return s.WriteMessage(conn.ID, NewMessage(util.SvcSystem, ServerCmdTeams, teams))
}
