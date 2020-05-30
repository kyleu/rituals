package socket

import (
	"emperror.dev/errors"
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/model/permission"
	"github.com/kyleu/rituals.dev/app/util"
)

type standupSessionSaveParams struct {
	Title       string                 `json:"title"`
	SprintID    string                 `json:"sprintID"`
	TeamID      string                 `json:"teamID"`
	Permissions permission.Permissions `json:"permissions"`
}

type addReportParams struct {
	D       string `json:"d"`
	Content string `json:"content"`
}

type editReportParams struct {
	ID      uuid.UUID `json:"id"`
	D       string    `json:"d"`
	Content string    `json:"content"`
}

func onStandupMessage(s *Service, conn *Connection, cmd string, param json.RawMessage) error {
	var err error
	userID := conn.Profile.UserID

	switch cmd {
	case ClientCmdConnect:
		var u uuid.UUID
		util.FromJSON(param, &u, s.Logger)
		err = onStandupConnect(s, conn, u)
	case ClientCmdUpdateSession:
		sss := standupSessionSaveParams{}
		util.FromJSON(param, &sss, s.Logger)
		err = onStandupSessionSave(s, *conn.Channel, userID, sss)
	case ClientCmdRemoveMember:
		var u uuid.UUID
		util.FromJSON(param, &u, s.Logger)
		err = onRemoveMember(s, s.standups.Data.Members, *conn.Channel, userID, u)
	case ClientCmdAddReport:
		arp := addReportParams{}
		util.FromJSON(param, &arp, s.Logger)
		err = onAddReport(s, *conn.Channel, userID, arp)
	case ClientCmdUpdateReport:
		erp := editReportParams{}
		util.FromJSON(param, &erp, s.Logger)
		err = onEditReport(s, *conn.Channel, userID, erp)
	case ClientCmdRemoveReport:
		var u uuid.UUID
		util.FromJSON(param, &u, s.Logger)
		err = onRemoveReport(s, *conn.Channel, userID, u)
	default:
		err = errors.New("unhandled standup command [" + cmd + "]")
	}
	return errors.Wrap(err, "error handling standup message")
}
