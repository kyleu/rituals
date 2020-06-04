package socket

import (
	"encoding/json"

	"emperror.dev/errors"
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

func onSprintMessage(s *Service, conn *connection, cmd string, param json.RawMessage) error {
	dataSvc := s.sprints
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
	case ClientCmdUpdateMember:
		u := updateMemberParams{}
		util.FromJSON(param, &u, s.Logger)
		err = onUpdateMember(s, dataSvc.Data.Members, *conn.Channel, userID, u)
	case ClientCmdRemoveMember:
		var u uuid.UUID
		util.FromJSON(param, &u, s.Logger)
		err = onRemoveMember(s, dataSvc.Data.Members, *conn.Channel, userID, u)
	default:
		err = errors.New("unhandled sprint command [" + cmd + "]")
	}
	return errors.Wrap(err, "error handling sprint message")
}

func sendSprints(s *Service, conn *connection, userID uuid.UUID) error {
	sprints := s.sprints.GetByMember(userID, nil)
	return s.WriteMessage(conn.ID, NewMessage(util.SvcSystem, ServerCmdSprints, sprints))
}
