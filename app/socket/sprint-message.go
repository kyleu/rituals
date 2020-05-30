package socket

import (
	"emperror.dev/errors"
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/model/permission"
	"github.com/kyleu/rituals.dev/app/util"
)

type sprintSessionSaveParams struct {
	Title       string                 `json:"title"`
	StartDate   string                 `json:"startDate"`
	EndDate     string                 `json:"endDate"`
	TeamID      string                 `json:"teamID"`
	Permissions permission.Permissions `json:"permissions"`
}

func onSprintMessage(s *Service, conn *Connection, cmd string, param json.RawMessage) error {
	var err error
	userID := conn.Profile.UserID

	switch cmd {
	case ClientCmdConnect:
		var u uuid.UUID
		util.FromJSON(param, &u, s.Logger)
		err = onSprintConnect(s, conn, u)
	case ClientCmdUpdateSession:
		sss := sprintSessionSaveParams{}
		util.FromJSON(param, &sss, s.Logger)
		err = onSprintSessionSave(s, *conn.Channel, userID, sss)
	case ClientCmdRemoveMember:
		var u uuid.UUID
		util.FromJSON(param, &u, s.Logger)
		err = onRemoveMember(s, s.sprints.Data.Members, *conn.Channel, userID, u)
	default:
		err = errors.New("unhandled sprint command [" + cmd + "]")
	}
	return errors.Wrap(err, "error handling sprint message")
}
