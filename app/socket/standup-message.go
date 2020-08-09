package socket

import (
	"encoding/json"

	"github.com/kyleu/npn/npncore"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/permission"
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

func onStandupMessage(s *Service, conn *connection, cmd string, param json.RawMessage) error {
	dataSvc := s.standups
	var err error
	userID := conn.Profile.UserID

	switch cmd {
	case ClientCmdConnect:
		var u uuid.UUID
		_ = npncore.FromJSON(param, &u)
		err = onStandupConnect(s, conn, u)
	case ClientCmdUpdateSession:
		sss := standupSessionSaveParams{}
		_ = npncore.FromJSON(param, &sss)
		err = onStandupSessionSave(s, *conn.Channel, userID, sss)
	case ClientCmdUpdateMember:
		u := updateMemberParams{}
		_ = npncore.FromJSON(param, &u)
		err = onUpdateMember(s, dataSvc.Data.Members, *conn.Channel, userID, u)
	case ClientCmdRemoveMember:
		var u uuid.UUID
		_ = npncore.FromJSON(param, &u)
		err = onRemoveMember(s, dataSvc.Data.Members, *conn.Channel, userID, u)
	case ClientCmdAddReport:
		arp := addReportParams{}
		_ = npncore.FromJSON(param, &arp)
		err = onAddReport(s, *conn.Channel, userID, arp)
	case ClientCmdUpdateReport:
		erp := editReportParams{}
		_ = npncore.FromJSON(param, &erp)
		err = onEditReport(s, *conn.Channel, userID, erp)
	case ClientCmdRemoveReport:
		var u uuid.UUID
		_ = npncore.FromJSON(param, &u)
		err = onRemoveReport(s, *conn.Channel, userID, u)
	default:
		err = errors.New("unhandled standup command [" + cmd + "]")
	}
	return errors.Wrap(err, "error handling standup message")
}
