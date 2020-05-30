package socket

import (
	"emperror.dev/errors"
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/model/permission"
	"github.com/kyleu/rituals.dev/app/util"
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

func onRetroMessage(s *Service, conn *Connection, cmd string, param json.RawMessage) error {
	var err error
	userID := conn.Profile.UserID

	switch cmd {
	case ClientCmdConnect:
		var u uuid.UUID
		util.FromJSON(param, &u, s.Logger)
		err = onRetroConnect(s, conn, u)
	case ClientCmdUpdateSession:
		rss := retroSessionSaveParams{}
		util.FromJSON(param, &rss, s.Logger)
		err = onRetroSessionSave(s, *conn.Channel, userID, rss)
	case ClientCmdRemoveMember:
		var u uuid.UUID
		util.FromJSON(param, &u, s.Logger)
		err = onRemoveMember(s, s.retros.Data.Members, *conn.Channel, userID, u)
	case ClientCmdAddFeedback:
		afp := addFeedbackParams{}
		util.FromJSON(param, &afp, s.Logger)
		err = onAddFeedback(s, *conn.Channel, userID, afp)
	case ClientCmdUpdateFeedback:
		efp := editFeedbackParams{}
		util.FromJSON(param, &efp, s.Logger)
		err = onEditFeedback(s, *conn.Channel, userID, efp)
	case ClientCmdRemoveFeedback:
		var u uuid.UUID
		util.FromJSON(param, &u, s.Logger)
		err = onRemoveFeedback(s, *conn.Channel, userID, u)
	default:
		err = errors.New("unhandled retro command [" + cmd + "]")
	}
	return errors.Wrap(err, "error handling retro message")
}
