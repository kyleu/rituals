package socket

import (
	"fmt"
	"strings"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/retro"
	"github.com/kyleu/rituals.dev/app/util"
)

type RetroSessionJoined struct {
	Profile  *util.Profile    `json:"profile"`
	Session  *retro.Session   `json:"session"`
	Members  []*member.Entry   `json:"members"`
	Online   []uuid.UUID      `json:"online"`
	Feedback []*retro.Feedback `json:"feedback"`
}

func onRetroMessage(s *Service, conn *connection, userID uuid.UUID, cmd string, param interface{}) error {
	var err error
	switch cmd {
	case util.ClientCmdConnect:
		err = onRetroConnect(s, conn, userID, param.(string))
	case util.ClientCmdUpdateSession:
		err = onRetroSessionSave(s, *conn.Channel, userID, param.(map[string]interface{}))
	case util.ClientCmdAddFeedback:
		err = onAddFeedback(s, *conn.Channel, userID, param.(map[string]interface{}))
	case util.ClientCmdUpdateFeedback:
		err = onEditFeedback(s, *conn.Channel, userID, param.(map[string]interface{}))
	case util.ClientCmdRemoveFeedback:
		err = onRemoveFeedback(s, *conn.Channel, userID, param.(string))
	default:
		err = errors.New("unhandled retro command [" + cmd + "]")
	}
	return errors.WithStack(errors.Wrap(err, "error handling retro message"))
}

func onRetroSessionSave(s *Service, ch channel, userID uuid.UUID, param map[string]interface{}) error {
	title := util.ServiceTitle(param["title"].(string))
	categoriesString, ok := param["categories"].(string)
	if !ok {
		return errors.WithStack(errors.New(fmt.Sprintf("cannot parse [%v] as string", param["categories"])))
	}
	categories := util.StringToArray(categoriesString)
	if len(categories) == 0 {
		categories = retro.DefaultCategories
	}
	s.logger.Debug(fmt.Sprintf("saving retro session [%s] with categories [%s]", title, strings.Join(categories, ", ")))

	err := s.retros.UpdateSession(ch.ID, title, categories, userID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error updating retro session"))
	}

	err = sendRetroSessionUpdate(s, ch)
	return errors.WithStack(errors.Wrap(err, "error sending retro session"))
}

func sendRetroSessionUpdate(s *Service, ch channel) error {
	sess, err := s.retros.GetByID(ch.ID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error finding retro session [" + ch.ID.String() + "]"))
	}
	if sess == nil {
		return errors.WithStack(errors.Wrap(err, "cannot load retro session [" + ch.ID.String() + "]"))
	}

	msg := Message{Svc: util.SvcRetro.Key, Cmd: util.ServerCmdSessionUpdate, Param: sess}
	err = s.WriteChannel(ch, &msg)
	return errors.WithStack(errors.Wrap(err, "error sending retro session"))
}
