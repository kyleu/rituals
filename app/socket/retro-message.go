package socket

import (
	"encoding/json"
	"github.com/kyleu/npn/npnconnection"

	"github.com/kyleu/npn/npncore"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/permission"
)

type retroSessionSaveParams struct {
	Title       string                 `json:"title"`
	Categories  string                 `json:"categories"`
	SprintID    string                 `json:"sprintID"`
	TeamID      string                 `json:"teamID"`
	Permissions permission.Permissions `json:"permissions"`
}

type addFeedbackParams struct {
	Category string `json:"category"`
	Content  string `json:"content"`
}

type editFeedbackParams struct {
	ID       uuid.UUID `json:"id"`
	Category string    `json:"category"`
	Content  string    `json:"content"`
}

func onRetroMessage(s *npnconnection.Service, conn *npnconnection.Connection, cmd string, param json.RawMessage) error {
	dataSvc := retros(s)
	var err error
	userID := conn.Profile.UserID

	switch cmd {
	case ClientCmdConnect:
		var u uuid.UUID
		_ = npncore.FromJSON(param, &u)
		err = onRetroConnect(s, conn, u)
	case ClientCmdUpdateSession:
		rss := retroSessionSaveParams{}
		_ = npncore.FromJSON(param, &rss)
		err = onRetroSessionSave(s, *conn.Channel, userID, rss)
	case ClientCmdUpdateMember:
		u := updateMemberParams{}
		_ = npncore.FromJSON(param, &u)
		err = onUpdateMember(s, dataSvc.Data.Members, *conn.Channel, userID, u)
	case ClientCmdRemoveMember:
		var u uuid.UUID
		_ = npncore.FromJSON(param, &u)
		err = onRemoveMember(s, dataSvc.Data.Members, *conn.Channel, userID, u)
	case ClientCmdAddFeedback:
		afp := addFeedbackParams{}
		_ = npncore.FromJSON(param, &afp)
		err = onAddFeedback(s, *conn.Channel, userID, afp)
	case ClientCmdUpdateFeedback:
		efp := editFeedbackParams{}
		_ = npncore.FromJSON(param, &efp)
		err = onEditFeedback(s, *conn.Channel, userID, efp)
	case ClientCmdRemoveFeedback:
		var u uuid.UUID
		_ = npncore.FromJSON(param, &u)
		err = onRemoveFeedback(s, *conn.Channel, userID, u)
	default:
		err = errors.New("unhandled retro command [" + cmd + "]")
	}
	return errors.Wrap(err, "error handling retro message")
}
