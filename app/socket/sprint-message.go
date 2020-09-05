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

type sprintSessionSaveParams struct {
	Title       string                 `json:"title"`
	StartDate   string                 `json:"startDate"`
	EndDate     string                 `json:"endDate"`
	TeamID      string                 `json:"teamID"`
	Permissions permission.Permissions `json:"permissions"`
}

func onSprintMessage(s *npnconnection.Service, a *auth.Service, conn *npnconnection.Connection, cmd string, param json.RawMessage) error {
	dataSvc := sprints(s)
	var err error
	userID := conn.Profile.UserID

	switch cmd {
	case ClientCmdConnect:
		var u uuid.UUID
		_ = npncore.FromJSON(param, &u)
		err = onSprintConnect(s, conn, u)
	case ClientCmdUpdateSession:
		sss := sprintSessionSaveParams{}
		_ = npncore.FromJSON(param, &sss)
		err = onSprintSessionSave(s, a, *conn.Channel, userID, sss)
	case ClientCmdUpdateMember:
		u := updateMemberParams{}
		_ = npncore.FromJSON(param, &u)
		err = onUpdateMember(s, dataSvc.Data.Members, *conn.Channel, userID, u)
	case ClientCmdRemoveMember:
		var u uuid.UUID
		_ = npncore.FromJSON(param, &u)
		err = onRemoveMember(s, dataSvc.Data.Members, *conn.Channel, userID, u)
	default:
		err = errors.New("unhandled sprint command [" + cmd + "]")
	}
	return errors.Wrap(err, "error handling sprint message")
}

func sendSprints(s *npnconnection.Service, conn *npnconnection.Connection, userID uuid.UUID) error {
	sprints := sprints(s).GetByMember(userID, nil)
	return s.WriteMessage(conn.ID, npnconnection.NewMessage(util.SvcSystem.Key, ServerCmdSprints, sprints))
}
