package socket

import (
	"encoding/json"

	"github.com/kyleu/npn/npnservice/auth"

	"github.com/kyleu/npn/npnconnection"

	"github.com/kyleu/npn/npncore"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/permission"
	"github.com/kyleu/rituals.dev/app/util"
)

type teamSessionSaveParams struct {
	Title       string                 `json:"title"`
	Permissions permission.Permissions `json:"permissions"`
}

func onTeamMessage(s *npnconnection.Service, a auth.Service, conn *npnconnection.Connection, cmd string, param json.RawMessage) error {
	dataSvc := ctx(s).teams
	var err error
	userID := conn.Profile.UserID
	switch cmd {
	case ClientCmdConnect:
		var u uuid.UUID
		_ = npncore.FromJSON(param, &u)
		err = onTeamConnect(s, conn, u)
	case ClientCmdUpdateSession:
		tss := teamSessionSaveParams{}
		_ = npncore.FromJSON(param, &tss)
		err = onTeamSessionSave(s, a, *conn.Channel, userID, tss)
	case ClientCmdUpdateMember:
		u := updateMemberParams{}
		_ = npncore.FromJSON(param, &u)
		err = onUpdateMember(s, dataSvc.Data.Members, *conn.Channel, userID, u)
	case ClientCmdRemoveMember:
		var u uuid.UUID
		_ = npncore.FromJSON(param, &u)
		err = onRemoveMember(s, dataSvc.Data.Members, *conn.Channel, userID, u)
	default:
		err = errors.New("unhandled team command [" + cmd + "]")
	}
	return errors.Wrap(err, "error handling team message")
}

func sendTeams(s *npnconnection.Service, conn *npnconnection.Connection, userID uuid.UUID) error {
	teams := ctx(s).teams.GetByMember(userID, nil)
	return s.WriteMessage(conn.ID, npnconnection.NewMessage(util.SvcSystem.Key, ServerCmdTeams, teams))
}
