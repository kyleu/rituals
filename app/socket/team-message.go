package socket

import (
	"encoding/json"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/permission"
	"github.com/kyleu/rituals.dev/app/util"
)

type teamSessionSaveParams struct {
	Title       string                 `json:"title"`
	Permissions permission.Permissions `json:"permissions"`
}

func onTeamMessage(s *Service, conn *connection, cmd string, param json.RawMessage) error {
	dataSvc := s.teams
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
	case ClientCmdUpdateMember:
		u := updateMemberParams{}
		util.FromJSON(param, &u, s.Logger)
		err = onUpdateMember(s, dataSvc.Data.Members, *conn.Channel, userID, u)
	case ClientCmdRemoveMember:
		var u uuid.UUID
		util.FromJSON(param, &u, s.Logger)
		err = onRemoveMember(s, dataSvc.Data.Members, *conn.Channel, userID, u)
	default:
		err = errors.New("unhandled team command [" + cmd + "]")
	}
	return errors.Wrap(err, "error handling team message")
}

func sendTeams(s *Service, conn *connection, userID uuid.UUID) error {
	teams := s.teams.GetByMember(userID, nil)
	return s.WriteMessage(conn.ID, NewMessage(util.SvcSystem, ServerCmdTeams, teams))
}
